package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerPing method to handler the ping endpoint
func HandlerPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
