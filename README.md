A golang wrapper for the BART API

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

### Fields of a *Station* ###
***Name***: "San Leandro", ***Abbr***: "SANL", ***Latitude***: 37.72261921, ***Longitude***: -122.1613112, ***Address***: "1401 San Leandro Blvd.", ***City***: "San Leandro", ***County***: "alameda", ***State***: "CA", ***Zip***: 94577

### Fields of *StationAccess*###

**Entering**: <p>The BART trains at 19th Street/Oakland Station are located two and three levels below the street. The first level below ground is the concourse level. The second level is the upper platform level and third level is the lower platform level.</p>
........  
***Exiting***: <p>The BART trains at 19th Street/Oakland Station are located two and three levels below the street.</p>
<p>Two separate elevators (referred to as Platform and Street) are needed to get from the train platforms to the street.</p>....  
***Parking***: <p>There's no parking at 19th Street Station. The closest station with parking is <a href="http://www.bart.gov/stations/MCAR/">MacArthur Station</a>.</p>
***FillTime***:  
***CarShare***:  
***Lockers***: <p>8 electronic bicycle lockers managed by City of Oakland.&nbsp;For more information on these lockers, contact the City of Oakland at&nbsp;<a href="http://www.bart.gov/stations/19TH/mailto:bikeped@oaklandnet.com">bikeped@oaklandnet.com</a> or (510) 238-3983.</p>
***Destinations***: <p>To find a destination near this station, visit the <a href="http://www.bart.gov/stations/19th/neighborhood">neighborhood section</a>.</p>
***TransitInfo***: <p>19th St. Oakland Station is served by <a rel="external" href="http://www.actransit.org">AC Transit</a>. The <a target="_blank" rel="external" href="http://www.greyhound.com/">Greyhound</a> bus terminal is located three blocks down 20th Street at San Pablo Avenue.</p>.....  

### Fields of *StationInfo* ###
***NorthRoutes***: [ROUTE 2  ROUTE 6  ROUTE 8  ROUTE 12]  
***SouthRoutes***: [ROUTE 1  ROUTE 5  ROUTE 7  ROUTE 11]  
***NorthPlatforms***: [2]  
***SouthPlatforms***: [1]  
***PlatformInfo***: Always check destination signs and listen for departure announcements.  
***Intro***: "The Mission" refers to the San Francisco de Asis Mission, also known as Mission Dolores, which was founded 1776. Today the neighborhood is host to an eclectic mix of restaurants, markets, performance spaces, shops, and nightspots.  
***CrossStreet***: Nearby Cross: 24th St.  
***Food***: Nearby restaurant reviews from <a rel="external" href="http://www.yelp.com/search?find_desc=Restaurant+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
***Shopping***: Local shopping from <a rel="external" href="http://www.yelp.com/search?find_desc=Shopping+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
***Attractions***: More station area attractions from <a rel="external" href="http://www.yelp.com/search?find_desc=+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
***Link***: http://www.bart.gov/stations/24TH
