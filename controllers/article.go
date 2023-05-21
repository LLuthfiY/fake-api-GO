package controllers

import (
	"net/http"

	"github.com/LLuthfiY/fake-api-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type getArticle struct {
	ID string `json:"id"`
}

type updateArticleRequest struct {
	ID      string `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type createArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

type articleResponse struct {
	ID      string
	Title   string
	Content string
}

func GetAllArticle(ctx *gin.Context) {
	var article []models.Article
	result := db.Preload("User").Find(&article)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": &result})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &article})
}

func GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	var article models.Article
	result := db.Model(&article).Preload("User").Take(&article, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": result.Error})
		return
	}

	response := articleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &response})
}

func CreateArticle(ctx *gin.Context) {
	var data createArticleRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erroes": err.Error()})
		return
	}
	article := models.Article{
		ID:      uuid.NewString(),
		Title:   data.Title,
		Content: data.Content,
		UserId:  data.UserId,
		User:    &models.User{ID: data.UserId},
	}
	result := db.Create(&article)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": result.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &article})
}

func UpdateArticle(ctx *gin.Context) {
	var data updateArticleRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var article models.Article
	result := db.Model(&article).Take(&article, "ID = ?", data.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}
	article.Title = data.Title
	article.Content = data.Content
	saveResult := db.Save(&article)
	if saveResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": saveResult.Error})
	}
	response := articleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &response})
}

func DeleteArticle(ctx *gin.Context) {
	var data getArticle
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var article models.Article
	result := db.Model(&article).Take(&article, "iD = ?", data.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}
	deleteResult := db.Delete(&article)
	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": deleteResult.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &article})
}
