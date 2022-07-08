package router

import (
	"project1/domain"
)

type taskRoute struct {
	usecase domain.TaskUsecase
}

func (t *taskRoute) BuildRouteTask(router *gin.Engine) {
	router.GET("/", t.Fetch)
	router.GET("/:id", t.GetTaskByID)
	router.POST("/", t.CreateTask)
	router.DELETE("/:id", t.DeleteTask)
	router.PUT("/:id", t.UpdateTask)
}