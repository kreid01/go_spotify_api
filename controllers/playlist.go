package controllers 

import (
	"go_spotify_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlaylists(c *gin.Context) {
	var playlists []models.Playlist
	models.DB.Find(&playlists)

	c.JSON(http.StatusOK, gin.H{"data": playlists})
}

func GetPlaylist(c *gin.Context) {
	var playlist models.Playlist
	models.DB.Find(&playlist)

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&playlist).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return 
	}

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

func PostPlaylist(c * gin.Context) {
	var input models.PlaylistInput

	 if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var playlists  []models.Playlist;
  var count int64
  models.DB.Find(&playlists).Count(&count)

  playlist := models.Playlist{Name: input.Name, ID: uint(count + 1), UserID: 1}
  models.DB.Create(&playlist)

    c.JSON(http.StatusOK, gin.H{"data": playlist})
}


func UpdatePlaylist(c * gin.Context) {
	var input models.UpdatePlaylistInput
	 if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }


	var playlist models.Playlist
	if err := models.DB.Where("ID = ?", c.Param("id")).First(&playlist).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return 
	}

	
	models.DB.Model(&playlist).Updates(input)

    c.JSON(http.StatusOK, gin.H{"data": playlist})
}

func DeletePlaylist(c *gin.Context) {
	var playlist models.Playlist
	models.DB.Find(&playlist)

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&playlist).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return 
	}

	models.DB.Delete(&playlist)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
