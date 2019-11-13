package rfc7807_test

import (
	"fmt"
	"io"

	"github.com/longkai/rfc7807"
	"golang.org/x/xerrors"
)

func ExampleNew() {
	err := rfc7807.New(rfc7807.NotFound, "the item is not found")
	err.Instance = "example only, do not set this field"
	fmt.Println(err)
	// Output: {"type":"NOT_FOUND","title":"A specified resource is not found, or the request is rejected by undisclosed reasons, such as whitelisting.","status":404,"detail":"the item is not found","instance":"example only, do not set this field"}
}

func ExampleWrap() {
	cause := xerrors.Errorf("read /path/to/file: %w", io.EOF)
	err := rfc7807.Wrap(rfc7807.Internal, "", cause)
	err.Instance = "example only, do not set this field"
	// Note: the output prints the stacktrace of the error wrapping chain, so it won't be passed in other environment.
	// Just showing the wrapping output format.
	fmt.Printf("%+v\n", err)
	// Output: {"type":"INTERNAL","title":"Internal server error. Typically a server bug.","status":500,"detail":"","instance":"example only, do not set this field"}:
	//     github.com/longkai/rfc7807_test.ExampleWrap
	//         /Users/longkai/Dropbox/src/rfc7807/example_test.go:20
	//   - read /path/to/file:
	//     github.com/longkai/rfc7807_test.ExampleWrap
	//         /Users/longkai/Dropbox/src/rfc7807/example_test.go:19
	//   - EOF
}

func ExampleCustomize() {
	err := rfc7807.Customize("com.example", "METHOD_NOT_ALLOWED", "The request method is not allowed.", "The api only allows HTTP GET, got POST.", 405, nil, nil)
	err.Instance = "example only, do not set this field"
	fmt.Println(err)
	// Output: {"type":"com.example.METHOD_NOT_ALLOWED","title":"The request method is not allowed.","status":405,"detail":"The api only allows HTTP GET, got POST.","instance":"example only, do not set this field"}
}
