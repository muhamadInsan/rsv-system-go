package controllers

import (
	"fmt"
	"net/http"
	"rsv-system-go/business"
	"rsv-system-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func FetchUsers(c *gin.Context) {
	var users []models.User
	if err := business.FetchUsers(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := business.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func FetchUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	var user models.User

	validate := validator.New()
	if err := validate.Struct(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data is not available!"})
		return
	}

	if err := business.FetchUserById(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	validate := validator.New()
	if err := validate.Struct(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data is not available!"})
		return
	}

	var user models.User
	c.BindJSON(&user)

	if err = business.UpdateUser(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	validate := validator.New()
	if err := validate.Struct(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data is not available!"})
		return
	}

	var user models.User
	if err = business.DeleteUser(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

func Login(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)

	token, err := business.Login(&user)
	if err != nil {
		fmt.Println("error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": token,
	})
}
