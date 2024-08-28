package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/CatalogAPI/pkg/response"
)

func (m *middleware) Authentication(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		response.Failed(c, http.StatusUnauthorized, "Invalid token. Please provide valid authentication credentials", "bearer toke is empty", 40101)
		c.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	_, err := m.redis.GetData(token)
	if err != nil {
		response.Failed(c, http.StatusUnauthorized, "Invalid token. Token is not store in redis. Please login", err.Error(), 40101)
		c.Abort()
		return
	}

	userId, err := m.jwtAuth.VerifyJWTToken(token)

	if err != nil {
		response.Failed(c, http.StatusUnauthorized, "Invalid token. Please provide valid authentication credentials", err.Error(), 40101)
		c.Abort()
		return
	}

	user, err := m.service.UserService.FindByID(userId)

	if err != nil {
		response.Failed(c, http.StatusUnauthorized, "Invalid ID", err.Error(), 40101)
		c.Abort()
		return
	}

	c.Set("user_id", user.ID)
	c.Next()

}
