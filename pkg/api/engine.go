package api

import "github.com/gin-gonic/gin"

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine() (*gin.Engine, error) {
	r := gin.New()

	return r, nil
}