package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/ms-sui-http/middleware"
	"github.com/urfave/negroni"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle(
		"/ideas/add",
		negroni.New(negroni.HandlerFunc(middleware.JWTMiddleware.HandlerWithNext), negroni.Wrap(http.HandlerFunc(addIdea))),
	).Methods("GET")
	return router
}
