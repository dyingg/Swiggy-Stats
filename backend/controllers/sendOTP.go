package controllers

import (
	"sgstory/requests"

	"github.com/gin-gonic/gin"
)

type SendOTP struct {
	Number string `json:"number" binding:"required"`
}

func SendOTPCon(c *gin.Context) {
	var sendOTP SendOTP
	if err := c.ShouldBindJSON(&sendOTP); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tid, sid, err := requests.SendOTP(sendOTP.Number)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"tid": tid, "sid": sid})
}
