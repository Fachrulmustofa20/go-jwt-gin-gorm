package controller

import (
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegistration(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	user := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}
	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	ComparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !ComparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
