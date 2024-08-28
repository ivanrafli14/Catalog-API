package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/ivanrafli14/CatalogAPI/pkg/response"
	"github.com/ivanrafli14/CatalogAPI/pkg/validation"
)

func (r *Rest) CreateProduct(c *gin.Context) {
	var productReq entity.ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {

		var val validator.ValidationErrors
		errorList, errorCode := validation.GetError(err, val)
		response.Failed(c, http.StatusBadRequest, "bad request", errorList, errorCode)
		return
	}

	_, err := r.service.ProductService.CreateProduct(&productReq)

	if err != nil {
		response.Failed(c, http.StatusInternalServerError, "Failed to create product", err.Error(), 50001)
		return
	}

	

	response.Success(c, http.StatusCreated, "create product success", nil,nil)
}

func (r *Rest) GetProduct(c *gin.Context) {
	name := c.Query("name")
	product,productTotal, err := r.service.ProductService.GetProducts(name)

	if err != nil {
		response.Failed(c, http.StatusInternalServerError, "Failed to get product", err.Error(), 50001)
		return
	}
	result := map[string]interface{}{
		"query": name,
		"total" : productTotal,
	}
	response.SuccessSearch(c, http.StatusOK, "Product has been retrieved", product, result)
}

func (r *Rest) UploadProductImage(c *gin.Context) {
	var photoReq entity.ProductPhotoRequest
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed(c, http.StatusBadRequest, "Failed to get photo product", err.Error(), 40001)
		return
	}

	// Manually populate the Photo struct
	photoReq.File = file

	if err := c.ShouldBind(&photoReq); err != nil {
		var val validator.ValidationErrors
		errorList, errorCode := validation.GetError(err, val)
		response.Failed(c, http.StatusBadRequest, "bad request", errorList, errorCode)
		return
	}

	if photoReq.Type != "PRODUCT" {
		response.Failed(c, http.StatusBadRequest, "bad request", "Type must be PRODUCT", 40003)
		return
	}

	// Manually populate the Photo struct
	id, ok := c.Get("user_id")
	if !ok {
		response.Failed(c, http.StatusUnauthorized, "Failed to get user", "User not found", 40101)
		return
	}
	result, err := r.service.ProductService.UploadPhoto(file, id.(string))

	if err != nil {
		response.Failed(c, http.StatusInternalServerError, "Failed to upload photo", err.Error(), 50001)
		return
	}
	response.Success(c, http.StatusOK, "Photo has been uploaded", result,nil)
}
