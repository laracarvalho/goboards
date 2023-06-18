package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/schema"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendCreatedSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func sendPaginatedSuccess(ctx *gin.Context, op string, data interface{}, pagination interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("operation from handler: %s successfull", op),
		"data":       data,
		"pagination": pagination,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateListingResponse struct {
	Message string                 `json:"message"`
	Data    schema.ListingResponse `json:"data"`
}

type UpdateListingResponse struct {
	Message string                 `json:"message"`
	Data    schema.ListingResponse `json:"data"`
}

type DeleteListingResponse struct {
	Message string                 `json:"message"`
	Data    schema.ListingResponse `json:"data"`
}

type ListListingResponse struct {
	Message    string                   `json:"message"`
	Data       []schema.ListingResponse `json:"data"`
	Pagination Pagination               `json:"pagination"`
}

type ShowListingResponse struct {
	Message string                 `json:"message"`
	Data    schema.ListingResponse `json:"data"`
}

type Pagination struct {
	CurrentPage  int `json:"currentPage"`
	PageSize     int `json:"pageSize"`
	NextPage     int `json:"nextPage"`
	PreviousPage int `json:"previousPage"`
}
