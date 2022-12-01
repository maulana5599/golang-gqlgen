package handler

import (
	"log"

	"github.com/maulana5599/golang-gqlgen/config"
	models "github.com/maulana5599/golang-gqlgen/models"

	"github.com/maulana5599/golang-gqlgen/graph/model"
)

func GetAllUser() []*model.User {

	todos := []*model.User{
		{
			ID:   "Maulana Muhammad Rizky",
			Name: "hahahah",
		},
	}
	return todos
}

func GetCustomUser() []*models.Learn {
	var user []*models.Learn
	tx := config.DB.Find(&user)

	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	return user
}
