package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API Server for touring")
	router := gin.New()

	router.Run(":8080")
}
