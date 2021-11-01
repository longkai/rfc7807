package rfc7807_test

import (
	"testing"

	"github.com/skailhq/rfc7807"
	"golang.org/x/xerrors"
)

func TestIs(t *testing.T) {
	err := xerrors.New("err")
	wrapErr := rfc7807.Wrap(rfc7807.NotFound, "", err)
	if !xerrors.Is(wrapErr, err) {
		t.Errorf("xerrors.Is(%v, %v) = false", wrapErr, err)
	}
}

func TestAs(t *testing.T) {
	err := xerrors.New("err")
	wrapErr := rfc7807.Wrap(rfc7807.NotFound, "", err)
	e := &rfc7807.ProblemDetails{}
	if !xerrors.As(wrapErr, &e) { // Set pointer's pointer.
		t.Errorf("xerror.As(%v, %v) = false", wrapErr, &e)
	}
}
