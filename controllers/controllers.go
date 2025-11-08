package controllers

import (
	"chatters-REST/config"
	"chatters-REST/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CreateAdminInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Function used to create an admin
func CreateAdmin(c *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateAdminInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We make sure the entries are not empty
	if in.Username == "" || in.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
		return
	}

	// We encrypt the password since thats safer right?
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
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
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Function used to create a new user
func CreateUser(c *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateUserInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We make sure the entries are not empty
	if in.Username == "" || in.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
		return
	}

	// We encrypt the password since thats safer right?
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
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
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Function to login a user that has an account
func Login(c *gin.Context) {
	// We create a credential from our user input and bind the JSON body
	var cred CreateUserInput
	if err := c.BindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	// We now try to find our user
	var user models.User
	if err := models.DB.Where("username = ?", cred.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	// We now try to compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expiration := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	// We return out the user without the password
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": signed,
	})

	// Cookie instead
	// c.SetCookie("auth_token", signed, 3600*24, "/", "domain-example.com", true, true)
}

// Function to create a channel
func CreateChannel(c *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateChannelInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We can now create our channel from the input
	channel := models.Channel{Name: in.Name}

	// We save the channel to the database
	models.DB.Create(&channel)

	// We return out the JSON
	c.JSON(http.StatusOK, gin.H{"channel": channel})
}

// Function to create a message
func CreateMessage(c *gin.Context) {
	// We create an input and attempt to bind the JSON body
	var in CreateMessageInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Function to get our channel names for the front end
func GetChannels(c *gin.Context) {
	// We create a slice of Channels
	var channels []models.Channel

	// We then try to find the channels
	models.DB.Find(&channels)

	// We then return out the JSON
	c.JSON(http.StatusOK, gin.H{"channels": channels})
}

// Function to get the channel messages for the front end
func GetChannelMessages(c *gin.Context) {
	// We query the id
	// .GET("/users/:id")
	id := c.Param("id")

	// We create a slice of messages
	var messages []models.Message
	// We need to query given our ID, we order by ascending order then find
	// our messages
	if err := models.DB.Where("channel_id = ?", id).Order("id ASC").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// We can now return our messages
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// Function to get our messages for our front end
func GetMessages(c *gin.Context) {
	// We create a slice of Messages
	var messages []models.Message

	// We then try to find the messages
	models.DB.Find(&messages)

	// We return out the proper JSON
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// Function to get our users for the admin page
func GetUsers(c *gin.Context) {
	// We create a slice of users
	var users []models.User

	// we then try to find the users
	models.DB.Find(&users)

	// We return out the proper JSON
	c.JSON(http.StatusOK, gin.H{"users": users})
}
