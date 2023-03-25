package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"yggdra/proxy"
)

var logger = log.New(os.Stderr, "httpsproxy:", log.Llongfile|log.LstdFlags)

func Serve(listenAdress string) {
	cert, err := genCertificate()
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Addr:      listenAdress,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy.Serve(w, r)
			logger.Print("HandlerFunc", r)
		}),
	}

	logger.Fatal(server.ListenAndServe())
}
