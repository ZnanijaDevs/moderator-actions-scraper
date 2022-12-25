package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/deletionReasons", func(c *gin.Context) {
		c.JSON(200, deletionReasons)
	})

	router.GET("/actions/:userId", func(ctx *gin.Context) {
		limit := ToNumber(ctx.DefaultQuery("limit", "15"))
		offsetPage := ToNumber(ctx.DefaultQuery("offsetPage", "1"))
		userId := ToNumber(ctx.Param("userId"))
		authToken := ctx.Query("authToken")

		if (
			limit < 1 || 
			limit > 9000 || 
			offsetPage < 1 || 
			offsetPage > 1_000_000 ||
			authToken == "") {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
			return
		}

		actions, errors := GetActions(userId, limit, offsetPage, authToken)

		ctx.JSON(200, gin.H{
			"actions": actions,
			"errors": errors,
			"count": len(actions),
		})
	})

	router.Run()
}