package rfc7807

import (
	"testing"

	"golang.org/x/xerrors"
)

func TestIs(t *testing.T) {
	err := xerrors.New("err")
	wrapErr := Wrap(NotFound, "", err)
	if !xerrors.Is(wrapErr, err) {
		t.Errorf("xerrors.Is(%v, %v) = false", wrapErr, err)
	}
}

func TestAs(t *testing.T) {
	err := xerrors.New("err")
	wrapErr := Wrap(NotFound, "", err)
	e := &ProblemDetail{}
	if !xerrors.As(wrapErr, &e) { // Set pointer's pointer.
		t.Errorf("xerror.As(%v, %v) = false", wrapErr, &e)
	}
}
