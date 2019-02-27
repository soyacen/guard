package handler

import (
	"github.com/yacen/guard/util"
	"net/http"
)

func RootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.SetJSONResponseHeader(w)
		util.WriteMessage(w, http.StatusOK, "hello guard")
	}
}
