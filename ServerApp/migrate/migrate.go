package main

import (
	"github.com/savkela/property-app/initializers"
	"github.com/savkela/property-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
