package params

import (
    "ohttp/util"
)

type Params map[string]interface{}

func New() (*Params) {
    return &Params{}
}
func (this *Params) Set(k string, v interface{}) (*Params) {
    (*this)[k] = v
    return this
}
func (this *Params) Get(k string) (interface{}) {
    return (*this)[k]
}
func (this *Params) GetInt(k string) (int) {
    return util.Int((*this)[k])
}
func (this *Params) GetUInt(k string) (uint) {
    return util.UInt((*this)[k])
}
func (this *Params) GetString(k string) (string) {
    return util.String((*this)[k])
}
func (this *Params) GetBool(k string) (bool) {
    return util.Bool((*this)[k])
}

func (this *Params) Empty() (bool) {
    for k, v := range (*this) {
        _ = k; _ = v
        return false
    }
    return true
}

func (this *Params) String() (string) {
    m := util.Map()
    for k, v := range (*this) {
        m[k] = v
    }
    return util.UrlQueryUnparse(m)
}