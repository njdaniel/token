package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenicationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")

	router.HandleFunc("/refresh-token-auth", negroni.New(negroni.HandlerFunc(controllers.RefreshToken))).Methods("GET")

	router.Handle("/logout", negroni.New(negroni.HandlerFunc(authenication.RequireTokenAuthenication), negroni.HandlerFunc(controllers.Logout))).Methods("GET")

	return router
}
