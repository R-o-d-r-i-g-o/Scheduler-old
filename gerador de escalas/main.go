package main

import (
	"scheduler/repository"
	"scheduler/routes"

	"github.com/gin-gonic/gin"
)

func init() {

	repository.ConnectToDataBase()
}

func main() {

	router := gin.Default()

	routes.Avaible(router)

	router.Run("localhost:8000")
}
