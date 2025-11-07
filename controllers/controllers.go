package controllers

import (
	"chatters-REST/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateChannelInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateMessageInput struct {
	ChannelID int    `json:"channel_id" binding:"required"`
	UserID    int    `json:"user_id" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
	Text      string `json:"text" binding:"required"`
}

func CreateUser(ctx *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateUserInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We make sure the entries are not empty
	if in.Username == "" || in.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
		return
	}

	// We encrypt the password since thats safer right?
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	in.Password = string(hashed)

	// We create the new user
	user := models.User{
		Username: in.Username,
		Password: in.Password,
	}

	// We add the user to the DB
	models.DB.Create(&user)

	// We can now return out the proper JSON, we don't return the password
	in.Password = ""
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(ctx *gin.Context) {
	// We create a credential from our user input and bind the JSON body
	var cred CreateUserInput
	if err := ctx.BindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	// We now try to find our user
	var user models.User
	if err := models.DB.Where("username = ?", cred.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	// We now try to compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// We return out the user without the password
	user.Password = ""
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateChannel(ctx *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateChannelInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We can now create our channel from the input
	channel := models.Channel{Name: in.Name}

	// We save the channel to the database
	models.DB.Create(&channel)

	// We return out the JSON
	ctx.JSON(http.StatusOK, gin.H{"channel": channel})
}

func CreateMessage(ctx *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateMessageInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We can now create our model from the input
	message := models.Message{
		ChannelID: in.ChannelID,
		UserID:    in.UserID,
		UserName:  in.UserName,
		Text:      in.Text,
	}

	// We create a new message in the db
	models.DB.Create(&message)

	// We return out the corresponding JSON
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func FindChannels(ctx *gin.Context) {
	// We create a slice of Channels
	var channels []models.Channel

	// We then try to find the channels
	models.DB.Find(&channels)

	// We then return out the JSON
	ctx.JSON(http.StatusOK, gin.H{"channels": channels})
}

func FindMessages(ctx *gin.Context) {
	// We create a slice of Messages
	var messages []models.Message

	// We then try to find the messages
	models.DB.Find(&messages)

	// We return out the proper JSON
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
