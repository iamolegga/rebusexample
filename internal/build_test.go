package internal_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/iamolegga/rebusexample/internal"
)

func TestBuild(t *testing.T) {
	h := internal.Build()
	srv := httptest.NewServer(h)
	defer srv.Close()

	type Case struct {
		request          string
		expectedResponse string
	}

	tests := []Case{
		{"GET::/todos::", "200::[]"},
		{"POST::/todos::foo", `200::{"ID":1,"Payload":"foo"}`},
		{"GET::/todos::", `200::[{"ID":1,"Payload":"foo"}]`},
		{"PUT::/todos/1::bar", `200::{"ID":1,"Payload":"bar"}`},
		{"GET::/todos/1::", `200::{"ID":1,"Payload":"bar"}`},
		{"GET::/todos/2::", `404::`},
		{"POST::/todos::baz", `200::{"ID":2,"Payload":"baz"}`},
		{"GET::/todos::", `200::[{"ID":1,"Payload":"bar"},{"ID":2,"Payload":"baz"}]`},
		{"DELETE::/todos/1::", `200::`},
		{"GET::/todos::", `200::[{"ID":2,"Payload":"baz"}]`},
	}
	for _, tt := range tests {
		t.Run(tt.request, func(t *testing.T) {
			req := NewRequest(tt.request)
			res := req.For(srv.URL)
			if !res.Equals(tt.expectedResponse) {
				t.Error("not equals")
			}
		})
	}
}

type Request struct {
	method  string
	path    string
	payload string
}

func NewRequest(s string) *Request {
	data := strings.Split(s, "::")
	return &Request{data[0], data[1], data[2]}
}

func (r *Request) For(server string) *Response {
	client := http.Client{Timeout: time.Minute}
	req, _ := http.NewRequest(r.method, server+r.path, bytes.NewBufferString(r.payload))
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	bb, _ := ioutil.ReadAll(resp.Body)

	return &Response{
		status:  resp.StatusCode,
		payload: string(bb),
	}
}

type Response struct {
	status  int
	payload string
}

func (r *Response) Equals(s string) bool {
	data := strings.Split(s, "::")
	if len(data) != 2 {
		return false
	}
	status, _ := strconv.Atoi(data[0])
	payload := data[1]
	return status == r.status && payload == r.payload
}
