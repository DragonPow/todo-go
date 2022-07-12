package router

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"project1/domain"
	"project1/util/api_handle"
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
	router.DELETE("/", task.DeleteAll)
	router.DELETE("/ids", task.Delete)
}

func (t *taskRoute) Fetch(c *gin.Context) {

}

func (t *taskRoute) GetByID(c *gin.Context) {
	defer handlePanic(c, "Get task by ID")

	// Get ID from uri
	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "ID must be a integer")
		return
	}

	// Get
	task, err := t.taskUsecase.GetByID(c.Request.Context(), id.ID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.NotFoundResponse(c, "The task does not exist, id "+strconv.Itoa(int(task.ID)))
		} else {
			api_handle.ServerErrorResponse(c)
		}

		return
	}

	api_handle.SuccessResponse(c, getTaskJsonResponse(task))
}

func (t *taskRoute) Create(c *gin.Context) {
	defer handlePanic(c, "Create task")

	// Parse token to user information
	var creator_id int32 = 1

	// Get new task information
	if _, err := c.MultipartForm(); err != nil {
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	// Define new_task_info
	name := c.PostForm("name")
	description := c.PostForm("description")
	is_done, _ := strconv.ParseBool(c.PostForm("is_done"))
	list_tag := []domain.Tag{}
	for _, v := range c.PostFormArray("tags_id") {
		value_int, err := strconv.Atoi(v)
		if err != nil {
			api_handle.BadRequesResponse(c, "Tags is must be an integer")
			return
		}
		list_tag = append(list_tag, domain.Tag{ID: int32(value_int)})
	}

	new_task := domain.Task{
		Name:        name,
		Description: description,
		IsDone:      is_done,
		Tags:        list_tag,
		UserCreator: domain.User{ID: creator_id},
	}

	// Create
	new_task, err := t.taskUsecase.Create(c.Request.Context(), creator_id, new_task)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.NotFoundResponse(c, "The user does not exist")
		} else if errors.Is(err, domain.ErrTagNotExists) {
			api_handle.NotFoundResponse(c, "The tag does not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}

		return
	}

	new_task, err = t.taskUsecase.GetByID(c.Request.Context(), new_task.ID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.NotFoundResponse(c, "The task does not exist, id "+strconv.Itoa(int(new_task.ID)))
		} else {
			api_handle.ServerErrorResponse(c)
		}

		return
	}

	api_handle.SuccessResponse(c, new_task)
}

func (t *taskRoute) DeleteAll(c *gin.Context) {

}

func (t *taskRoute) Delete(c *gin.Context) {
	// Get ID from uri
	var ids IdsUri
	if err := c.ShouldBind(&ids); err != nil {
		api_handle.BadRequesResponse(c, "ID must be a integer")
		return
	}

	if err := t.taskUsecase.Delete(c.Request.Context(), ids.IDs); err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.BadRequesResponse(c, "Task does not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, "Delete id success")
}

func (t *taskRoute) Update(c *gin.Context) {

}
