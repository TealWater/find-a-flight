package utils

import (
	"time"
)

type FlightAwareFlight struct {
	Links struct {
		Next string `json:"next"`
	} `json:"links"`
	NumPages            int `json:"num_pages"`
	ScheduledDepartures []struct {
		Ident             string   `json:"ident"`
		IdentIcao         string   `json:"ident_icao"`
		IdentIata         string   `json:"ident_iata"`
		FaFlightID        string   `json:"fa_flight_id"`
		Operator          string   `json:"operator"`
		OperatorIcao      string   `json:"operator_icao"`
		OperatorIata      string   `json:"operator_iata"`
		FlightNumber      string   `json:"flight_number"`
		Registration      string   `json:"registration"`
		AtcIdent          string   `json:"atc_ident"`
		InboundFaFlightID string   `json:"inbound_fa_flight_id"`
		Codeshares        []string `json:"codeshares"`
		CodesharesIata    []string `json:"codeshares_iata"`
		Blocked           bool     `json:"blocked"`
		Diverted          bool     `json:"diverted"`
		Cancelled         bool     `json:"cancelled"`
		PositionOnly      bool     `json:"position_only"`
		Origin            struct {
			Code           string `json:"code"`
			CodeIcao       string `json:"code_icao"`
			CodeIata       string `json:"code_iata"`
			CodeLid        string `json:"code_lid"`
			Timezone       string `json:"timezone"`
			Name           string `json:"name"`
			City           string `json:"city"`
			AirportInfoURL string `json:"airport_info_url"`
		} `json:"origin"`
		Destination struct {
			Code           string `json:"code"`
			CodeIcao       string `json:"code_icao"`
			CodeIata       string `json:"code_iata"`
			CodeLid        string `json:"code_lid"`
			Timezone       string `json:"timezone"`
			Name           string `json:"name"`
			City           string `json:"city"`
			AirportInfoURL string `json:"airport_info_url"`
		} `json:"destination"`
		DepartureDelay      int       `json:"departure_delay"`
		ArrivalDelay        int       `json:"arrival_delay"`
		FiledEte            int       `json:"filed_ete"`
		ProgressPercent     int       `json:"progress_percent"`
		Status              string    `json:"status"`
		AircraftType        string    `json:"aircraft_type"`
		RouteDistance       int       `json:"route_distance"`
		FiledAirspeed       int       `json:"filed_airspeed"`
		FiledAltitude       int       `json:"filed_altitude"`
		Route               string    `json:"route"`
		BaggageClaim        string    `json:"baggage_claim"`
		SeatsCabinBusiness  int       `json:"seats_cabin_business"`
		SeatsCabinCoach     int       `json:"seats_cabin_coach"`
		SeatsCabinFirst     int       `json:"seats_cabin_first"`
		GateOrigin          string    `json:"gate_origin"`
		GateDestination     string    `json:"gate_destination"`
		TerminalOrigin      string    `json:"terminal_origin"`
		TerminalDestination string    `json:"terminal_destination"`
		Type                string    `json:"type"`
		ScheduledOut        time.Time `json:"scheduled_out"`
		EstimatedOut        time.Time `json:"estimated_out"`
		ActualOut           time.Time `json:"actual_out"`
		ScheduledOff        time.Time `json:"scheduled_off"`
		EstimatedOff        time.Time `json:"estimated_off"`
		ActualOff           time.Time `json:"actual_off"`
		ScheduledOn         time.Time `json:"scheduled_on"`
		EstimatedOn         time.Time `json:"estimated_on"`
		ActualOn            time.Time `json:"actual_on"`
		ScheduledIn         time.Time `json:"scheduled_in"`
		EstimatedIn         time.Time `json:"estimated_in"`
		ActualIn            time.Time `json:"actual_in"`
	} `json:"scheduled_departures"`
}

type AirlineData struct {
	Departure_time      string
	Carrier_iata        string
	Flight_Number       string
	Source_Airport      string
	Destination_Airport string
	Best_Flight         string
	Cheapest_Flight     string
	Fastest_Flight      string
}

type SkyscannerData struct {
	Query struct {
		Market    string `json:"market"`
		Locale    string `json:"locale"`
		Currency  string `json:"currency"`
		QueryLegs []struct {
			OriginPlaceID struct {
				Iata string `json:"iata"`
				//EntityID string `json:"entityId"`
			} `json:"originPlaceId"`
			DestinationPlaceID struct {
				Iata string `json:"iata"`
				//EntityID string `json:"entityId"`
			} `json:"destinationPlaceId"`
			Date struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"date"`
		} `json:"queryLegs"`
		CabinClass string `json:"cabinClass"`
		Adults     int    `json:"adults"`
		//ChildrenAges              []int    `json:"childrenAges"`
		IncludedCarriersIds []string `json:"includedCarriersIds"`
		//ExcludedCarriersIds       []string `json:"excludedCarriersIds"`
		//IncludedAgentsIds         []string `json:"includedAgentsIds"`
		//ExcludedAgentsIds         []string `json:"excludedAgentsIds"`
		//IncludeSustainabilityData bool     `json:"includeSustainabilityData"`
		//NearbyAirports            bool     `json:"nearbyAirports"`
	} `json:"query"`
}

type SkyscannerResults struct {
	SessionToken string `json:"sessionToken"`
	Status       string `json:"status"`
	Action       string `json:"action"`
	Content      struct {
		Results struct {
			Itineraries struct {
				Property1 struct {
					PricingOptions []struct {
						Price struct {
							Amount       string `json:"amount"`
							Unit         string `json:"unit"`
							UpdateStatus string `json:"updateStatus"`
						} `json:"price"`
						AgentIds []string `json:"agentIds"`
						Items    []struct {
							Price struct {
							} `json:"price"`
							AgentID  string `json:"agentId"`
							DeepLink string `json:"deepLink"`
							Fares    []any  `json:"fares"`
						} `json:"items"`
						TransferType string `json:"transferType"`
						ID           string `json:"id"`
					} `json:"pricingOptions"`
					LegIds             []string `json:"legIds"`
					SustainabilityData struct {
						IsEcoContender    bool `json:"isEcoContender"`
						EcoContenderDelta int  `json:"ecoContenderDelta"`
					} `json:"sustainabilityData"`
				} `json:"property1"`
				Property2 struct {
					PricingOptions []struct {
						Price struct {
							Amount       string `json:"amount"`
							Unit         string `json:"unit"`
							UpdateStatus string `json:"updateStatus"`
						} `json:"price"`
						AgentIds []string `json:"agentIds"`
						Items    []struct {
							Price struct {
							} `json:"price"`
							AgentID  string `json:"agentId"`
							DeepLink string `json:"deepLink"`
							Fares    []any  `json:"fares"`
						} `json:"items"`
						TransferType string `json:"transferType"`
						ID           string `json:"id"`
					} `json:"pricingOptions"`
					LegIds             []string `json:"legIds"`
					SustainabilityData struct {
						IsEcoContender    bool `json:"isEcoContender"`
						EcoContenderDelta int  `json:"ecoContenderDelta"`
					} `json:"sustainabilityData"`
				} `json:"property2"`
			} `json:"itineraries"`
			Legs struct {
				Property1 struct {
					OriginPlaceID      string `json:"originPlaceId"`
					DestinationPlaceID string `json:"destinationPlaceId"`
					DepartureDateTime  struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"departureDateTime"`
					ArrivalDateTime struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"arrivalDateTime"`
					DurationInMinutes   int      `json:"durationInMinutes"`
					StopCount           int      `json:"stopCount"`
					MarketingCarrierIds []string `json:"marketingCarrierIds"`
					OperatingCarrierIds []string `json:"operatingCarrierIds"`
					SegmentIds          []string `json:"segmentIds"`
				} `json:"property1"`
				Property2 struct {
					OriginPlaceID      string `json:"originPlaceId"`
					DestinationPlaceID string `json:"destinationPlaceId"`
					DepartureDateTime  struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"departureDateTime"`
					ArrivalDateTime struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"arrivalDateTime"`
					DurationInMinutes   int      `json:"durationInMinutes"`
					StopCount           int      `json:"stopCount"`
					MarketingCarrierIds []string `json:"marketingCarrierIds"`
					OperatingCarrierIds []string `json:"operatingCarrierIds"`
					SegmentIds          []string `json:"segmentIds"`
				} `json:"property2"`
			} `json:"legs"`
			Segments struct {
				Property1 struct {
					OriginPlaceID      string `json:"originPlaceId"`
					DestinationPlaceID string `json:"destinationPlaceId"`
					DepartureDateTime  struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"departureDateTime"`
					ArrivalDateTime struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"arrivalDateTime"`
					DurationInMinutes     int    `json:"durationInMinutes"`
					MarketingFlightNumber string `json:"marketingFlightNumber"`
					MarketingCarrierID    string `json:"marketingCarrierId"`
					OperatingCarrierID    string `json:"operatingCarrierId"`
				} `json:"property1"`
				Property2 struct {
					OriginPlaceID      string `json:"originPlaceId"`
					DestinationPlaceID string `json:"destinationPlaceId"`
					DepartureDateTime  struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"departureDateTime"`
					ArrivalDateTime struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"arrivalDateTime"`
					DurationInMinutes     int    `json:"durationInMinutes"`
					MarketingFlightNumber string `json:"marketingFlightNumber"`
					MarketingCarrierID    string `json:"marketingCarrierId"`
					OperatingCarrierID    string `json:"operatingCarrierId"`
				} `json:"property2"`
			} `json:"segments"`
			Places struct {
				Property1 struct {
					EntityID    string `json:"entityId"`
					ParentID    string `json:"parentId"`
					Name        string `json:"name"`
					Type        string `json:"type"`
					Iata        string `json:"iata"`
					Coordinates struct {
						Latitude  int `json:"latitude"`
						Longitude int `json:"longitude"`
					} `json:"coordinates"`
				} `json:"property1"`
				Property2 struct {
					EntityID    string `json:"entityId"`
					ParentID    string `json:"parentId"`
					Name        string `json:"name"`
					Type        string `json:"type"`
					Iata        string `json:"iata"`
					Coordinates struct {
						Latitude  int `json:"latitude"`
						Longitude int `json:"longitude"`
					} `json:"coordinates"`
				} `json:"property2"`
			} `json:"places"`
			Carriers struct {
				Property1 struct {
					Name       string `json:"name"`
					AllianceID string `json:"allianceId"`
					ImageURL   string `json:"imageUrl"`
					Iata       string `json:"iata"`
				} `json:"property1"`
				Property2 struct {
					Name       string `json:"name"`
					AllianceID string `json:"allianceId"`
					ImageURL   string `json:"imageUrl"`
					Iata       string `json:"iata"`
				} `json:"property2"`
			} `json:"carriers"`
			Agents struct {
				Property1 struct {
					Name            string `json:"name"`
					Type            string `json:"type"`
					ImageURL        string `json:"imageUrl"`
					FeedbackCount   int    `json:"feedbackCount"`
					Rating          int    `json:"rating"`
					RatingBreakdown struct {
						CustomerService int `json:"customerService"`
						ReliablePrices  int `json:"reliablePrices"`
						ClearExtraFees  int `json:"clearExtraFees"`
						EaseOfBooking   int `json:"easeOfBooking"`
						Other           int `json:"other"`
					} `json:"ratingBreakdown"`
					IsOptimisedForMobile bool `json:"isOptimisedForMobile"`
				} `json:"property1"`
				Property2 struct {
					Name            string `json:"name"`
					Type            string `json:"type"`
					ImageURL        string `json:"imageUrl"`
					FeedbackCount   int    `json:"feedbackCount"`
					Rating          int    `json:"rating"`
					RatingBreakdown struct {
						CustomerService int `json:"customerService"`
						ReliablePrices  int `json:"reliablePrices"`
						ClearExtraFees  int `json:"clearExtraFees"`
						EaseOfBooking   int `json:"easeOfBooking"`
						Other           int `json:"other"`
					} `json:"ratingBreakdown"`
					IsOptimisedForMobile bool `json:"isOptimisedForMobile"`
				} `json:"property2"`
			} `json:"agents"`
			Alliances struct {
				Property1 struct {
					Name string `json:"name"`
				} `json:"property1"`
				Property2 struct {
					Name string `json:"name"`
				} `json:"property2"`
			} `json:"alliances"`
		} `json:"results"`
		Stats struct {
			Itineraries struct {
				MinDuration int `json:"minDuration"`
				MaxDuration int `json:"maxDuration"`
				Total       struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"total"`
				Stops struct {
					Direct struct {
						Total struct {
							Count    int `json:"count"`
							MinPrice struct {
								Amount       string `json:"amount"`
								Unit         string `json:"unit"`
								UpdateStatus string `json:"updateStatus"`
							} `json:"minPrice"`
						} `json:"total"`
						TicketTypes struct {
							SingleTicket struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"singleTicket"`
							MultiTicketNonNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNonNpt"`
							MultiTicketNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNpt"`
						} `json:"ticketTypes"`
					} `json:"direct"`
					OneStop struct {
						Total struct {
							Count    int `json:"count"`
							MinPrice struct {
								Amount       string `json:"amount"`
								Unit         string `json:"unit"`
								UpdateStatus string `json:"updateStatus"`
							} `json:"minPrice"`
						} `json:"total"`
						TicketTypes struct {
							SingleTicket struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"singleTicket"`
							MultiTicketNonNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNonNpt"`
							MultiTicketNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNpt"`
						} `json:"ticketTypes"`
					} `json:"oneStop"`
					TwoPlusStops struct {
						Total struct {
							Count    int `json:"count"`
							MinPrice struct {
								Amount       string `json:"amount"`
								Unit         string `json:"unit"`
								UpdateStatus string `json:"updateStatus"`
							} `json:"minPrice"`
						} `json:"total"`
						TicketTypes struct {
							SingleTicket struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"singleTicket"`
							MultiTicketNonNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNonNpt"`
							MultiTicketNpt struct {
								Count    int `json:"count"`
								MinPrice struct {
									Amount       string `json:"amount"`
									Unit         string `json:"unit"`
									UpdateStatus string `json:"updateStatus"`
								} `json:"minPrice"`
							} `json:"multiTicketNpt"`
						} `json:"ticketTypes"`
					} `json:"twoPlusStops"`
				} `json:"stops"`
				HasChangeAirportTransfer bool `json:"hasChangeAirportTransfer"`
			} `json:"itineraries"`
		} `json:"stats"`
		SortingOptions struct {
			Best []struct {
				Score       float32 `json:"score"`
				ItineraryID string  `json:"itineraryId"`
			} `json:"best"`
			Cheapest []struct {
				Score       float32 `json:"score"`
				ItineraryID string  `json:"itineraryId"`
			} `json:"cheapest"`
			Fastest []struct {
				Score       float32 `json:"score"`
				ItineraryID string  `json:"itineraryId"`
			} `json:"fastest"`
		} `json:"sortingOptions"`
	} `json:"content"`
}
