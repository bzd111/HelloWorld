package main

import (
	"gin_graphql"
	"gin_graphql/models"
)

func main() {
	db := gin_graphql.GetInstance()
	defer db.Close()
	db.AutoMigrate(&models.Todo{}, &models.User{})
}
