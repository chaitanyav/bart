package bart

import (
	"encoding/xml"
	"fmt"
)

type StationInfo struct {
	NorthRoutes    []string `xml:"north_routes>route"`
	SouthRoutes    []string `xml:"south_routes>route"`
	NorthPlatforms []string `xml:"north_platforms>platform"`
	SouthPlatforms []string `xml:"south_platforms>platform"`
	PlatformInfo   string   `xml:"platform_info"`
	Intro          string   `xml:"intro"`
	CrossStreet    string   `xml:"cross_street"`
	Food           string   `xml:"food"`
	Shopping       string   `xml:"shopping"`
	Attractions    string   `xml:"attraction"`
	Link           string   `xml:"link"`
}

type StationAccess struct {
	Entering     string `xml:"entering"`
	Exiting      string `xml:"exiting"`
	Parking      string `xml:"parking"`
	FillTime     string `xml:"fill_time"`
	CarShare     string `xml:"car_share"`
	Lockers      string `xml:"lockers"`
	Destinations string `xml:"destinations"`
	TransitInfo  string `xml:"transit_info"`
}

type Station struct {
	Name      string  `xml:"name"`
	Abbr      string  `xml:"abbr"`
	Latitude  float64 `xml:"gtfs_latitude"`
	Longitude float64 `xml:"gtfs_longitude"`
	Address   string  `xml:"address"`
	City      string  `xml:"city"`
	County    string  `xml:"county"`
	State     string  `xml:"state"`
	Zip       int     `xml:"zipcode"`
	StationInfo
	StationAccess
}

func GetStationInfo(stationCode string) (*StationInfo, error) {
	uri := fmt.Sprintf("http://api.bart.gov/api/stn.aspx?cmd=stninfo&orig=%s&key=MW9S-E7SL-26DU-VV8V", stationCode)
	resp, err := GetDataFromUri(uri)
	if err != nil {
		return nil, err
	}
	data := new(Root)
	err = xml.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}

	stationInfo := new(StationInfo)
	station := data.Stations[0]
	stationInfo.NorthRoutes = station.NorthRoutes
	stationInfo.SouthRoutes = station.SouthRoutes
	stationInfo.NorthPlatforms = station.NorthPlatforms
	stationInfo.SouthPlatforms = station.SouthPlatforms
	stationInfo.PlatformInfo = station.PlatformInfo
	stationInfo.Intro = station.Intro
	stationInfo.CrossStreet = station.CrossStreet
	stationInfo.Food = station.Food
	stationInfo.Shopping = station.Shopping
	stationInfo.Attractions = station.Attractions
	stationInfo.Link = station.Link

	return stationInfo, nil
}

func GetStationAccessInfo(stationCode string) (*StationAccess, error) {
	uri := fmt.Sprintf("http://api.bart.gov/api/stn.aspx?cmd=stnaccess&orig=%s&key=MW9S-E7SL-26DU-VV8V&l=1", stationCode)
	resp, err := GetDataFromUri(uri)
	if err != nil {
		return nil, err
	}
	data := new(Root)
	err = xml.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}

	stationAccess := new(StationAccess)
	station := data.Stations[0]
	stationAccess.Entering = station.Entering
	stationAccess.Exiting = station.Exiting
	stationAccess.Parking = station.Parking
	stationAccess.FillTime = station.FillTime
	stationAccess.CarShare = station.CarShare
	stationAccess.Lockers = station.Lockers
	stationAccess.Destinations = station.Destinations
	stationAccess.TransitInfo = station.TransitInfo
	return stationAccess, nil
}
