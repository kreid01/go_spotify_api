package main

import (
  	"github.com/gin-gonic/gin"
        "go_spotify_api/controllers" 
        "go_spotify_api/models" 


)


func main() {
	models.ConnectDatabase()

	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.PostUser)
	router.GET("/users/:id", controllers.GetUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	
	router.PUT("/users/likes/:id", controllers.UpdateUserLikes)
	router.DELETE("/users/likes/:id", controllers.DeleteUserLikes)

	router.PUT("/users/liked-albums/:id", controllers.UpdateLikedAlbum)
	router.DELETE("/users/liked-albums/:id", controllers.DeleteLikedAlbum)

	router.GET("/playlists", controllers.GetPlaylists)
	router.POST("/playlists", controllers.PostPlaylist)
	router.GET("/playlists/:id", controllers.GetPlaylist)
	router.DELETE("/playlists/:id", controllers.DeletePlaylist)
	router.PUT("/playlists/:id", controllers.UpdatePlaylist)

	router.Run()
}

