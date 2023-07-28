package controllers

import (
	"TwitterApi/models"
	"TwitterApi/utils"
	"github.com/gin-gonic/gin"
)
func DeleteTweet(c *gin.Context) {
	var request models.UnTweet
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}
	resp, err := utils.Gql("POST", "DeleteTweet", map[string]interface{}{"tweet_id": request.TweetId, "dark_request": false}, c)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    resp,
	})
}
