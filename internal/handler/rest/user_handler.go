package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/ivanrafli14/CatalogAPI/internal/service"
	"github.com/ivanrafli14/CatalogAPI/pkg/response"
	"github.com/ivanrafli14/CatalogAPI/pkg/validation"
)

func (r *Rest) Register(c *gin.Context) {
	var userReq entity.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		var val validator.ValidationErrors
		errorList, errorCode := validation.GetError(err, val)
		response.Failed(c, http.StatusBadRequest, "bad request", errorList, errorCode)
		return

	}

	_, err := r.service.UserService.Register(&userReq)

	if err != nil {
		if errors.Is(err, service.ErrDuplicateEmail) {
			response.Failed(c, http.StatusBadRequest, "bad request", "email has been used", 40002)
			return
		}
		response.Failed(c, http.StatusInternalServerError, "error repository", err.Error(), 50001)
		return
	}

	response.Success(c, http.StatusCreated, "registration success", nil,nil)
}

func (r *Rest) Login(c *gin.Context) {
	var userReq entity.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		var val validator.ValidationErrors
		errorList, errorCode := validation.GetError(err, val)
		response.Failed(c, http.StatusBadRequest, "bad request", errorList, errorCode)
		return
	}

	loginRes, err := r.service.UserService.Login(&userReq)

	if err != nil {
		response.Failed(c, http.StatusInternalServerError, "Failed to login", err.Error(), 50001)
		return
	}
	response.Success(c, http.StatusOK, "Login success", loginRes,nil)
}
