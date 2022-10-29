package routes

import (
	"github.com/LLuthfiY/fake-api-GO/controllers"
	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {
	userRoute := router.Group("/user")

	userRoute.GET("/", controllers.GetAllUser)
	userRoute.GET("/get", controllers.GetOneUser)
	userRoute.POST("/", controllers.CreateUser)
	userRoute.POST("/update", controllers.UpdateUser)
	userRoute.DELETE("/delete", controllers.DeleteUser)
}
