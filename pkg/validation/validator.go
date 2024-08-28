package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetField(fe validator.FieldError) string {
	var builder strings.Builder

	for i, char := range fe.Field() {
		if i > 0 && char >= 'A' && char <= 'Z' {
			builder.WriteRune('_')
		}
		builder.WriteRune(char)
	}

	return strings.ToLower(builder.String())
}

func GetError(err error, ve validator.ValidationErrors) (string, int) {
	if errors.As(err, &ve) {
		for _, fe := range ve {
			// Return the first error message and code
			return GetErrorMsg(fe)
		}
	}

	return "unknown error", 40000
}

func GetErrorMsg(fe validator.FieldError) (string, int) {
	switch fe.Field() {
	case "Email" :
		if fe.Tag() == "required" {
			return "email is required", 40001
		}
		if fe.Tag() == "email" {
			return "invalid email", 40002
		}

	case "Password":
		if fe.Tag() == "required" {
			return "password is required", 40003
		}
		if fe.Tag() == "min" {
			return "invalid password", 40004
		}

	//product
	case "Name":
		if fe.Tag() == "required" {
			return "name is required", 40001
		}
	case "Stock":
		if fe.Tag() == "required" {
			return "stock is required", 40002
		}
	case "Price":
		if fe.Tag() == "required" {
			return "price is required", 40003
		}
	case "CategoryID":
		if fe.Tag() == "required" {
			return "category ID is required", 40004
		}
	case "ImageUrl":
		if fe.Tag() == "required" {
			return "imageUrl is required", 40005
		}

	case "File":
		if fe.Tag() == "required" {
			return "file is required", 40001
		}

	case "Type":
		if fe.Tag() == "required" {
			return "type is required", 40003
		}
	
	}


	return fe.Tag(), 40000
}

