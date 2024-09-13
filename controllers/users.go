package controllers

import (
"net/http"
    "github.com/lib/pq" 
    "fmt"


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

	 user.Likes = append(user.Likes, input.Likes...)

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

	var input models.DeleteUserLikesInput 
	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    		return
  	}

	user.Likes = Filter(user.Likes, func(str string) bool {return str != input.LikeId })

	fmt.Print(user.Likes)
	

    if err := models.DB.Model(&user).Update("likes", pq.Array(user.Likes)).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes"})
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

