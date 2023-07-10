package main

import (
	"fmt"

	"github.com/adasarpan404/tours/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API Server for touring")
	router := gin.New()
	service.ConnectMongoDB()
	router.Run(":8080")
}
