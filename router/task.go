package router

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"project1/domain"
	"project1/util/api_handle"
	"project1/util/helper"
)

type taskRoute struct {
	taskUsecase domain.TaskUsecase
}

type fetchReqStruct struct {
	Name       string  `form:"name" json:"name"`
	Tags       []int32 `form:"tags" json:"tags"`
	StartIndex int32   `form:"startIndex" json:"startIndex"`
	Number     int32   `form:"number" json:"number"`
}

type TaskUpdateReqStruct struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	IsDone      bool    `form:"is_done:`
	TagsAdd     []int32 `form:"tags_add"`
	TagsDelete  []int32 `form:"tags_delete"`
}

func BuildTaskRoute(router *gin.RouterGroup, t domain.TaskUsecase) {
	task := taskRoute{taskUsecase: t}

	router.GET("/search", task.Fetch)
	router.GET("/:id", task.GetByID)
	router.POST("/", task.Create)
	router.PUT("/:id", task.Update)
	router.DELETE("/", task.DeleteAll)
	router.DELETE("/ids", task.Delete)
}

func (route *taskRoute) Fetch(c *gin.Context) {
	defer handlePanic(c, "Fetch task")

	creator_id, err := authenticateUser(c)
	if err != nil {
		return
	}

	var reqJson fetchReqStruct
	if err := c.ShouldBindQuery(&reqJson); err != nil {
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	// Define args
	args := make(map[string]interface{})
	if strings.TrimSpace(reqJson.Name) != "" {
		args["name"] = strings.TrimSpace(reqJson.Name)
	}
	// if reqJson.Tags != nil {
	// 	args["tags"] = reqJson.Tags
	// }

	tasks, err := route.taskUsecase.Fetch(c.Request.Context(), creator_id, reqJson.StartIndex, reqJson.Number, args)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.BadRequesResponse(c, "User is not exists")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, tasks)
}

func (route *taskRoute) GetByID(c *gin.Context) {
	defer handlePanic(c, "Get task by ID")

	// Get ID from uri
	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "ID must be a integer")
		return
	}

	// Get
	task, err := route.taskUsecase.GetByID(c.Request.Context(), id.ID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.NotFoundResponse(c, "The task does not exist, id "+strconv.Itoa(int(id.ID)))
		} else {
			api_handle.ServerErrorResponse(c)
		}

		return
	}

	api_handle.SuccessResponse(c, getTaskJsonResponse(task))
}

func (route *taskRoute) Create(c *gin.Context) {
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

	// Tranfer to domain
	new_task := domain.Task{
		Name:        name,
		Description: description,
		IsDone:      is_done,
		Tags:        list_tag,
		UserCreator: domain.User{ID: creator_id},
	}

	// Create
	new_task, err := route.taskUsecase.Create(c.Request.Context(), creator_id, new_task)
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

	// Get
	new_task, err = route.taskUsecase.GetByID(c.Request.Context(), new_task.ID)
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

func (route *taskRoute) DeleteAll(c *gin.Context) {
	// Get creator_id
	creator_id, err := authenticateUser(c)
	if err != nil {
		return
	}

	// Delete
	if err := route.taskUsecase.DeleteAll(c.Request.Context(), creator_id); err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.BadRequesResponse(c, "User is not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, "Delete all task of user "+strconv.Itoa(int(creator_id))+" success")
}

func (route *taskRoute) Delete(c *gin.Context) {
	// Get ID from body
	var ids IdsUri
	if err := c.ShouldBind(&ids); err != nil {
		api_handle.BadRequesResponse(c, "ID must be a integer")
		return
	}

	// Delete
	if err := route.taskUsecase.Delete(c.Request.Context(), ids.IDs); err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.BadRequesResponse(c, "Task does not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, "Delete ids: ["+strings.Join(helper.IntToString(ids.IDs), ",")+"] success")
}

func (route *taskRoute) Update(c *gin.Context) {
	// Get ID from uri
	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "Id must be a integer")
		return
	}

	// Get task information
	var taskReq TaskUpdateReqStruct
	if err := c.ShouldBind(&taskReq); err != nil {
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	// Tranfer
	var new_task_info map[string]interface{}
	var new_tags_add []int32
	var new_tags_remove []int32
	// TODO: Implement code here

	// Update
	if err := route.taskUsecase.Update(c.Request.Context(), id.ID, new_task_info, new_tags_add, new_tags_remove); err != nil {
		if errors.Is(err, domain.ErrTaskNotExists) {
			api_handle.NotFoundResponse(c, "Task with ID "+strconv.Itoa(int(id.ID))+" does not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, nil)
}
