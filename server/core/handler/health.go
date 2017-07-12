package handler

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// Health handle /health route
func Health(c *gin.Context) {
	c.String(http.StatusOK)
}
