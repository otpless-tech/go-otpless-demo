package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"net/http"
)

func VerifyAuthToken(c *gin.Context, config Config) {
	authToken := c.PostForm("authToken")

	result, err := otplessAuthSdk.VerifyAuthToken(authToken, config.ClientID, config.ClientSecret)
	fmt.Printf("Result: for [%s] %v , error : %v\n", authToken, result, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
