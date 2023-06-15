package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateListingsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "POST Opening",
	})
}