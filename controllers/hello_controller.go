package controllers

import (
	"net/http"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandleFunc) {
	w.Write([]byte("Hello, World!"))
}
