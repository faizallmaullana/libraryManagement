package profile

import (
	"errors"
	"net/http"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /profile/:id
func GetProfile(c *gin.Context) {
	// Fetch user based on the user ID from the route parameter
	var user models.Users
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User record not found"})
		return
	}

	// Fetch profile with eager loading of the associated user
	var profile models.Profile
	if err := models.DB.Preload("Users").Where("id_user = ?", user.ID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile record not found for the user"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}
