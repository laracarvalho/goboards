package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/schema"
)

func CreateListingHandler(ctx *gin.Context) {
	req := CreateListingRequest{}
	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	list := schema.Listing{
		Role:        req.Role,
		Company:     req.Company,
		Location:    req.Location,
		Remote:      *req.Remote,
		Link:        req.Link,
		Salary:      req.Salary,
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

func DeleteListingHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	listing := schema.Listing{}
	if err := db.First(&listing, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("listing with id: %s not found", id))
		return
	}

	if err := db.Delete(&listing).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting listing with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-listing", listing)
}

func ShowListingsHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	listing := schema.Listing{}
	if err := db.First(&listing, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "listing not found")
		return
	}

	sendSuccess(ctx, "show-listing", listing)
}

func UpdateListingsHandler(ctx *gin.Context) {
	req := UpdateListingRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	listing := schema.Listing{}

	if err := db.First(&listing, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "listing not found")
		return
	}

	if req.Role != "" {
		listing.Role = req.Role
	}

	if req.Company != "" {
		listing.Company = req.Company
	}

	if req.Location != "" {
		listing.Location = req.Location
	}

	if req.Remote != nil {
		listing.Remote = *req.Remote
	}

	if req.Link != "" {
		listing.Link = req.Link
	}

	if req.Salary > 0 {
		listing.Salary = req.Salary
	}

	if err := db.Save(&listing).Error; err != nil {
		logger.Errorf("error updating listing: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating listing")
		return
	}

	sendSuccess(ctx, "update-listing", listing)
}
