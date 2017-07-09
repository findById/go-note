package main

import (
	"net/http"
	"log"
	"time"
)

type Router struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewMux(router []Router) *http.ServeMux {
	mux := http.NewServeMux()
	for _, item := range router {
		mux.HandleFunc(routerHandler(item.Method, item.Pattern, item.HandlerFunc))
	}
	return mux
}

func routerHandler(method, pattern string, handler func(http.ResponseWriter, *http.Request)) (string, http.HandlerFunc) {
	return pattern, func(w http.ResponseWriter, r *http.Request) {
		log.Println("[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
			r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]")
		if r.Method != method {
			http.Error(w, "HTTP method '"+r.Method+"' is not supported by this URL", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
