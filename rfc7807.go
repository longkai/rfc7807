package rfc7807

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// ProblemDetails is the entity of RFC 7807.
type ProblemDetails struct {
	Type       string      `json:"type" xml:"type"`
	Title      string      `json:"title" xml:"title"`
	Status     int         `json:"status" xml:"status"`
	Detail     string      `json:"detail" xml:"detail"`
	Instance   string      `json:"instance" xml:"instance"`
	Extensions interface{} `json:"extensions,omitempty" xml:"extensions,omitempty"`

	err   error
	frame xerrors.Frame
} // @name ProblemDetails

func (e *ProblemDetails) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("rfc7807: %v", err) // Never happens.
	}
	return string(b)
}

// Format implements fmt.Formatter.
func (e *ProblemDetails) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

// FormatError implements xerrors.Formatter.
func (e *ProblemDetails) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

// Unwrap implements xerrors.Wrapper.
func (e *ProblemDetails) Unwrap() error { return e.err }

// New creates a new problem detail with the typo and detail without wrapping any error.
func New(code Code, detail string) *ProblemDetails {
	val := mappings[code]
	return &ProblemDetails{
		Type:     val.typo,
		Title:    val.title,
		Status:   val.status,
		Detail:   detail,
		Instance: uuid.New().String(),
	}
}

// Wrap wraps the err with the typo and detail into a problem detail.
func Wrap(code Code, detail string, err error) *ProblemDetails {
	val := mappings[code]
	return &ProblemDetails{
		Type:     val.typo,
		Title:    val.title,
		Status:   val.status,
		Detail:   detail,
		Instance: uuid.New().String(),
		err:      err,
		frame:    xerrors.Caller(1),
	}
}

// Customize creates a limited domain rfc7807 problem detail.
// Note you should provide a domain to limit the namespace of the problem, since it's different from canonical error codes for Google APIs.
// So the problem type will be `domain.typo`.
// Each typo and title should correspond one-to-one.
// The cause may be nil if you do not have a cause error to wrap.
// The extensions normally nil if no additional infomation passing.
// Please consider using `New` or `Wrap` first, if they cannot convery your needs, then come back this func.
func Customize(domain, typo, title, detail string, status int, cause error, extensions interface{}) *ProblemDetails {
	if domain != "" {
		typo = domain + "." + typo
	}
	pd := &ProblemDetails{
		Type:       typo,
		Title:      title,
		Detail:     detail,
		Status:     status,
		Extensions: extensions,
		Instance:   uuid.New().String(),
	}
	if cause != nil {
		pd.err, pd.frame = cause, xerrors.Caller(1)
	}
	return pd
}
