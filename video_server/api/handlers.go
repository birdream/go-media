package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser hello
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// Login login
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u_name := p.ByName("user_name")
	io.WriteString(w, u_name)
}
