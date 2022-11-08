package controllers

import (
	"sgstory/requests"

	"github.com/gin-gonic/gin"
)

type verifyOTP struct {
	TID string `json:"tid" binding:"required"`
	SID string `json:"sid" binding:"required"`
	OTP string `json:"otp" binding:"required"`
}

func VerifyOTP(c *gin.Context) {

	var verifyOTP verifyOTP

	if err := c.ShouldBindJSON(&verifyOTP); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authToken, err := requests.VerifyOTP(verifyOTP.TID, verifyOTP.SID, verifyOTP.OTP)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//In the future do not want to expose this token to the user.

	c.JSON(200, gin.H{"authToken": authToken})
}
