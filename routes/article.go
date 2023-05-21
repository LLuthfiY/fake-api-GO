package routes

import (
	"github.com/LLuthfiY/fake-api-GO/controllers"
	"github.com/gin-gonic/gin"
)

func Article(router *gin.Engine) {
	articleRoute := router.Group("/article")

	articleRoute.GET("/", controllers.GetAllArticle)
	articleRoute.GET("/:id", controllers.GetArticle)
	articleRoute.POST("/", controllers.CreateArticle)
	articleRoute.POST("/update", controllers.UpdateArticle)
	articleRoute.DELETE("/delete", controllers.DeleteArticle)
}
