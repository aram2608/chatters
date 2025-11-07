package main

import (
	"log"

	"chatters-REST/controllers"
	"chatters-REST/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// We connect our database
	models.ConnectDB()

	// We create the gin router
	r := gin.Default()
	r.Use(cors.Default())

	// We make some endpoints for gin router
	// we pass in the proper handlers for each endpoint
	r.POST("/users", controllers.CreateUser)
	r.POST("/channels", controllers.CreateChannel)
	r.POST("/messages", controllers.CreateMessage)

	// We make two GET requests that can return the list of channels and
	// messages from our endpoints
	r.GET("/channels", controllers.FindChannels)
	r.GET("/messages", controllers.FindMessages)

	// We make a login endpoint with the proper login handler
	r.POST("/login", controllers.Login)

	// We try to run on local host
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
