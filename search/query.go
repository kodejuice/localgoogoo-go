package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

// location of the search script wrt localgoogoo root dir
const path = "/php/search/api.search.php"

// Search for a query
func Search(host string, query string) Result {
	urlString := fmt.Sprintf("%s%s?q=%s", host, path, url.PathEscape(query))

	resp, err := http.Get(urlString)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf(`
%d error, Please make sure your localGoogoo is up to date
If it is, then make sure your HOST is correctly set  (see '--config' flag)

Try opening the following url on your browser, it should return a JSON string
(%s)

Create an issue if you think this is a bug: (http://github.com/kodejuice/localgoogoo-go/issues/new).
`, resp.StatusCode, urlString)
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
