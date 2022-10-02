package router

import (
	"FGA_Hacktiv8/jwt/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	rUser := r.Group("/users")
	rUser.POST("/register", controllers.UserRegister)
	rUser.POST("/login", controllers.UserLogin)

	return r
}
