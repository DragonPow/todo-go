package api_handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type any = interface{}
type jsonResposne = map[string]any

func BadRequesResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, jsonResposne{"message": message})
}

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, jsonResposne{"message": message})
}

func SuccessResponse(c *gin.Context, object any) {
	c.JSON(http.StatusOK, jsonResposne{"data": object})
}

func ServerErrorResponse(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, jsonResposne{"message": "Server error"})
}
