package bart

import (
	"io/ioutil"
	"net/http"
	"strings"
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

	//******this is a hack***********
	// i am tricking the parser to think that the
	// iso-8859-1 is utf-8
	str := string(body)
	str = strings.Replace(str, "iso-8859-1", "utf-8", -1)
	return []byte(str), nil
}
