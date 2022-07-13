package api_handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type any = interface{}
type jsonResposne = map[string]any

func BadRequesResponse(c *gin.Context, message string, args ...jsonResposne) {
	response := jsonResposne{"message": message}
	if len(args) > 0 {
		for _, v := range args {
			response = concat2Map(response, v)
		}
	}
	c.JSON(http.StatusBadRequest, response)
}

func Unauthorized(c *gin.Context, message string, args ...jsonResposne) {
	response := jsonResposne{"message": message}
	if len(args) > 0 {
		for _, v := range args {
			response = concat2Map(response, v)
		}
	}
	c.JSON(http.StatusUnauthorized, response)
}

func NotFoundResponse(c *gin.Context, message string, args ...jsonResposne) {
	response := jsonResposne{"message": message}
	if len(args) > 0 {
		for _, v := range args {
			response = concat2Map(response, v)
		}
	}
	c.JSON(http.StatusNotFound, response)
}

func SuccessResponse(c *gin.Context, object any, args ...jsonResposne) {
	response := jsonResposne{"data": object}
	if len(args) > 0 {
		for _, v := range args {
			response = concat2Map(response, v)
		}
	}
	c.JSON(http.StatusOK, response)
}

func ServerErrorResponse(c *gin.Context, args ...jsonResposne) {
	response := jsonResposne{"message": "Server error"}
	if len(args) > 0 {
		for _, v := range args {
			response = concat2Map(response, v)
		}
	}
	c.JSON(http.StatusInternalServerError, response)
}

func concat2Map(map1 jsonResposne, map2 jsonResposne) jsonResposne {
	for key, value := range map2 {
		map1[key] = value
	}
	return map1
}
