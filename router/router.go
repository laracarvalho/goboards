package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/handler"
)

func Start() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Project reachable.",
		})
	})

	router.POST("/listings", handler.CreateListingsHandler)

	router.Run(":8080")
}