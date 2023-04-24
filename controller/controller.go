package controller

import (
	"io/ioutil"
	"log"
	"net/http"
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

func GetFlights(c *gin.Context) {
	client := &http.Client{}
	url := "https://aeroapi.flightaware.com/aeroapi/airports/KJFK/flights/scheduled_departures?type=Airline"
	start := "start=2023-04-23T23:00:00Z"
	end := "end=2023-04-24T00:00:00Z"
	url += "&" + start + "&" + end

	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	req.Header = http.Header{
		"content-type": {"application/json; charset=UTF-8"},
		"x-apikey":     {os.Getenv("FLIGHT_AWARE_KEY")},
	}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, string(body))

	/*c.JSON(200, gin.H{
		"message": "Welcome to the home page",
	})*/
}

func GetFares(c *gin.Context) {

}
