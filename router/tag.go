package router

import (
	"github.com/gin-gonic/gin"

	"project1/domain"
)

type tagRoute struct {
	tagUsecase domain.TagUsecase
}

func BuildTagRoute(router *gin.RouterGroup, t domain.TagUsecase) {
	tag := tagRoute{tagUsecase: t}

	router.GET("/", tag.Fetch)
	router.POST("/", tag.Create)
	router.DELETE("/:id", tag.Delete)
}

func (t *tagRoute) Fetch(c *gin.Context) {

}

func (t *tagRoute) Create(c *gin.Context) {

}

func (t *tagRoute) Delete(c *gin.Context) {

}
