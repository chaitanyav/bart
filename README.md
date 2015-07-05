A golang wrapper for the BART API

### Sample code to get BART Stations along with Station info and Access information ###
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

### Truncated output of the above###
Name: "San Leandro", Abbr: "SANL", Latitude: 37.72261921, Longitude: -122.1613112, Address: "1401 San Leandro Blvd.", City: "San Leandro", County: "alameda", State: "CA", Zip: 94577

Name: "South Hayward", Abbr: "SHAY", Latitude: 37.63479954, Longitude: -122.0575506, Address: "28601 Dixon Street", City: "Hayward", County: "alameda", State: "CA", Zip: 94544

Name: "South San Francisco", Abbr: "SSAN", Latitude: 37.664174, Longitude: -122.444116, Address: "1333 Mission Road", City: "South San Francisco", County: "sanmateo", State: "CA", Zip: 94080

Name: "Union City", Abbr: "UCTY", Latitude: 37.591208, Longitude: -122.017867, Address: "10 Union Square", City: "Union City", County: "alameda", State: "CA", Zip: 94587

Name: "Walnut Creek", Abbr: "WCRK", Latitude: 37.905628, Longitude: -122.067423, Address: "200 Ygnacio Valley Road", City: "Walnut Creek", County: "contracosta", State: "CA", Zip: 94596

Name: "West Dublin/Pleasanton", Abbr: "WDUB", Latitude: 37.699759, Longitude: -121.928099, Address: "6501 Golden Gate Drive", City: "Dublin", County: "alameda", State: "CA", Zip: 94568

Name: "West Oakland", Abbr: "WOAK", Latitude: 37.80467476, Longitude: -122.2945822, Address: "1451 7th Street", City: "Oakland", County: "alameda", State: "CA", Zip: 94607

Entering: <p>The BART trains at 19th Street/Oakland Station are located two and three levels below the street. The first level below ground is the concourse level. The second level is the upper platform level and third level is the lower platform level.</p>
<p>Two separate elevators (referred to as Street and Platform) are needed to get from the street to the train platforms.</p>
<p>The street access to the elevator is located at 1746 &ndash; 1750 Broadway between 17th Street and 19th Street.</p>
<p>From the street, the Street elevator will take you to the concourse level where the station agent and paid area are located. The elevator buttons are labeled "C" for Concourse and "S" for Street.</p>
<p>When you arrive at the concourse level you will need to travel approximately 90 feet to reach the Platform elevator. The Platform elevator buttons are labeled "C" for Concourse, "UL" upper platform, and "LL" for lower platform."</p>
<p>When going between the Street elevator and Platform elevator there is no danger of encountering train tracks.</p>
<p>Depending on your destination, the Platform elevator will take you to the upper platform or the lower platform. When the elevator arrives at eith er platform you will be at the far end of the platforms. Go straight to reach the center of the train boarding area.</p>
Exiting: <p>The BART trains at 19th Street/Oakland Station are located two and three levels below the street.</p>
<p>Two separate elevators (referred to as Platform and Street) are needed to get from the train platforms to the street.</p>
<p>Prior to 8:45 AM weekdays Fremont and San Francisco &ndash; Daly City- bound trains arrive at both the upper platform and lower platform. After 8:45 AM on weekdays Fremont and San Francisco &ndash; Daly City bound trains arrive at the lower platform only. When departing a Fremont or San Francisco - Daly City train, the elevator will be towards the front of the train at the far end of the platform. The elevator buttons are labeled "C" for concourse, "UL" for upper platform, and "LL" for lower platform.</p>
<p>When departing a Richmond or Pittsburg &ndash; Bay Point train you will be on the upper platform two levels below the street. When exiting the trains the elevator will be towards the rear of the train at the far end of the platform. The elevator buttons are labeled "C" for concourse, "UL" for upper platform, and "LL" for lower platform.</p>
<p>Once you arrive at the concourse level you will be exiting the Platform elevator directly across from an information booth. The Street elevator is approximately 90 feet to your right. The Street elevator buttons are labeled "C" for Concourse and "S" for Street.</p>
<p>When going between the Platform elevator and Street elevator there is no danger of encountering train tracks.</p>
<p>The Street elevator will take you to the street. When you exit the Street elevator you will be at 1746 &ndash; 1750 Broadway between 17th Street and 19th Street.</p>
Parking: <p>There's no parking at 19th Street Station. The closest station with parking is <a href="http://www.bart.gov/stations/MCAR/">MacArthur Station</a>.</p>
FillTime:
CarShare:
Lockers: <p>8 electronic bicycle lockers managed by City of Oakland.&nbsp;For more information on these lockers, contact the City of Oakland at&nbsp;<a href="http://www.bart.gov/stations/19TH/mailto:bikeped@oaklandnet.com">bikeped@oaklandnet.com</a> or (510) 238-3983.</p>
Destinations: <p>To find a destination near this station, visit the <a href="http://www.bart.gov/stations/19th/neighborhood">neighborhood section</a>.</p>
TransitInfo: <p>19th St. Oakland Station is served by <a rel="external" href="http://www.actransit.org">AC Transit</a>. The <a target="_blank" rel="external" href="http://www.greyhound.com/">Greyhound</a> bus terminal is located three blocks down 20th Street at San Pablo Avenue.</p>
<p><strong>Try BART's free trip planning service!</strong><br />For a personalized trip plan with BART and connecting transit, call the BART <a href="http://www.bart.gov/siteinfo/contact">Transit Information Center</a>. It's fast, it's easy, and it's tailored just for you!</p>
NorthRoutes: [ROUTE 2  ROUTE 6  ROUTE 8  ROUTE 12]  
SouthRoutes: [ROUTE 1  ROUTE 5  ROUTE 7  ROUTE 11]  
NorthPlatforms: [2]  
SouthPlatforms: [1]  
PlatformInfo: Always check destination signs and listen for departure announcements.  
Intro: "The Mission" refers to the San Francisco de Asis Mission, also known as Mission Dolores, which was founded 1776. Today the neighborhood is host to an eclectic mix of restaurants, markets, performance spaces, shops, and nightspots.  
CrossStreet: Nearby Cross: 24th St.  
Food: Nearby restaurant reviews from <a rel="external" href="http://www.yelp.com/search?find_desc=Restaurant+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
Shopping: Local shopping from <a rel="external" href="http://www.yelp.com/search?find_desc=Shopping+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
Attractions: More station area attractions from <a rel="external" href="http://www.yelp.com/search?find_desc=+&amp;ns=1&amp;rpp=10&amp;find_loc=2800 Mission Street San Francisco, CA 94110">yelp.com</a>  
Link: http://www.bart.gov/stations/24TH
