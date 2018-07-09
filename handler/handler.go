package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/ms-sui-http/middleware"
	"github.com/jianhan/ms-sui-ideas/proto/idea"
	"github.com/urfave/negroni"
)

func Router(ideaClient idea.IdeaServiceClient) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// start api
	api := router.PathPrefix("/api/v1").Subrouter()

	// ideas
	ideaHandler := &ideaHandler{ideaClient: ideaClient}

	// create new idea
	api.Handle(
		"/ideas",
		negroni.New(negroni.HandlerFunc(middleware.JWTMiddleware.HandlerWithNext), negroni.Wrap(http.HandlerFunc(ideaHandler.addIdeas))),
	).Methods("POST")

	// list ideas
	api.HandleFunc(
		"/ideas",
		ideaHandler.listIdeas,
	).Methods("GET")

	return router
}
