package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var version = "dev-build"

// SetupRouter creates the Gin.Engine with its routers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// add version endpoint
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "version:"+version)
		return
	})

	return r
}
