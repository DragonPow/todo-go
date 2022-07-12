package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	repository "project1/repository/postgresql"
	usecase "project1/usecase"
	"project1/util/api_handle"
	db "project1/util/db"
)

func Init(server *gin.Engine, DB db.Database) {
	taskRepo := repository.NewTaskRepository(DB)
	tagRepo := repository.NewTagRepository(DB)
	userRepo := repository.NewUserRepository(DB)

	taskUsecase := usecase.NewTaskUsecase(DB, taskRepo, userRepo)
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
