package bart

import (
	"encoding/xml"
	"fmt"
)

type Count struct {
	Uri       string `xml:"uri"`
	Date      string `xml:"date"`
	Time      string `xml:"time"`
	NumTrains int    `xml:"traincount"`
	Message
}

type Elevator struct {
	Uri         string `xml:"uri"`
	Date        string `xml:"date"`
	Time        string `xml:"time"`
	Posted      string `xml:"bsa>posted"`
	Description string `xml:"bsa>description"`
	Message
}

type BSA struct {
	Uri         string `xml:"uri"`
	Type        string `xml:"type"`
	Date        string `xml:"date"`
	Time        string `xml:"time"`
	Posted      string `xml:"bsa>posted"`
	Description string `xml:"bsa>description"`
	Expires     string `xml:"bsa>expires"`
	Message
}

func GetBSA() (*BSA, error) {
	resp, err := GetDataFromUri("http://api.bart.gov/api/bsa.aspx?cmd=bsa&key=MW9S-E7SL-26DU-VV8V")
	if err != nil {
		return nil, err
	}

	bsa := new(BSA)
	err = xml.Unmarshal(resp, bsa)
	if err != nil {
		return nil, err
	}

	if bsa.ErrorText != "" && bsa.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", bsa.ErrorText, bsa.ErrorDetails)
	}

	return bsa, nil
}

func GetElevatorStatus() (*Elevator, error) {
	resp, err := GetDataFromUri("http://api.bart.gov/api/bsa.aspx?cmd=elev&key=MW9S-E7SL-26DU-VV8V")
	if err != nil {
		return nil, err
	}

	elevator := new(Elevator)
	err = xml.Unmarshal(resp, elevator)
	if err != nil {
		return nil, err
	}

	if elevator.ErrorText != "" && elevator.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", elevator.ErrorText, elevator.ErrorDetails)
	}

	return elevator, nil
}

func GetActiveTrainCount() (*Count, error) {
	resp, err := GetDataFromUri("http://api.bart.gov/api/bsa.aspx?cmd=count&key=MW9S-E7SL-26DU-VV8V")
	if err != nil {
		return nil, err
	}

	count := new(Count)

	err = xml.Unmarshal(resp, count)
	if err != nil {
		return nil, err
	}

	if count.ErrorText != "" && count.ErrorDetails != "" {
		return nil, fmt.Errorf("Error: %s, %s", count.ErrorText, count.ErrorDetails)
	}

	return count, nil
}
