package bart


import (
  "net/http"
  "io/ioutil"
)

func GetDataFromUri(uri string) ([]byte, error) {
  resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

  return body, nil
}
