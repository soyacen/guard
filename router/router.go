package router

import (
	"github.com/yacen/guard/handler"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// 默认
	mux.HandleFunc("/", handler.RootHandler())

	// 登录
	mux.HandleFunc("/signin", handler.SignIn)

	// 注册
	mux.HandleFunc("/signup", handler.SignUp)

	// 退出
	mux.HandleFunc("/signout", handler.SignOut)

	return mux
}
