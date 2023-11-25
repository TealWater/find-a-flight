package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	util "find-a-flight/utils"
	iata_codes "find-a-flight/utils/json_files"

	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

var wg sync.WaitGroup
var mu sync.Mutex
var list_of_flights []util.AirlineData
var price_data = &util.SkyscannerResults{}
var airport string

func Get_flights(c *gin.Context) {
	defer wg.Done()
	client := &http.Client{}
	flights := &util.FlightAwareFlight{}
	mu.Lock()
	list_of_flights = make([]util.AirlineData, 0)

	url := "https://aeroapi.flightaware.com/aeroapi/airports/" + airport + "/flights/scheduled_departures?type=Airline"
	time_start, time_end := util.Convert_to_iso(1)
	start := "start=" + time_start
	end := "end=" + time_end
	url += "&" + start + "&" + end

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &flights)
	if err != nil {
		log.Println()
		log.Println("can't parse json")
		log.Fatal(err)
	}

	departures := flights.ScheduledDepartures
	for _, v := range departures {
		list_of_flights = append(list_of_flights, util.AirlineData{
			Departure_time:      v.ScheduledOut.String(),
			Carrier_iata:        v.OperatorIata,
			Flight_Number:       v.FlightNumber,
			Source_Airport:      v.Origin.CodeIata,
			Destination_Airport: v.Destination.CodeIata,
		})
	}
	mu.Unlock()
}

func Get_fares(c *gin.Context) {
	defer wg.Done()
	mu.Lock()
	data := util.Build_skyscanner_data(list_of_flights)
	for i := 0; i < len(list_of_flights); i++ {
		jsonBytes, err := json.Marshal(data[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("sorry")
		}

		client := &http.Client{}
		url := "https://partners.api.skyscanner.net/apiservices/v3/flights/live/search/create"
		c.SetAccepted("application/json; charset=UTF-8")

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Println("unable to create request")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.Header = http.Header{
			"content-type": {"application/json; charset=UTF-8"},
			"x-api-key":    {os.Getenv("SKY_SCANNER_KEY")},
		}

		res, err := client.Do(req)
		if err != nil {
			log.Println("unable to send request")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("unable to read response body")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if err := json.Unmarshal(body, &price_data); err != nil {
			log.Println("unable to unmarshal json")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		json_data, err := gabs.ParseJSON(body)
		if err != nil {
			log.Println("unable to parse json")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		iata_codes.ValidateAirlineName(&list_of_flights[i])
		iata_codes.ValidateAirportName(&list_of_flights[i])

		sorting_options := []string{"best", "cheapest", "fastest"}
		for _, option := range sorting_options {
			booking, price := util.GetBookinglink_And_Price(json_data, option, *price_data)
			if booking == "" {
				log.Println("no flights")

				//Remove flight data from slice in-place
				list_of_flights = append(list_of_flights[:i], list_of_flights[i+1:]...)
				data = append(data[:i], data[i+1:]...)
				i--
				break
			}

			switch option {
			case "best":
				list_of_flights[i].Best_Flight = booking
				list_of_flights[i].Best_Flight_Price = price
			case "cheapest":
				list_of_flights[i].Cheapest_Flight = booking
				list_of_flights[i].Cheapest_Flight_Price = price
			case "fastest":
				list_of_flights[i].Fastest_Flight = booking
				list_of_flights[i].Fastest_Flight_Price = price
			}
			log.Println("flight found")
		}
	}
	c.JSON(http.StatusOK, list_of_flights)
	mu.Unlock()
}

func Search(c *gin.Context) {
	airport = c.DefaultQuery("airport", "KJFK")
	log.Println(airport)
	Populate(c)
}

func Populate(c *gin.Context) {
	enableCors(c)
	wg.Add(1)
	go Get_flights(c)
	wg.Wait()
	wg.Add(1)
	go Get_fares(c)
	wg.Wait()
}

func enableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", os.Getenv("TRUSTED_URL"))
}
