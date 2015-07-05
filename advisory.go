package bart


import (
  "encoding/xml"
)

type Count struct {
    Uri string `xml:"uri"`
    Date string `xml:"date"`
    Time string `xml:"time"`
    NumTrains int `xml:"traincount"`
}

type Elevator struct {
  Uri string `xml:"uri"`
  Date string `xml:"date"`
  time string `xml:"time"`
  Posted string `xml:"bsa>posted"`
  Description string `xml:"bsa>description"`
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
  return count, nil
}
