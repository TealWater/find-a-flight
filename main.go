package main

import (
	control "find-a-flight/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/getFlights", control.GetFlights)

	router.Run(":8080")
}
