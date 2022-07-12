package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"project1/domain"
	repository "project1/repository/postgresql"
	usecase "project1/usecase"
	"project1/util/api_handle"
	db "project1/util/db"
)

type IdUri struct {
	ID int32 `uri:"id"`
}

type jsonResponse map[string]interface{}

func Init(server *gin.Engine, DB db.Database) {
	taskRepo := repository.NewTaskRepository(DB)
	tagRepo := repository.NewTagRepository(DB)
	userRepo := repository.NewUserRepository(DB)

	taskUsecase := usecase.NewTaskUsecase(DB, taskRepo, userRepo, tagRepo)
	tagUsecase := usecase.NewTagUsecase(DB, tagRepo)
	userUsecase := usecase.NewUserUsecase(DB, userRepo)

	BuildTaskRoute(server.Group("/tasks"), taskUsecase)
	BuildTagRoute(server.Group("/tags"), tagUsecase)
	BuildUserRoute(server.Group("/users"), userUsecase)

}

func handlePanic(c *gin.Context, routeName string) {
	if r := recover(); r != nil {
		fmt.Printf("Panic in route: %s\n", routeName)
		fmt.Println(r)
		api_handle.ServerErrorResponse(c)
	}
}

func getUserJsonResponse(user domain.User) jsonResponse {
	return map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
	}
}

func getTagJsonResponse(tag domain.Tag) jsonResponse {
	return map[string]interface{}{
		"id":          tag.ID,
		"value":       tag.Value,
		"description": tag.Description,
	}
}

func getTaskJsonResponse(task domain.Task) jsonResponse {
	tags := []jsonResponse{}
	for _, tag := range task.Tags {
		tags = append(tags, getTagJsonResponse(tag))
	}

	res := map[string]interface{}{
		"id":           task.ID,
		"name":         task.Name,
		"description":  task.Description,
		"is_done":      task.IsDone,
		"user_creator": getUserJsonResponse(task.UserCreator),
		"tags":         tags,
	}

	return res
}
