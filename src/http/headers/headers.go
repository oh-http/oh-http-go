package headers

import (
    "util"
)

type Headers struct {
    data map[string]string
}

func New() (*Headers) {
    return &Headers{
        data: util.MapString(),
    }
}

func (this *Headers) Set(name, value string) (*Headers) {
    this.data[name] = value
    return this
}
func (this *Headers) Get(name string) (string) {
    if value, ok := this.data[name]; ok {
        return value
    }
    return ""
}
