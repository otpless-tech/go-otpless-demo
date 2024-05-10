package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
	"math/rand"
	"net/http"
	"strconv"
)

func SendOTP(c *gin.Context, config Config) {
	mobile := c.PostForm("mobile")
	email := c.PostForm("email")
	v := rand.Int()
	result, err := otplessAuthSdk.SendOTP(otplessAuthSdk.SendOTPRequest{
		PhoneNumber: mobile,
		Email:       email,
		Channel:     "SMS",
		Hash:        "GOSDK",
		OrderId:     "GO_SDK" + strconv.Itoa(v),
		Expiry:      60,
		OtpLength:   4,
	}, config.ClientID, config.ClientSecret)
	fmt.Printf("Result:  %v , error : %v\n", result, err)

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
