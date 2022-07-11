package router

import (
	"github.com/gin-gonic/gin"

	"project1/domain"
)

type taskRoute struct {
	taskUsecase domain.TaskUsecase
}

func BuildTaskRoute(router *gin.RouterGroup, t domain.TaskUsecase) {
	task := taskRoute{taskUsecase: t}

	router.GET("/", task.Fetch)
	router.GET("/:id", task.GetByID)
	router.POST("/", task.Create)
	router.PUT("/:id", task.Update)
	router.DELETE("/:id", task.Delete)
}

func (t *taskRoute) Fetch(c *gin.Context) {

}

func (t *taskRoute) GetByID(c *gin.Context) {

}

func (t *taskRoute) Create(c *gin.Context) {

}

func (t *taskRoute) Delete(c *gin.Context) {

}

func (t *taskRoute) Update(c *gin.Context) {

}
