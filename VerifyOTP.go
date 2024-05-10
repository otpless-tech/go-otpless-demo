package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"net/http"
)

func VerifyOTP(c *gin.Context, config Config) {
	orderID := c.PostForm("orderID")
	otp := c.PostForm("otp")
	email := c.PostForm("email")
	phoneNumber := c.PostForm("phoneNumber")
	result, err := otplessAuthSdk.VerifyOTP(orderID, otp, email, phoneNumber, config.ClientID, config.ClientSecret)
	fmt.Printf("Result:  %v , error : %v\n", result, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
