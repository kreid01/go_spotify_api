package controllers

import (
        "net/http"
        "github.com/lib/pq" 
	
	"github.com/gin-gonic/gin"
	"go_spotify_api/models"
)

func PostUser(c *gin.Context) {
	var input models.UserInput

	 if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

 var users []models.User
 var count int64

  models.DB.Find(&users).Count(&count)

  user := models.User{Name: input.Name, ID: uint(count + 1)}
  models.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	var user models.User

	  result := models.DB.Preload("Playlists").First(&user, c.Param("id"))

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }


	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c * gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func UpdateUser(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateUserInput
	  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	models.DB.Model(&user).Updates(input)
  c.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteUser(c * gin.Context) {
	var user models.User

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateUserLikes(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateUserLikesInput	
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	 user.Likes = append(user.Likes, input.LikeId)

	 print(user.Likes, input.LikeId)

    if err := models.DB.Model(&user).Update("likes", pq.Array(user.Likes)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserLikes(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateUserLikesInput 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.Likes = Filter(user.Likes, func(str string) bool {return str != input.LikeId })

    if err := models.DB.Model(&user).Update("likes", pq.Array(user.Likes)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateLikedAlbum(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateLikedAlbum 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.LikedAlbums = append(user.LikedAlbums, input.AlbumId)

    if err := models.DB.Model(&user).Update("liked_albums", pq.Array(user.LikedAlbums)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update liked albums"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteLikedAlbum(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateLikedAlbum 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.LikedAlbums = Filter(user.LikedAlbums, func(str string) bool {return str != input.AlbumId })
	

    if err := models.DB.Model(&user).Update("liked_albums", pq.Array(user.LikedAlbums)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update liked albums"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateFollowedArtists(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateFollowedArtistInput 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.FollowedArtists = append(user.FollowedArtists, input.ArtistsId)

    if err := models.DB.Model(&user).Update("followed_artists", pq.Array(user.FollowedArtists)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update liked albums"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteFollowedArtist(c* gin.Context) {
	var user models.User;

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    		return
	}

	var input models.UpdateFollowedArtistInput 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.FollowedArtists = Filter(user.FollowedArtists, func(str string) bool {return str != input.ArtistsId })
	

    if err := models.DB.Model(&user).Update("followed_artists", pq.Array(user.FollowedArtists)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update liked albums"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"data": user})
}


func Filter(strings []string, predicate func(string) bool) []string {
	var result []string
	for _, string := range strings {
		if predicate(string) {
			result = append(result, string)
		}
	}
	return result
}

