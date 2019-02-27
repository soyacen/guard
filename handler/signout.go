package handler

import (
	"github.com/yacen/guard/resp"
	"github.com/yacen/guard/util"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	util.SetJSONResponseHeader(w)
	util.WriteMessage(w, resp.SUCCESS, "sign out")
}
