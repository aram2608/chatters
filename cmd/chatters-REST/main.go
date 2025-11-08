package main

import (
	"log"
	"time"

	"chatters-REST/controllers"
	"chatters-REST/middleware"
	"chatters-REST/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// We connect our database
	models.ConnectDB()

	// We create the gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length", "Authorization"},
		MaxAge:        12 * time.Hour,
	}))

	// We set the authentification
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())

	// We make some endpoints for gin router
	// we pass in the proper handlers for each endpoint
	r.POST("/admin", controllers.CreateAdmin)
	r.POST("/users", controllers.CreateUser)
	r.POST("/channels", controllers.CreateChannel)
	r.POST("/messages", controllers.CreateMessage)

	// We make GET requests that can return the list of channels and
	// messages from our endpoints
	auth.GET("/channels", controllers.GetChannels)
	auth.GET("/channels/:id/messages", controllers.GetChannelMessages)
	auth.GET("/messages", controllers.GetMessages)
	auth.GET("/users", controllers.GetUsers)

	// We make a login endpoint with the proper login handler
	r.POST("/login", controllers.Login)

	// We try to run on local host
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
