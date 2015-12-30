package http

import (
    // _fmt "fmt"
)

import (
    "util"
    "http/headers"
)

type Stream struct {
    type_       string
    httpVersion string
    headers     *headers.Headers
    body        string
    StreamBody
    StreamString
}

type StreamBody interface {
    SetBody(body interface{})
}

type StreamString interface {
    String() (string)
}

const (
    TYPE_REQUEST  = "request"
    TYPE_RESPONSE = "response"
)

const (
    HTTP_VERSION_1_0 = "1.0"
    HTTP_VERSION_1_1 = "1.1"
    HTTP_VERSION_2_0 = "2.0"
)

func Shutup() {
    util.Shutup()
}

func NewStream(type_, httpVersion string) (*Stream) {
    return &Stream{
              type_: type_,
        httpVersion: httpVersion,
            headers: headers.New(),
               body: "",
    }
}

func (this *Stream) SetType(type_ string) {
    this.type_ = type_
}
func (this *Stream) GetType() (string) {
    return this.type_
}

func (this *Stream) SetHttpVersion(httpVersion string) {
    this.httpVersion = httpVersion
}
func (this *Stream) GetHttpVersion() (string) {
    return this.httpVersion
}
