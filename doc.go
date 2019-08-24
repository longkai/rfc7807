// Package rfc7807 implements RFC 7807, Problem Details for HTTP APIs:
//   https://tools.ietf.org/html/rfc7807.
//
// This package predefined Google's global domain errors.
// To create a new problem detail, using
//
//    rfc7807.New(NotFound, "this item xxx is not found.")
//
// Or wrapping with a undelaying error
//
//    rfc7807.Wrap(InternalError, "internal server error.", err)
//
// It supports Go 2 error as values proposal:
//    https://go.googlesource.com/proposal/+/master/design/29934-error-values.md.
//
// Before using this package, you MUST provide a UUID generator function like
//
//    rfc7807.NewInstanceID = func() string { return uuid.New().String() }
//
// If the predefined errors domain doesn't safatify the requirements,
// you can patche the exported fields after obtaining one from the two methods above, `New` or `Wrap`.
package rfc7807
