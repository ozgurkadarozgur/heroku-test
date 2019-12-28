package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   ComputeData(2, 3),
		})
	})
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func ComputeData(a, b int) int {
	return a * b
}
