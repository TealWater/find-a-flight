package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Jeffail/gabs"
)

// look 1 hour ahead in UTC
func Convert_to_iso(off_set int) (start string, end string) {
	var end_time time.Time
	var cur_time time.Time = time.Now().UTC().Add(1 * time.Hour)

	hour_offset := time.Duration(off_set)
	end_time = cur_time.Add(time.Hour * hour_offset)

	return cur_time.Format(time.RFC3339), end_time.Format(time.RFC3339)
}

func Build_skyscanner_data(data []AirlineData) (query []SkyscannerData) {
	var err error
	query = make([]SkyscannerData, len(data))
	for k := range data {
		query[k].Query.Market = "US"
		query[k].Query.Locale = "en-US"
		query[k].Query.Currency = "USD"

		slice := query[k].Query.QueryLegs
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

		slice[0].OriginPlaceID.Iata = data[k].Source_Airport
		slice[0].DestinationPlaceID.Iata = data[k].Destination_Airport
		year_and_month := strings.Split(data[k].Departure_time, "-")
		day := strings.Split(year_and_month[2], " ")

		//properly calibrate the year month day for EST (-4 UTC)
		if new_year, new_day := calibrate_date_to_local_time(day, year_and_month, day, 0); new_year != nil && new_day != nil {
			year_and_month = new_year
			day = new_day
		}

		slice[0].Date.Year, err = strconv.Atoi(year_and_month[0])
		if err != nil {
			log.Fatal(err)
		}

		slice[0].Date.Month, err = strconv.Atoi(year_and_month[1])
		if err != nil {
			log.Fatal(err)
		}

		slice[0].Date.Day, err = strconv.Atoi(day[0])
		if err != nil {
			log.Fatal(err)
		}

		query[k].Query.QueryLegs = append(query[k].Query.QueryLegs, slice...)
		query[k].Query.CabinClass = "CABIN_CLASS_ECONOMY"
		query[k].Query.Adults = 1
		query[k].Query.IncludedCarriersIds = append(query[k].Query.IncludedCarriersIds, data[k].Carrier_iata)
	}
	return query
}

func calibrate_date_to_local_time(utc_time []string, year_and_month_old []string, day_old []string, hour_offset int) (year_and_month []string, day []string) {
	//properly calibrate the year month day for EST
	utc_Time_parsed := strings.Split(utc_time[1], ":")
	hour, _ := strconv.Atoi(utc_Time_parsed[0])
	if (hour - hour_offset) < 0 {
		layout := "2006-01-02"
		cur_day, err := time.Parse(layout, year_and_month_old[0]+"-"+year_and_month_old[1]+"-"+day_old[0])
		if err != nil {
			log.Println("look in /utils")
			log.Fatal(err)
		}
		//go back one day
		cur_day = cur_day.Add(time.Duration(time.Now().UTC().Day()) * -1)
		year_and_month = strings.Split(cur_day.String(), "-")
		day = strings.Split(year_and_month[2], " ")
		return year_and_month, day
	}
	return nil, nil
}

func GetBookinglink_And_Price(json *gabs.Container, sorting_option string, price_data SkyscannerResults) (string, string) {
	defer SafeExit("the itinerary does not exist")

	var itineraryId string
	switch sorting_option {
	case "best":
		itineraryId = price_data.Content.SortingOptions.Best[0].ItineraryID
	case "cheapest":
		itineraryId = price_data.Content.SortingOptions.Cheapest[0].ItineraryID
	case "fastest":
		itineraryId = price_data.Content.SortingOptions.Fastest[0].ItineraryID
	}

	mp, err := json.Search("content", "results", "itineraries", itineraryId).ChildrenMap()
	if err != nil {
		log.Println("I want a spaceship")
		panic(err)
	}

	pricing_options, err := mp["pricingOptions"].ArrayElement(0)
	if err != nil {
		log.Println("you broke me inside another loop! why!!!!")
		panic(err)
	}

	items := pricing_options.Search("items").Search("deepLink")
	price := pricing_options.Search("price").Search("amount")
	price_string := formatCurrency(price.String())
	link := items.String()
	formatLink(&link)
	fmt.Println("the price is: " + price_string)
	return link, price_string
}

func SafeExit(s string) {
	if r := recover(); r != nil {
		log.Println(s)
	}
}

func formatCurrency(data string) string {
	fmt.Println("//////////////////")
	fmt.Println(data)
	fmt.Println("//////////////////")
	result := ""
	for _, r := range data {
		if unicode.IsDigit(r) {
			result += string(r)
		}
	}
	data = result
	slice := []string{}
	data_len := len(data) - 1
	fmt.Println("*************")
	fmt.Println(data)
	fmt.Println("*************")

	for i := 0; i < data_len; i += 1 {
		if i == len(data)-3 {
			slice = append(slice, ".")
		}

		slice = append(slice, string(data[i]))
	}

	data = strings.Join(slice, "")
	fmt.Print("after: ")
	fmt.Println(data)
	return addCommas(data)
}

func addCommas(data string) string {
	if len(data) <= 6 {
		return data
	}
	parsed := strings.Split(data, ".")
	data = parsed[0]

	end := len(data) - 1

	slice := []string{}
	for count := 1; end > -1; end, count = end-1, count+1 {

		if count == 4 {
			slice = append(slice, ",")
			count = 1

		}

		slice = append(slice, string(data[end]))

	}

	swap(&slice)
	slice = append(slice, ".", parsed[1])
	fmt.Println(slice)
	return strings.Join(slice, "")

}

func swap(data *[]string) {
	start := 0
	end := len(*data) - 1
	for ; start < end; start, end = start+1, end-1 {
		(*data)[start], (*data)[end] = (*data)[end], (*data)[start]
	}
}

func formatLink(input *string) {
	runes := []rune(*input)
	end := len(runes)
	runes = append(runes[2:], runes[:end-2]...)
	*input = string(runes)
}
