package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type item struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

// Result for search
type Result struct {
	Results   []item  `json:"results"`   // array of results
	Total     int64   `json:"total"`     // count of result
	QueryTime float64 `json:"queryTime"` // tim taken to run search
}

// TODO: let user set this
const host = "http://localhost/localgoogoo/"

// location of the search script wrt host
const path = "/php/search/api.search.php"

// Search for a query
func Search(query string) Result {
	url := fmt.Sprintf("%s%s?q=%s", host, path, query)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf(`
%d error, Please make sure your localGoogoo is up to date

If it is, then make sure your HOST is correctly set
(%s)
`, resp.StatusCode, host)
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body to type string
	// result := string(body)
	// log.Printf(result)

	// decode API response which is
	// currently in JSON format
	var r Result
	err = json.Unmarshal(body, &r)

	return r
}
