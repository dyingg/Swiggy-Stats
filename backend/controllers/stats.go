package controllers

import (
	orders "sgstory/orders/compute"

	"github.com/gin-gonic/gin"
)

type statsRequest struct {
	AuthToken string `json:"authToken" binding:"required"`
}

func Stats(c *gin.Context) {

	var statsRequest statsRequest

	if err := c.ShouldBindJSON(&statsRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userStats, err := orders.ComputeStats(statsRequest.AuthToken)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	UIStats := userStats.ToUIStats()

	c.JSON(200, gin.H{"stats": UIStats})
}
