# RFC 7807

[![GoDoc](https://godoc.org/github.com/skailhq/rfc7807?status.svg)](http://godoc.org/github.com/skailhq/rfc7807)

Package rfc7807 implements RFC 7807, Problem Details for HTTP APIs:
	https://tools.ietf.org/html/rfc7807.

This package predefined Google's gRPC canonical error codes:
	https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto.

To create a new problem detail:

	rfc7807.New(rfc7807.NotFound, "xxx is not found.")

Or wrap with a underlying error:

	rfc7807.Wrap(rfc7807.Internal, "", causeError)

It supports Go 2 error as values proposal:
	https://go.googlesource.com/proposal/+/master/design/29934-error-values.md.

If the predefined errors doesn't satisfy your needs:

	rfc7807.Customize("my.error.domain", "MY_ERROR_TYPE", 400, nil, nil)

### Install

`go get github.com/skailhq/rfc7807`
