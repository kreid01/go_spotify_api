package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	ID uint `gorm:"primaryKey"` 
	Name string `json:"name"`
	SongIds []string `json:"songIds" gorm:"type:json"` 
	UserID uint `json:"userId"`
}

type PlaylistInput struct {
	Name string `json:"name"`
	UserID uint `json:"userId"`
}

type UpdatePlaylistInput struct {
	Name string `json:"name"`
	SongIds []string `json:"songIds" gorm:"type:json"` 
}
