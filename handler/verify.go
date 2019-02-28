package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/yacen/guard/util/log"

	"github.com/yacen/guard/resp"
	"github.com/yacen/guard/util"

	"github.com/dgrijalva/jwt-go/request"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (key interface{}, e error) {
		key = util.RsaPub
		return
	})
	if err != nil {
		log.Error("parse from request error,", err)
		util.WriteError(w, resp.NewError(resp.ParseFromError, "verify error"))
		return
	}
	if token.Valid {
		util.WriteMessage(w, resp.SUCCESS, "success")
	} else {
		util.WriteError(w, resp.NewError(resp.VerifyError, "verify error"))
	}
}
