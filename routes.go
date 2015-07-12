package bart

import (
	"encoding/xml"
	"fmt"
)

type Route struct {
	Name        string   `xml:"name"`
	Abbr        string   `xml:"abbr"`
	RouteID     string   `xml:"routeID"`
	Number      int      `xml:"number"`
	Origin      string   `xml:"origin"`
	Color       string   `xml:"color"`
	Direction   string   `xml:"direction"`
	Holidays    int      `xml:"holidays"`
	NumStations int      `xml:"num_stns"`
	Stns        []string `xml:"config>station"`
}

type RoutesRoot struct {
	Uri    string  `xml:"uri"`
	Routes []Route `xml:"routes>route"`
	Message
}

type RouteInfoRoot struct {
	Uri       string `xml:"uri"`
	RouteInfo Route  `xml:"routes>route"`
	Message
}

func GetRoutes() ([]Route, error) {

	resp, err := GetDataFromUri("http://api.bart.gov/api/route.aspx?cmd=routes&key=MW9S-E7SL-26DU-VV8V")
	if err != nil {
		return make([]Route, 0), err
	}

	routes := new(RoutesRoot)
	err = xml.Unmarshal(resp, routes)
	if err != nil {
		return make([]Route, 0), err
	}

	if routes.ErrorText != "" && routes.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", routes.ErrorText, routes.ErrorDetails)
	}

	return routes.Routes, nil
}

func GetRouteInfo(routeNumber int) (*Route, error) {
	uri := fmt.Sprintf("http://api.bart.gov/api/route.aspx?cmd=routeinfo&route=%d&key=MW9S-E7SL-26DU-VV8V", routeNumber)
	resp, err := GetDataFromUri(uri)
	if err != nil {
		return nil, err
	}

	routeInfo := new(RouteInfoRoot)
	err = xml.Unmarshal(resp, routeInfo)
	if err != nil {
		return nil, err
	}

	if routeInfo.ErrorText != "" && routeInfo.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", routeInfo.ErrorText, routeInfo.ErrorDetails)
	}

	return &routeInfo.RouteInfo, nil
}
