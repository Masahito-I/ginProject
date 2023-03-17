package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ginProject/models"
)

func GetUser(ctx *gin.Context) {
	// get user ID from path parameter
	id := ctx.Param("id")

	// get user data from database
	user, err := models.GetUser(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	// parse JSON request body into User struct
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// save user data to database
	if err := models.CreateUser(user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
