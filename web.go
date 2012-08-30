// Steve Phillips / elimisteve
// 2012.06.09

package fun

import (
	// "fmt"
	"io/ioutil"
	"net/http"
)

func UrlToHtml(url string) (string, error) {
	// Download contents of `url`
	req, err := http.Get(url)
	if err != nil {
		// fmt.Printf("Error GET'ing %s: %s\n", url, err)
		return "", err
	}
	// Save contents of page to `html` variable
	html, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// fmt.Printf("Error reading from req.Body: %s\n", err)
		return "", err
	}
	req.Body.Close()
	return string(html), nil
}
