package router

import (
	repository "project1/repository/postgresql"
	"project1/usecase"
	"project1/util/db"

	"github.com/gin-gonic/gin"
)

func Init(server *gin.Engine, DB db.Database) {
	taskRepo := repository.NewTaskRepository(DB)
	tagRepo := repository.NewTagRepository(DB)
	userRepo := repository.NewUserRepository(DB)

	taskUsecase := usecase.NewTaskUsecase(taskRepo, userRepo)
	tagUsecase := usecase.NewTagUsecase(tagRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	BuildTaskRoute(server.Group("/tasks"), taskUsecase)
	BuildTagRoute(server.Group("/tags"), tagUsecase)
	BuildUserRoute(server.Group("/users"), userUsecase)

}
