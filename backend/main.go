package main

import (
	c "sgstory/controllers"
	m "sgstory/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(m.CORSMiddleware())

	router.POST("/sendOTP", c.SendOTPCon)
	router.POST("/verifyOTP", c.VerifyOTP)
	router.POST("/stats", c.Stats)

	router.Run(":80")
}
