package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/CatalogAPI/pkg/response"
)

func (r *Rest) GetAllCategory(c *gin.Context) {
	category, err := r.service.CategoryService.GetAllCategory()
	if err != nil {
		response.Failed(c, http.StatusInternalServerError, "Failed to get category", err.Error(), 50001)
		return
	}
	response.Success(c, http.StatusOK, "get categories success", category,nil)
}
