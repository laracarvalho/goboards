package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/handler"
)

func Start() {
	handler.InitHandler()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Project reachable.",
		})
	})

	router.POST("/listing", handler.CreateListingHandler)
	router.GET("/listings", handler.ListListingsHandler)
	router.DELETE("/listing", handler.DeleteListingHandler)
	router.GET("/listing", handler.ShowListingsHandler)


	router.Run(":8080")
}