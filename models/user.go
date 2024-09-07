package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Playlists  []Playlist `json:"playlists" gorm:"foreignKey:UserID"`
	Likes []string `json:"likes" gorm:"type:json"`
}

type UserInput struct {
	Name string `json:"name"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}
