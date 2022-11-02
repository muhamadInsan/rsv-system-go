package routes

import (
	"rsv-system-go/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("users", controllers.FetchUsers)
	r.POST("users", controllers.CreateUser)
	r.GET("users/:id", controllers.FetchUserById)
	r.PUT("users/:id", controllers.UpdateUser)
	r.DELETE("users/:id", controllers.DeleteUser)

	r.POST("login", controllers.Login)
}
