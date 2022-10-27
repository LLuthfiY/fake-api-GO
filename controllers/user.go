package controllers

import (
	"net/http"

	"github.com/LLuthfiY/fake-api-GO/config"
	"github.com/LLuthfiY/fake-api-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type userCreateResponse struct {
	Name     string
	Username string
	Password string
}

type userCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAllUser(ctx *gin.Context) {
	var users []models.User
	db.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": &users})
}

func CreateUser(ctx *gin.Context) {
	var data userCreateRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	newUUID := uuid.New()
	userModels := models.User{}
	userModels.Id = newUUID.String()
	userModels.Username = data.Username
	userModels.Password = data.Password
	user := db.Create(&userModels)
	if user.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": user.Error})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": &userModels})
}
