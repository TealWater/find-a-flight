package json_files

import (
	"encoding/json"
	util "find-a-flight/utils"
	"log"
	"os"

	"github.com/Jeffail/gabs"
)

var airport_codes = make(map[string]interface{})
var json_data = &gabs.Container{}

func init() {
	if err := loadAirports(); err == nil {
		log.Println("Airports loaded!")
	}
}

func loadAirports() error {
	defer util.SafeExit("Unable to load airport names")
	data, err := os.ReadFile("utils/json_files/airports.json")
	if err != nil {
		log.Panicln("Error reading file:", err)
		return err
	}

	err = json.Unmarshal(data, &airport_codes)
	if err != nil {
		log.Panic("Error parsing JSON: \n", err)
		return err
	}

	json_data, err = gabs.ParseJSON(data)
	if err != nil {
		log.Panic("unable to parse json. \n", err)
		return err
	}
	return nil
}

func ValidateAirportName(flight *util.AirlineData) {
	dest_iata := flight.Destination_Airport
	source_iata := flight.Source_Airport
	if _, ok := airport_codes[dest_iata]; ok {
		flight.Destination_Airport = json_data.Search(dest_iata).Search("name").String()
		flight.Destination_Country = json_data.Search(dest_iata).Search("country").String()
		flight.Destination_City = json_data.Search(dest_iata).Search("city").String()
	}

	if _, ok := airport_codes[source_iata]; ok {
		flight.Source_Airport = json_data.Search(source_iata).Search("name").String()
	}
}
