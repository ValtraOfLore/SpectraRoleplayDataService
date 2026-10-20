package guilddata

import (
	"encoding/json"
	"fmt"
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
	guildId := r.PathValue("guild_id")

	guildData, err := h.service.GetGuildDataById(guildId)
	if err != nil {
		fmt.Println("Error while reading guild data for: " + guildId)
		return
	}

	body := GuildInfoResponse{
		GuildId:   guildId,
		GuildData: guildData,
	}

	json.NewEncoder(rw).Encode(body)
}
