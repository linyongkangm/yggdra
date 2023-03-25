package proxy

import (
	"net/http"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		handleHttps(w, r)
	} else {
		handleHttp(w, r)
	}
}
