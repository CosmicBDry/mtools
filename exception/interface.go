package exception

import (
	"encoding/json"
)

type ApiException interface {
	error
	ToJson() string
	SetHttpCode(int)
	GetHttpCode() int
	GetErrorCode() int
	SetData(interface{}) ApiException
	GetData() interface{}
	SetMeta(interface{}) ApiException
	GetMeta() interface{}
	SetNamespace(ns string)
	GetNamespace() string
	GetReason() string
	Is(error) bool
}

type exception struct {
	Namespace string `json:"namespace"`

	Httpcode int `json:"http_code"`

	ErrorCode int `json:"error_code"`

	Message string `json:"message"`

	Reason string `json:"reason"`

	Data interface{} `json:"data"`

	Meta interface{} `json:"meta"`
}

func (ex *exception) Error() string {
	return ex.Message
}

func (ex *exception) ToJson() string {

	bytes, _ := json.MarshalIndent(ex, "", "  ")

	return string(bytes)

}

func (ex *exception) Is(e error) bool {

	if v, ok := e.(ApiException); ok {
		return v.GetErrorCode() == ex.GetErrorCode()
	}

	return ex.Message == e.Error()

}

func (ex *exception) SetNamespace(ns string) {
	ex.Namespace = ns
}

func (ex *exception) GetNamespace() string {
	return ex.Namespace
}

func (ex *exception) GetErrorCode() int {
	return ex.ErrorCode
}

func (ex *exception) SetHttpCode(httpcode int) {
	ex.Httpcode = httpcode
}
func (ex *exception) GetHttpCode() int {
	return ex.Httpcode
}

func (ex *exception) SetData(d interface{}) ApiException {
	ex.Data = d
	return ex
}

func (ex *exception) GetData() interface{} {
	return ex.Data
}

func (ex *exception) SetMeta(m interface{}) ApiException {
	ex.Meta = m
	return ex
}

func (ex *exception) GetMeta() interface{} {
	return ex.Meta
}

func (ex *exception) GetReason() string {
	return ex.Reason
}
