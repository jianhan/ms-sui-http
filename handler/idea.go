package handler

import "net/http"

func addIdea(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
