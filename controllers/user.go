package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/LLuthfiY/fake-api-GO/config"
	"github.com/LLuthfiY/fake-api-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type userGetUsername struct {
	Username string `json:"username" binding:"required"`
}

type userUpdateRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
}

type userCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetAllUser(ctx *gin.Context) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}
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

func GetOneUser(ctx *gin.Context) {
	var data userGetUsername
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var user = models.User{}
	user.Username = data.Username
	result := db.Model(&user).Take(&user, "username = ?", &data.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": result.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &user})
}

func UpdateUser(ctx *gin.Context) {
	var data userUpdateRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error})
		return
	}
	fmt.Println(err)
	var user models.User
	result := db.Model(&user).Take(&user, "username = ? ", &data.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": result.Error})
		return
	}
	user.Name = sql.NullString{String: data.Name, Valid: true}
	fmt.Println(&user)
	db.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": &user})
}

func DeleteUser(ctx *gin.Context) {
	var data userGetUsername
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var user models.User
	result := db.Model(&user).Take(&user, "username = ? ", &data.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": result.Error})
		return
	}
	db.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": &user})
}
