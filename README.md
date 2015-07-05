A golang wrapper for the BART API (work in progress)

### To install the package ###
    go get github.com/chaitanyav/bart

### For Docs ###
    godoc -http=:8000

    http://localhost:8000/pkg/github.com/chaitanyav/bart/

### Sample code to get BART Stations along with Station info and Station Access information ###
    package main

    import (
    "github.com/chaitanyav/bart"
    "fmt"
    "log"
    )

    func main() {
      stationsInfo, err := bart.GetStations()
        if err != nil {
          log.Fatalf("Error getting stations info: %s\n", err)
        }

      for _, station := range stationsInfo.Stations {
        fmt.Printf("Name: %q, Abbr: %q, Latitude: %g, Longitude: %g, Address: %q, City: %q, County: %q, State: %q, Zip: %d\n\n", station.Name, station.Abbr, station.Latitude, station.Longitude, station.Address, station.City, station.County, station.State, station.Zip)
      }

      stationAccessInfo, err := bart.GetStationAccessInfo("19th")
        if err != nil {
          log.Fatalf("Error getting station access info: %s\n", err)
        }

      fmt.Printf("Entering: %s\nExiting: %s\nParking: %s\nFillTime: %s\nCarShare: %s\nLockers: %s\nDestinations: %s\nTransitInfo: %s\n", stationAccessInfo.Entering, stationAccessInfo.Exiting, stationAccessInfo.Parking, stationAccessInfo.FillTime, stationAccessInfo.CarShare, stationAccessInfo.Lockers, stationAccessInfo.Destinations, stationAccessInfo.TransitInfo)
      stationInfo, err := bart.GetStationInfo("24th")
        if err != nil {
          log.Fatalf("Error getting station access info: %s\n",
              err)
        }
      fmt.Printf("NorthRoutes: %s\nSouthRoutes: %s\nNorthPlatforms: %s\nSouthPlatforms: %s\nPlatformInfo: %s\nIntro: %s\nCrossStreet: %s\nFood: %s\nShopping: %s\nAttractions: %s\nLink: %s\n", stationInfo.NorthRoutes, stationInfo.SouthRoutes, stationInfo.NorthPlatforms, stationInfo.SouthPlatforms, stationInfo.PlatformInfo, stationInfo.Intro, stationInfo.CrossStreet, stationInfo.Food, stationInfo.Shopping, stationInfo.Attractions, stationInfo.Link)
    }
