package health

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(serv *Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h Handler) Get(rw http.ResponseWriter, r *http.Request) {
	healthMsg := h.service.GetHealth()
	json.NewEncoder(rw).Encode(healthMsg)
}
