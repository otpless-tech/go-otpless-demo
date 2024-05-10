package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"net/http"
)

func ReSendOTP(c *gin.Context, config Config) {
	orderId := c.PostForm("order_id")
	result, err := otplessAuthSdk.ResendOTP(orderId, config.ClientID, config.ClientSecret)
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
