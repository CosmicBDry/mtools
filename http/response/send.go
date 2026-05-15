package response

import (
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/CosmicBDry/mtools/exception"
)

func Failed(w http.ResponseWriter, err error, ops ...Options) {

	var (
		errcode  int
		httpcode int
		message  string
		reason   string
		data     interface{}
		meta     interface{}
	)

	//判断是否为ApiException接口类型
	switch t := err.(type) {
	case exception.ApiException:
		errcode = t.GetErrorCode()
		httpcode = t.GetHttpCode()
		message = t.Error()
		reason = t.GetReason()
		data = t.GetData()
		meta = t.GetMeta()
	default:
		httpcode = exception.InternalServerError
	}

	if httpcode == 0 {
		httpcode = exception.InternalServerError
	}

	resp := &Data{
		Code:    &errcode,
		Message: message,
		Reason:  reason,
		Meta:    meta,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")

	for _, op := range ops {
		op.Apply(resp)
	}

	resp_bytes, err := json.Marshal(resp)

	if err != nil {
		httpcode = exception.InternalServerError
		w.WriteHeader(httpcode)
		w.Write([]byte(fmt.Sprintf(`{"status": "error","errMsg":"json.Marshal,%s"}`, err.Error())))
		return

	}

	w.WriteHeader(httpcode)

	w.Write(resp_bytes)

}

func Success(w http.ResponseWriter, data interface{}, ops ...Options) {

	code := 0

	resp := &Data{
		Code: &code,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")

	resp_bytes, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(exception.InternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"status": "error","errMsg":"json.Marshal,%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(resp_bytes)

}
