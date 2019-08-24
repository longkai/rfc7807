package rfc7807

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/xerrors"
)

// NewInstance provides the instance id each time.
var NewInstance = func() string {
	return "WARNING: PLEASE PROVIDE AN INSTANCE ID GENERATOR."
}

// Problem detail response content type.
const (
	ContentTypeXML  = "application/problem+xml; charset=utf8"
	ContentTypeJSON = "application/problem+json; charset=utf8"
	// ContentTypeJSONP = "application/problem+json; charset=utf8"
)

// ProblemDetail is the entity of RFC 7807.
type ProblemDetail struct {
	Type     string `json:"type" xml:"type"`
	Title    string `json:"title" xml:"title"`
	Status   int    `json:"status" xml:"status"`
	Detail   string `json:"detail" xml:"detail"`
	Instance string `json:"instance" xml:"instance"`

	err   error
	frame xerrors.Frame
}

func (e *ProblemDetail) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("rfc7807: %v", err) // Never happens.
	}
	return string(b)
}

// Format implements fmt.Formatter.
func (e *ProblemDetail) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

// FormatError implements xerrors.Formatter.
func (e *ProblemDetail) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

// Unwrap implements xerrors.Wrapper.
func (e *ProblemDetail) Unwrap() error { return e.err }

// New creates a new problem detail with the typo and detail without wrapping any error.
func New(typo, detail string) *ProblemDetail {
	status, title, ok := ExplainError(typo)
	if ok && detail == "" {
		detail = title
	}
	return &ProblemDetail{
		Type:     typo,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: NewInstance(),
	}
}

// Wrap wraps the err with the typo and detail into a problem detail.
func Wrap(typo, detail string, err error) *ProblemDetail {
	status, title, ok := ExplainError(typo)
	if ok && detail == "" {
		detail = title
	}
	return &ProblemDetail{
		Type:     typo,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: NewInstance(),
		err:      err,
		frame:    xerrors.Caller(1),
	}
}
