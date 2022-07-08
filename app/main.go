package main

import (
	"fmt"
	"log"
	"time"
	
	"github.com/gin-gonic/gin"
	"net/http"

	"db_setup" db
)

func main()  {
	DB := db.Init()
	defer DB.Close()
	
	server := gin.Default()

	server.Run(":8080")
}