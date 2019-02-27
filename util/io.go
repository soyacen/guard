package util

import (
	"github.com/yacen/guard/resp"
	"net/http"
)

func SetJSONResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func WriteMessage(w http.ResponseWriter, code int, msg string) {

	SetJSONResponseHeader(w)

	w.WriteHeader(http.StatusOK)
	res := resp.BaseResp{
		Error: resp.Error{
			Code:    code,
			Message: msg,
		},
	}
	CheckWriteReturn(w.Write(res.ToJson()))
}

func WriteError(w http.ResponseWriter, err *resp.Error) {

	SetJSONResponseHeader(w)

	w.WriteHeader(http.StatusOK)
	res := resp.BaseResp{
		Error: *err,
	}
	CheckWriteReturn(w.Write(res.ToJson()))
}

func WriteData(w http.ResponseWriter, data interface{}) {

	SetJSONResponseHeader(w)

	w.WriteHeader(http.StatusOK)
	res := resp.BaseResp{
		Error: resp.Error{
			Code:    resp.SUCCESS,
			Message: "success",
		},
		Data: data,
	}
	CheckWriteReturn(w.Write(res.ToJson()))
}
