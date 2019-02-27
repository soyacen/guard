package server

import (
	"fmt"
	"github.com/yacen/guard/config"
	"github.com/yacen/guard/router"
	"github.com/yacen/guard/util/log"
	"net/http"
)

func Start() {
	port := config.Cfg.Port
	https := config.Cfg.Https
	certFile := config.Cfg.CertFile
	keyFile := config.Cfg.KeyFile
	handler := router.NewRouter()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	log.Println("listening on", port)
	if https {
		log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
