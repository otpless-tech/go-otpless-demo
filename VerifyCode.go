package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"net/http"
)

func VerifyCode(c *gin.Context, config Config) {
	code := c.PostForm("code")

	result, err := otplessAuthSdk.VerifyCode(code, config.ClientID, config.ClientSecret)
	fmt.Printf("Result: for [%s] %v , error : %v\n", code, result, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
