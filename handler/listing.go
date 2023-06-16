package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/schema"
)

func CreateListingsHandler(ctx *gin.Context) {
	req := CreateListingRequest{}
	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	list := schema.Listing{
		Role: req.Role,
		Company: req.Company,
		Location: req.Location,
		Remote: *req.Remote,
		Link: req.Link,
		Salary: req.Salary,
		Description: req.Description,
	}

	if err := db.Create(&list).Error; err != nil {
		logger.Errorf("Error creating item: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating listing on database")
		return
	}

	sendSuccess(ctx, "create-listing", list)
}

func ListListingsHandler(ctx *gin.Context) {
	listing := []schema.Listing{}

	if err := db.Find(&listing).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error returning Listing")
		return
	}

	sendSuccess(ctx, "list-Listing", listing)
}