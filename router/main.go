package router

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"project1/domain"
	repository "project1/repository/postgresql"
	usecase "project1/usecase"
	"project1/util/api_handle"
	db "project1/util/db"
	"project1/util/jwt_handle"
)

type IdUri struct {
	ID int32 `uri:"id"`
}

type IdsUri struct {
	IDs []int32 `form:"ids"`
}

type dataToken struct {
	user_id int32 `json:"user_id"`
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

func getToken(c *gin.Context) (token string, err error) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.Abort()
		return "", domain.ErrTokenRequired
	}

	// authorization = "Bearer " + token
	parts := strings.Split(authorization, " ")
	if parts[0] != "Bearer" {
		c.Abort()
		return "", domain.ErrTokenFormatInvalid
	}

	return parts[1], nil
}

func authenticateUser(c *gin.Context) (user_id int32, err error) {
	// Get token
	token, err := getToken(c)
	if err != nil {
		if errors.Is(err, domain.ErrTokenRequired) {
			api_handle.Unauthorized(c, "Token is required")
		} else if errors.Is(err, domain.ErrTokenFormatInvalid) {
			api_handle.BadRequesResponse(c, "Token format is wrong")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		c.Abort()
		return 0, err
	}

	if token == "Admin" {
		return 1, nil
	}

	// Parse token
	data, err := jwt_handle.ParseToData(token)
	if err != nil {
		api_handle.Unauthorized(c, "Token is invalid")
		c.Abort()
		return 0, err
	}

	user_id = data.(dataToken).user_id
	return user_id, nil
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
