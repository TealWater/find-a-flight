package json_files

import (
	"encoding/json"
	util "find-a-flight/utils"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type airline_iata []struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsLowcost bool   `json:"is_lowcost"`
	Logo      string `json:"logo"`
}

var airline_codes = &airline_iata{}
var Airline_map = make(map[string]string)

func init() {
	loadAirlines(&gin.Context{}, Airline_map)
	log.Println("Airlines loaded!")
}

func loadAirlines(c *gin.Context, airline_map map[string]string) {
	client := &http.Client{}
	url := "https://cdn.jsdelivr.net/gh/besrourms/airlines@latest/airlines.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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

	err = json.Unmarshal(body, &airline_codes)
	if err != nil {
		defer util.SafeExit("cdn for Commerical Airline codes are down \r\n Airline iata codes will be on display instead.")
		log.Panic(err)
	}

	for _, v := range *airline_codes {
		airline_map[v.Code] = v.Name
	}
}

func ValidateAirlineName(flight *util.AirlineData) {
	if val, ok := Airline_map[flight.Carrier_iata]; ok {
		flight.Carrier_iata = val
	}
}
