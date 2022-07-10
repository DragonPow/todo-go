package main

import (
	"project1/router"
	"project1/util/db"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()

	server := gin.Default()
	router.Init(server, *DB)

	server.Run(":8080")
}
