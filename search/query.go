package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Make a request (w/ search query) to localGoogoo,
// and parse the json string response into a Result struct
// return that result

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

// API struct for handling http requests
type API struct {
	Client  *http.Client
	BaseURL string
}

// QueryDB queries the localGoogoo database for a search query
func (api *API) QueryDB(query string) (Result, error) {
	var r Result

	// location of the search script wrt localgoogoo root dir
	const path = "/php/search/api.search.php"

	// full url string,
	// {host}/{path}?{query}
	urlString := fmt.Sprintf("%s%s?q=%s", api.BaseURL, path, url.PathEscape(query))

	resp, err := api.Client.Get(urlString)
	if err != nil {
		// log.Fatalln(err)
		return r, err
	}

	if resp.StatusCode != 200 {
		log.Fatalf(`
%d error, Please make sure your localGoogoo is up and running (and at least at v1.0.4)
If it is, then make sure your HOST is correctly set  (see '--config' flag)

Try opening the following url in your browser, it should return a JSON string
(%s)

Create an issue if you think this is a bug: (http://github.com/kodejuice/localgoogoo-go/issues/new).
`, resp.StatusCode, urlString)
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Fatalln(err)
		return r, err
	}
	defer resp.Body.Close()

	// decode API response which is
	// currently in JSON format
	err = json.Unmarshal(body, &r)

	if err != nil {
		if string(body) != "[]" {
			return r, fmt.Errorf("failed to parse json response")
		}
	}

	return r, nil
}
