package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m

}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	router.GET("/user/:user_name/videos", Login)

	router.GET("/user/:user_name/videos/:vid-id", Login)

	router.DELETE("/user/:user_name/videos/:vid-id", Login)

	router.GET("/videos/:vid-id/comments", Login)

	router.POST("/videos/:vid-id/comments", Login)

	router.DELETE("/videos/:vid-id/comment/:comment-id", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
