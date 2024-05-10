package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"net/http"
)

func GenerateMagicLink(c *gin.Context, config Config) {
	mobileNumber := c.PostForm("mobileNumber")
	email := c.PostForm("email")
	redirectUri := "http://localhost"
	channel := "WHATSAPP" //SMS,EMAIL

	result, err := otplessAuthSdk.GenerateMagicLink(mobileNumber, email, config.ClientID, config.ClientSecret, redirectUri, channel)
	fmt.Printf("Result: for [%s] %v , error : %v\n", mobileNumber, result, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result,
			"error":   err.Error(), // Convert error to string
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
