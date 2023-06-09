package server

import (
	"log"
	"net/http"
	"os"
	"yggdra/proxy"
)

var logger = log.New(os.Stdout, "httpsproxy:", log.Llongfile|log.LstdFlags)

func Serve(listenAdress string) {
	server := &http.Server{
		Addr: listenAdress,
		// TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy.Serve(w, r)
		}),
	}

	logger.Fatal(server.ListenAndServe())
}
