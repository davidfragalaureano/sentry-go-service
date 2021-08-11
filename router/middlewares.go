package router

import (
	"log"
	"net/http"
	"time"
)

func HTTPLogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqTime := time.Now()
		log.Printf("%v Request to \"%v\"", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
		log.Printf("Request took: \"%f\" sec", time.Now().Sub(reqTime).Seconds())
	})
}
