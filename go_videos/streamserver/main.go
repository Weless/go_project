package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too Many Requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middlewareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddlewareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
