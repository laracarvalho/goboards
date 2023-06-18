package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/config"
	"github.com/laracarvalho/goboards/schema"
)

// @BasePath /api/v1

// @Summary Create listing
// @Description Create a new job listing
// @Tags Listings
// @Accept json
// @Produce json
// @Param request body CreateListingRequest true "Request body"
// @Success 201 {object} CreateListingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /listing [post]
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

	sendCreatedSuccess(ctx, "create-listing", list)
}

// @Summary List listings
// @Description Returns an array of listing
// @Tags Listings
// @Accept json
// @Produce json
// @Success 200 {object} ListListingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /listings [get]
func ListListingsHandler(ctx *gin.Context) {
	page := ctx.Query("page")
	page_size := ctx.Query("page_size")

	listing := []schema.Listing{}

	if err := db.Scopes(config.Paginate(page, page_size)).Find(&listing).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error returning Listing")
		return
	}

	p, _ := strconv.Atoi(page)
	ps, _ := strconv.Atoi(page_size)

	pag := genPagination(p, ps)

	sendPaginatedSuccess(ctx, "list-Listing", listing, pag)
}

// @Summary Delete listings
// @Description Deletes a listing
// @Tags Listings
// @Accept json
// @Produce json
// @Param id query string true "Listing identification"
// @Success 200 {object} DeleteListingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /listing [delete]
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

// @Summary Show listings
// @Description Show a single listing
// @Tags Listings
// @Accept json
// @Produce json
// @Param id query string true "Listing identification"
// @Param opening body UpdateListingRequest true "Listing data to Update"
// @Success 200 {object} ShowListingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /listing [get]
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

// @Summary Update listings
// @Description Updates a listing
// @Tags Listings
// @Accept json
// @Produce json
// @Param id query string true "Listing identification"
// @Success 200 {object} UpdateListingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /listing [put]
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
