package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	util "find-a-flight/utils"

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

var list_of_flights []util.AirlineData
var price_data = &util.SkyscannerResults{}

func Get_flights(c *gin.Context) {
	client := &http.Client{}
	flights := &util.FlightAwareFlight{}
	list_of_flights = make([]util.AirlineData, 0)

	url := "https://aeroapi.flightaware.com/aeroapi/airports/KJFK/flights/scheduled_departures?type=Airline"
	time_start, time_end := util.Convert_to_iso(1)
	start := "start=" + time_start
	end := "end=" + time_end
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

	for _, v := range list_of_flights {
		c.String(http.StatusOK, v.Flight_Number+"\n")
	}
	c.String(http.StatusOK, "done!")
}

func Get_fares(c *gin.Context) {

	data := util.Build_skyscanner_data(list_of_flights)
	for i := 0; i < len(list_of_flights); i++ {
		jsonBytes, err := json.Marshal(data[i])
		/* mock data
		SS := util.SkyscannerData{
			Query: struct {
				Market    string "json:\"market\""
				Locale    string "json:\"locale\""
				Currency  string "json:\"currency\""
				QueryLegs []struct {
					OriginPlaceID struct {
						Iata string "json:\"iata\""
					} "json:\"originPlaceId\""
					DestinationPlaceID struct {
						Iata string "json:\"iata\""
					} "json:\"destinationPlaceId\""
					Date struct {
						Year  int "json:\"year\""
						Month int "json:\"month\""
						Day   int "json:\"day\""
					} "json:\"date\""
				} "json:\"queryLegs\""
				CabinClass          string   "json:\"cabinClass\""
				Adults              int      "json:\"adults\""
				IncludedCarriersIds []string "json:\"includedCarriersIds\""
			}{},
		}

		SS.Query.Market = "US"
		SS.Query.Locale = "en-US"
		SS.Query.Currency = "USD"
		slice := SS.Query.QueryLegs
		slice = append(slice, struct {
			OriginPlaceID struct {
				Iata string "json:\"iata\""
				//EntityID string "json:\"entityId\""
			} "json:\"originPlaceId\""
			DestinationPlaceID struct {
				Iata string "json:\"iata\""
				//EntityID string "json:\"entityId\""
			} "json:\"destinationPlaceId\""
			Date struct {
				Year  int "json:\"year\""
				Month int "json:\"month\""
				Day   int "json:\"day\""
			} "json:\"date\""
		}{})

		slice[0].OriginPlaceID.Iata = "JFK"
		slice[0].DestinationPlaceID.Iata = "LIM"
		slice[0].Date.Year = 2023
		slice[0].Date.Month = 4
		slice[0].Date.Day = 30
		SS.Query.QueryLegs = append(SS.Query.QueryLegs, slice...)
		SS.Query.CabinClass = "CABIN_CLASS_ECONOMY"
		SS.Query.Adults = 1
		SS.Query.IncludedCarriersIds = append(SS.Query.IncludedCarriersIds, "LA")

		jsonBytes, err := json.Marshal(SS)
		*/
		log.Println("111")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("sorry")
		}
		log.Println("222")

		client := &http.Client{}
		url := "https://partners.api.skyscanner.net/apiservices/v3/flights/live/search/create"
		c.SetAccepted("application/json; charset=UTF-8")

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.Header = http.Header{
			"content-type": {"application/json; charset=UTF-8"},
			"x-api-key":    {os.Getenv("SKY_SCANNER_KEY")},
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

		if err := json.Unmarshal(body, &price_data); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		json, err := gabs.ParseJSON(body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		sorting_options := []string{"best", "cheapest", "fastest"}
		for _, option := range sorting_options {
			booking := util.GetBookinglink(json, option, *price_data)
			if booking == "" {
				fmt.Println("no flights")
				continue
			}

			switch option {
			case "best":
				list_of_flights[i].Best_Flight = booking
			case "cheapest":
				list_of_flights[i].Cheapest_Flight = booking
			case "fastest":
				list_of_flights[i].Fastest_Flight = booking
			}
			fmt.Println("flight found")
			fmt.Println(booking)
		}
		c.String(http.StatusOK, "%+v\n", list_of_flights)
	}
}
