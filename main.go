package main

import (
	"bart"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Exploring the BART API")
	stations_info, err := bart.GetStations()
	if err != nil {
		log.Fatalf("Error getting stations info: %s\n", err)
	}

	for _, station := range stations_info.Stations {
    stationAccess, err := bart.GetStationAccessInfo(station.Abbr)
    if err != nil {
  		log.Fatalf("Error getting station access info: %s\n", err)
  	}

    stationInfo, err := bart.GetStationInfo(station.Abbr)
    if err != nil {
  		log.Fatalf("Error getting station access info: %s\n", err)
  	}

    fmt.Println(stationAccess)
    fmt.Println(stationInfo)
  }
}
