package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var globalConfig Config

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	loadConfiguration()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.html", nil)
	})

	router.POST("/generate-magic-link", func(c *gin.Context) {
		GenerateMagicLink(c, globalConfig)
	})

	router.POST("/verify-code", func(c *gin.Context) {
		VerifyCode(c, globalConfig)
	})

	router.POST("/verify-auth-token", func(c *gin.Context) {
		VerifyAuthToken(c, globalConfig)
	})

	router.POST("/send-otp", func(c *gin.Context) {
		SendOTP(c, globalConfig)
	})

	router.POST("/re-send-otp", func(c *gin.Context) {
		ReSendOTP(c, globalConfig)
	})

	router.POST("/verify-otp", func(c *gin.Context) {
		ReSendOTP(c, globalConfig)
	})

	_ = router.Run(":9098")
}

func loadConfiguration() {

	globalConfig = Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	if globalConfig.ClientID == "" || globalConfig.ClientSecret == "" {
		log.Fatal("Missing configuration. Please set CLIENT_ID, CLIENT_SECRET")
	}
}
