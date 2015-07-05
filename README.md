A golang wrapper for the BART API

### Get BART Stations along with Station info and Access information ###
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
