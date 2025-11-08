package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Admin struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Message struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Text      string `json:"text"`
}

// We forward declare our database
var DB *gorm.DB

// Method to connect to our database
func ConnectDB() {
	// We use the sqlite driver to open our database using the base gorm config
	database, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	// We panic out if we are not successful
	if err != nil {
		panic(err)
	}

	// We then try to auto migrate all of our objects
	err = database.AutoMigrate(&Admin{}, &User{}, &Channel{}, &Message{})
	// We panic out if we are not successful
	if err != nil {
		panic(err)
	}

	// We can now assing our database
	DB = database
}
