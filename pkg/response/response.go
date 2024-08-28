package response

import "github.com/gin-gonic/gin"

type Res struct {
	Success    bool   `json:"succes"`
	Message    string `json:"message"`
	Payload    any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}



func Success(c *gin.Context, code int, message string, data any, pagination any) {
	c.JSON(code, Res{
		Success:    true,
		Message:    message,
		Payload:    data,
		Pagination: pagination,
	})
}

func SuccessSearch(c *gin.Context, code int, message string, data any, pagination any) {
	c.JSON(code, gin.H{
		"success":    true,
		"message":    message,
		"payload":    data,
		"pagination": pagination,
	})
}

func Failed(c *gin.Context, code int, message string, error_msg any, error_code int) {
	c.JSON(code, gin.H{
		"success":    false,
		"message":    message,
		"error":      error_msg,
		"error_code": error_code,
	})
}
