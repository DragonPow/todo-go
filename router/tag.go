package router

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"project1/domain"
	"project1/util/api_handle"
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

func (route *tagRoute) Fetch(c *gin.Context) {
	tags, err := route.tagUsecase.FetchAll(c.Request.Context())
	if err != nil {
		api_handle.ServerErrorResponse(c)
		return
	}

	api_handle.SuccessResponse(c, tags)
}

func (route *tagRoute) Create(c *gin.Context) {
	defer handlePanic(c, "Create tag")

	// Get tag information from body
	var tag_info domain.Tag
	if err := c.ShouldBind(&tag_info); err != nil {
		fmt.Println(err)
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	// Create
	new_tag, err := route.tagUsecase.Create(c.Request.Context(), tag_info)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.NotFoundResponse(c, "User is not exist")
		} else if errors.Is(err, domain.ErrTagValueDuplicated) {
			api_handle.BadRequesResponse(c, "Tag value already exists")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, new_tag)
}

func (route *tagRoute) Delete(c *gin.Context) {
	defer handlePanic(c, "Delete tag")

	// Get Id from uri
	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "ID must be integer")
		return
	}

	// Delete
	if err := route.tagUsecase.Delete(c, id.ID); err != nil {
		if errors.Is(err, domain.ErrTagNotExists) {
			api_handle.NotFoundResponse(c, "Tag with Id "+strconv.Itoa(int(id.ID))+" does not exist")
		} else if errors.Is(err, domain.ErrTagStillReference) {
			api_handle.BadRequesResponse(c, "Some task contains tag, please delete task before")
		} else {
			api_handle.ServerErrorResponse(c)
		}

		return
	}

	api_handle.SuccessResponse(c, "Delete tag success with id "+strconv.Itoa(int(id.ID)))
}
