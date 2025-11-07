package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&User{}, &Channel{}, &Message{})
	if err != nil {
		panic(err)
	}

	DB = database
}
