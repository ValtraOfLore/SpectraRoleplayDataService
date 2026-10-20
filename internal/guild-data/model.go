package guilddata

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GuildInfoResponse struct {
	GuildId   string `json:"guildId"`
	GuildData bson.M `json:"guildData"`
}
