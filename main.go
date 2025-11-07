package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int    `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Text      string `json:"text"`
}

// Method to create a new User
func createUser(ctx *gin.Context, db *sql.DB) {
	// We parse the JSON request body inter the User struct
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We try to insert the new user into our DB
	result, err := db.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username,
		user.Password,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can then get the id of the new user
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can then return the new ID
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// Method to create a Channel
func createChannel(ctx *gin.Context, db *sql.DB) {
	// We parse the JSON request body into the Channel struct
	var channel Channel
	if err := ctx.ShouldBindJSON(&channel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We try to insert into the db
	result, err := db.Exec("INSERT INTO channels (name) VALUES (?)", channel.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We try to retrieve the ID
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can now return the id from the new channel
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// Method to create a new Message
func createMessage(ctx *gin.Context, db *sql.DB) {
	// We parse the JSON request body into the Message struct
	var message Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We try to insert the message into our database
	result, err := db.Exec(
		"INSERT INTO messages (channel_id, user_id, message) VALUES (?, ?, ?)",
		message.ChannelID,
		message.UserID,
		message.Text,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We try to retrieve the ID
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can now return the ID
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// Method to return a list of channels for the frontend
func listChannels(ctx *gin.Context, db *sql.DB) {
	// We query the DB for channels
	rows, err := db.Query("SELECT id, name FROM channels")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We make a slice of channels
	var channels []Channel
	for rows.Next() {
		// We make a new channel
		var channel Channel

		// We scan row into channel
		err := rows.Scan(&channel.ID, &channel.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		channels = append(channels, channel)
	}
	// We can now return the list of channels
	ctx.JSON(http.StatusOK, channels)
}

func listMessages(ctx *gin.Context, db *sql.DB) {
	// We get the channel id from the url
	channelID, err := strconv.Atoi(ctx.Query("channelID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can now parse optional limit for queries
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		// We set a default limit to 100 if not provided
		limit = 100
	}

	// We now parse the last message ID
	lastMessageID, err := strconv.Atoi(ctx.Query("lastMessageID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We can now query our DB
	rows, err := db.Query(
		"SELECT m.id, channel_id, user_id, u.username AS user_name, message FROM messages m LEFT JOIN users u ON u.id = m.user_id WHERE channel_id = ? AND m.id > ? ORDER BY m.id ASC LIMIT ?",
		channelID,
		lastMessageID,
		limit,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// We make a slice of messages
	var messages []Message

	for rows.Next() {
		// We make a new message
		var message Message

		// We can now scan row into message
		err := rows.Scan(
			&message.ID,
			&message.ChannelID,
			&message.UserID,
			&message.UserName,
			&message.Text,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// We append the message
		messages = append(messages, message)
	}

	// We can now return the messages to the frontend
	ctx.JSON(http.StatusOK, messages)
}

// Method for the frontend to user for user logins
func login(ctx *gin.Context, db *sql.DB) {
	// We need to parse the JSON body
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We make a query to our db
	row := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password)

	// We can then get the user id
	var id int
	err := row.Scan(id)
	if err != nil {
		// We check if the user was not found
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		// Otherwise we return the error reported
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// We can then return the ID
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func main() {
	// We get our current directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working directory", wd)

	// We attempt to open the db
	db, err := sql.Open("sqlite", wd+"/database.db")
	// We make sure to make a defer function to safely close
	// the db even if we have a fatal error
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// We create the gin router
	r := gin.Default()
	if err != nil {
		log.Fatal(err)
	}

	// We make some endpoints for gin router
	r.POST("/users", func(ctx *gin.Context) { createUser(ctx, db) })
	r.POST("/channels", func(ctx *gin.Context) { createChannel(ctx, db) })
	r.POST("/messages", func(ctx *gin.Context) { createMessage(ctx, db) })

	// We list the endpoints
	r.GET("/channels", func(ctx *gin.Context) { listChannels(ctx, db) })
	r.GET("/messages", func(ctx *gin.Context) { listMessages(ctx, db) })

	// We make a login endpoint
	r.POST("/login", func(ctx *gin.Context) { login(ctx, db) })

	// We try to run on local host
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
