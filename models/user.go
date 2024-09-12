package models

import ("gorm.io/gorm"
  "github.com/lib/pq"
  )

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Playlists  []Playlist `json:"playlists" gorm:"foreignKey:UserID"`
	Likes pq.StringArray `json:"likes" gorm:"type:text[]"`
}

type UserInput struct {
	Name string `json:"name"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}


type UpdateUserLikesInput struct {
	Likes []string `json:"likes" gorm:"type:json"`
}

