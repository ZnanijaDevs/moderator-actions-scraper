package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/deletionReasons", func(c *gin.Context) {
		c.JSON(200, deletionReasons)
	})

	router.GET("/actions", func(ctx *gin.Context) {
		limit := ToNumber(ctx.DefaultQuery("limit", "15"))
		offsetPage := ToNumber(ctx.DefaultQuery("offsetPage", "1"))

		userIds, _ := ExtractUserIdsFromQuery(ctx.DefaultQuery("users", "1"))

		if (
			limit < 1 || 
			limit > 9000 || 
			offsetPage < 1 || 
			offsetPage > 1_000_000 ||
			userIds == nil) {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
			return
		}

		actions, errors := GetActions(userIds, limit, offsetPage)

		ctx.JSON(200, gin.H{
			"actions": actions,
			"errors": errors,
			"count": len(actions),
		})
	})

	router.Run()
}