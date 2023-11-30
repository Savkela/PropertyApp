package main

import (
	"ServerApp/entity"
	"ServerApp/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&entity.User{})
}
