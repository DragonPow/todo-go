package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()

	BuildRouteTask(server.Group("/task"))
	// BuildRouteTag(server.Group("/tag"))
	// BuildRouteUser(server.Group("/user"))
	CreateRoute(server)

	return server
}