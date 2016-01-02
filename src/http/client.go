package http

import (
    "util"
    "http/request"
    "http/response"
    "http/headers"
)

type Client struct {
    request    *request.Request
    response   *response.Response
}

func NewClient() (*Client) {
    return &Client{
        request: NewRequest(),
        response: NewResponse(),
    }
}

func (this *Client) Request() (*request.Request) {
    return this.request
}
func (this *Client) Response() (*response.Response) {
    return this.response
}

func (this *Client) Do(u string, up interface{}, b interface{}, h interface{}) (*response.Response) {
    m, _, err := util.RegExpMatch(u, "^([A-Z]+)\\s+(.+)")
    if len(m) < 3 {
        panic("Usage: <Method GET|POST...> <Scheme http|https>://<Host>/<Path>... !")
    }
    util.Dumps(m[1], err, len(m))
    this.request.SetMethod(m[1])
    this.request.SetUri(m[2], up)
    this.request.SetBody(b)
    this.request.SetHeaderAll(h)

    rs, err := this.request.Send()
    if err != nil {
        panic(err)
    }

    rt := util.Explode(rs, "\r\n\r\n", 2)
    if len(rt) != 2 {
        panic("No valid response returned from server!")
    }

    rh := headers.Parse(rt[0])
    if _, ok := rh["0"]; ok {
        this.response.SetStatus(rh["0"])
    }
    this.response.SetHeaderAll(rh)
    this.response.SetBody(rt[1])

    if this.response.Status().Code() >= 400 {
        this.response.SetError(
            this.response.Status().Code(),
            this.response.Status().Text(),
        )
    }

    return this.response
}

func (this *Client) Head() {}
func (this *Client) HeadFunc() {}

// ...
