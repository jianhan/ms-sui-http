package handler

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/jianhan/ms-sui-ideas/proto/idea"
)

type ideaHandler struct {
	ideaClient idea.IdeaServiceClient
}

func (h *ideaHandler) addIdeas(w http.ResponseWriter, r *http.Request) {
}

func (h *ideaHandler) listIdeas(w http.ResponseWriter, r *http.Request) {
	rsp, err := h.ideaClient.ListIdeas(r.Context(), &idea.IdeaFilter{})
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(rsp)
	w.Write([]byte("test"))
}
