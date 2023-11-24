package service

import (
	"micr/forth/model"
	"time"
)

func GetUsers() ([]*model.User, error) {
	//TODO storage
	return []*model.User{{
		Id: model.GenerateId(),
		Name: "John Doe",
		CreationTime: time.Now().UTC().Truncate(time.Millisecond),
	}}, nil
}