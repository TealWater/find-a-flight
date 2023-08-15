package main

import (
	control "find-a-flight/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/getFlights", control.Get_flights)
	router.GET("/getFares", control.Get_fares)
	router.GET("/getData", control.Populate)
	router.GET("/search", control.Search)
	router.Run(os.Getenv("GO_PORT"))
}
