package controllers

import (
	"FGA_Hacktiv8/jwt/database"
	"FGA_Hacktiv8/jwt/helpers"
	"FGA_Hacktiv8/jwt/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	var user models.User

	switch contentType {
	case "application/json":
		ctx.ShouldBindJSON(&user)
	default:
		ctx.ShouldBind(&user)
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.Fullname,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	var user models.User

	switch contentType {
	case "application/json":
		ctx.ShouldBindJSON(&user)
	default:
		ctx.BindJSON(&user)
	}

	password := user.Password

	if err := db.Debug().Where("email = ?", user.Email).Take(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "invalid email/password",
		})

		return
	}

	if match := helpers.ComparePassword(password, user.Password); !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "invalid email/password",
		})

		return
	}

	jwtToken := helpers.GenerateToken(user.ID, user.Email)

	fmt.Println("create token success")

	ctx.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
