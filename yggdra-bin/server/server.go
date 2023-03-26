package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"yggdra/ca"
	"yggdra/proxy"
)

var logger = log.New(os.Stdout, "httpsproxy:", log.Llongfile|log.LstdFlags)

func Serve(listenAdress string) {
	cert, err := ca.GetTLSConfig()
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Addr:      listenAdress,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Print(2, r)
			proxy.Serve(w, r)
		}),
	}

	logger.Fatal(server.ListenAndServe())
}
