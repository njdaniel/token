package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/njdaniel/token/routers"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
