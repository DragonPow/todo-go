package router

import (
	"github.com/gin-gonic/gin"

	repository "project1/repository/postgresql"
	usecase "project1/usecase"
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
