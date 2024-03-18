package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	if(c.Request.Method == "POST"){
		c.JSON(http.StatusOK, gin.H{"message": "Register - POST"})
	}

	if(c.Request.Method == "GET"){
		c.JSON(http.StatusOK, gin.H{"message": "Register - GET"})
	}
}