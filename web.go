// Steve Phillips / elimisteve
// 2012.06.09

package fun

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Scrape(url string) ([]byte, error) {
	// Download contents of `url`
	req, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error GET'ing %s: %s\n", url, err)
	}
	// Save contents of page to `html` variable
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading from req.Body: %s\n", err)
	}
	req.Body.Close()
	return data, nil
}
