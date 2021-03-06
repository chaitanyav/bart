package bart

import (
	"encoding/xml"
	"fmt"
)

type Message struct {
	ErrorText    string `xml:"message>error>text"`
	ErrorDetails string `xml:"message>error>details"`
}

type Root struct {
	Uri      string    `xml:"uri"`
	Stations []Station `xml:"stations>station"`
	Message
}

func (root *Root) String() string {
	ret := ""
	for _, station := range root.Stations {
		ret += fmt.Sprintf(" Name: %s,\n Abbr: %s,\n Latitude: %g,\n Longitude: %g,\n Address: %s,\n City: %s,\n County: %s,\n State: %s,\n ZipCode: %d\n", station.Name,
			station.Abbr,
			station.Latitude,
			station.Longitude,
			station.Address,
			station.City,
			station.County,
			station.State,
			station.Zip)
		ret += "\n"
	}

	return ret
}

func GetStations() (*Root, error) {
	resp, err := GetDataFromUri("http://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V")
	if err != nil {
		return nil, err
	}

	data := new(Root)
	err = xml.Unmarshal(resp, data)
	if err != nil {
		return nil, err
	}

	if data.ErrorText != "" && data.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", data.ErrorText, data.ErrorDetails)
	}
	return data, nil
}
