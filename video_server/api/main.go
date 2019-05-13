package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
	http.ListenAndServe(":8000", RegisterHandlers())
}
