package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/ms-sui-http/middleware"
	"github.com/urfave/negroni"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Handle(
		"/ideas",
		negroni.New(negroni.HandlerFunc(middleware.JWTMiddleware.HandlerWithNext), negroni.Wrap(http.HandlerFunc(addIdea))),
	).Methods("POST")
	return router
}
