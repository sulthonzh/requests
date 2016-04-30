package requests

import (
	"bytes"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

// HTTPResponse is an interface type implemented by *Response and *http.Response.
type HTTPResponse interface {
	Cookies() []*http.Cookie
	Location() (*url.URL, error)
	ProtoAtLeast(major, minor int) bool
	Write(w io.Writer) error
	Bytes() []byte
	String() string
	JSON() []byte
	Len() int
}

// Response is a *http.Response and implements HTTPResponse.
type Response struct {
	*http.Response
	Error error
}

// Len returns the response's body's unread portion's length,
// which is the full length provided it has not been read.
func (r *Response) Len() int {
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r.Body)
	bodyLen := buf.Len()
	return bodyLen
}

// String returns the response's body as string. Any errors
// reading from the body is ignored for convenience.
func (r *Response) String() string {
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r.Body)
	bodyStr := buf.String()
	return bodyStr
}

// Bytes returns the response's Body as []byte. Any errors
// reading from the body is ignored for convenience.
func (r *Response) Bytes() []byte {
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r.Body)
	bodyBytes := buf.Bytes()
	return bodyBytes
}

// JSON returns the response's body as []byte if Content-Type is
// in the header contains "application/json".
func (r *Response) JSON() []byte {
	jsn := []byte{}
	for _, ct := range r.Header["Content-Type"] {
		t, _, err := mime.ParseMediaType(ct)
		if err != nil {
			log.Panicln(err)
		}
		if strings.Contains(t, "application/json") {
			jsn = r.Bytes()
		}

	}
	return jsn
}
