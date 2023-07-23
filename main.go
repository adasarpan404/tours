package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adasarpan404/tours/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client mongo.Client

func loadEnvFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	err := loadEnvFromFile("config.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("API Server for touring")
	router := gin.New()
	Client = service.ConnectMongoDB()
	router.Run(":8080")
}
