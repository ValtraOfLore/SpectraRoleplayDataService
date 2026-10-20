package guilddata

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	DataColl *mongo.Collection
}

func NewService(coll *mongo.Collection) *Service {
	return &Service{
		DataColl: coll,
	}
}

func (s Service) GetGuildDataById(guildId string) (bson.M, error) {
	var guildData bson.M
	readErr := s.DataColl.FindOne(context.TODO(), bson.M{"guildId": guildId}).Decode(&guildData)
	if readErr != nil {
		fmt.Println("Fail to read guild data for guild: " + guildId + "\n" + readErr.Error())
		return nil, errors.New("Unable to read guild data")
	}
	return guildData, nil
}
