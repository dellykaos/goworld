package handler

import (
	"net/http"

	"github.com/dellykaos/goworld"

	"github.com/julienschmidt/httprouter"
)

// Handler controls request flow from client to service
type Handler struct {
}

// New initiate handler
func New() *Handler {
	return &Handler{}
}

// Hello is used to control the flow of GET /hello endpoint
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := goworld.Hello(params.ByName("name"))

	w.Write([]byte(res))
}
