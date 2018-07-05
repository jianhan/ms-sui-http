package handler

import "github.com/gorilla/mux"

func Router() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ideas/add", addIdea)
	return router
}
