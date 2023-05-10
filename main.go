package main

import (
	control "find-a-flight/controller"
	util "find-a-flight/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		start, end := util.Convert_to_iso(4)
		c.JSON(200, gin.H{
			"message":      "Hello World",
			"start_time: ": start,
			"end_time: ":   end,
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.File("./frontend/src/routes/+page.svelte")
	})

	router.GET("/getFlights", control.Get_flights)
	router.GET("/getFares", control.Get_fares)
	router.GET("/getData", control.Populate)
	router.Run(":8080")
}
