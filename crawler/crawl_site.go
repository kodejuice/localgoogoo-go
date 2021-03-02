package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// API struct for handling http requests
type API struct {
	Client  *http.Client
	BaseURL string

	Payload struct {
		siteName string
		siteURL  string
	}
}

// this inits the request to localGoogoo crawler script
func (api *API) launchCrawler(done chan bool) error {
	// location of the crawler script wrt localgoogoo root dir
	var path = "/php/start_crawler.php"

	var formData = url.Values{
		"web_name": {api.Payload.siteName},
		"web_url":  {api.Payload.siteURL},
	}

	var urlString = fmt.Sprintf("%s%s", api.BaseURL, path)

	fmt.Print("\nCrawling website...\n\n")

	_, err := api.Client.PostForm(urlString, formData)
	if err != nil {
		return err
	}

	done <- true
	return nil
}

// continuously poll the localGoogoo database, reporting
// the number of pages crawled
func (api *API) monitorProgress(done, progressDone chan bool, siteName, siteURL string) {
	reportStat := func() {
		count := countPagesCrawled(api, siteName, siteURL)
		fmt.Printf("%d Crawled Pages\r", count)
	}

	for {
		select {
		case <-done:
			{
				reportStat() // report one last time before exiting
				progressDone <- true
				return
			}
		default:
			{
				reportStat()
			}
		}
	}
}

// given a site name and url, return the number of pages
// of that site stored in the database
func countPagesCrawled(api *API, siteName, siteURL string) int {
	var pagesCount = 0
	// location of the pages count script wrt localgoogoo root dir
	var path = "/php/get_pages_count.php"

	var urlString = fmt.Sprintf("%s%s?sitename=%s&siteurl=%s", api.BaseURL, path, siteName, siteURL)

	// make request
	resp, err := api.Client.Get(urlString)
	if err != nil {
		return pagesCount
	}

	// read in response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pagesCount
	}
	defer resp.Body.Close()

	// convert response to integers
	v, err := strconv.Atoi(string(body))
	if err != nil {
		return pagesCount
	}

	return v
}

// InitCrawler initializes the crawlng process
func (api *API) InitCrawler(siteName, siteURL string) error {
	var done = make(chan bool, 1)
	var progressDone = make(chan bool, 1)

	api.Payload.siteName = siteName
	api.Payload.siteURL = siteURL

	// launch goroutine to monitor
	// progress of the crawling process
	go api.monitorProgress(done, progressDone, siteName, siteURL)

	err := api.launchCrawler(done)

	// make sure we've reported the final pages count
	// before returning to caller
	<-progressDone

	fmt.Println("\n\nCrawl complete!")
	return err
}
