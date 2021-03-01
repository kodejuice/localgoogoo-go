package crawler

import (
	"fmt"
	"net/http"
	"net/url"
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

	_, err := api.Client.PostForm(urlString, formData)
	if err != nil {
		return err
	}

	done <- true
	return nil
}

// continously poll the localGoogoo database, reporting
// the number of pages crawled
func monitorProgress(done chan bool, siteName, siteURL string) {
	reportStat := func() {
		count := countPagesCrawled(siteName, siteURL)
		fmt.Printf("%d Crawled Pages\r", count)
	}

	for {
		select {
		case <-done:
			{
				reportStat() // report one last time before exiting
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
func countPagesCrawled(siteName, siteURL string) int {
	// TODO: were here
	return 1
}

// InitCrawler initializes the crawlng process
func (api *API) InitCrawler(siteName, siteURL string) error {
	var done = make(chan bool, 1)

	api.Payload.siteName = siteName
	api.Payload.siteURL = siteURL

	// monitor progress of the crawling process
	go monitorProgress(done, siteName, siteURL)

	err := api.launchCrawler(done)

	fmt.Println("Crawl complete!")
	return err
}

/*
crawl(name, site, done):
	http request to localGoogoo
	done <- 1

monitor_progress(ch):
	select {
		case <-ch: {
			return
		}
		default: {
			count := http equest to get stats
			printf("%d %s\r", count, "Crawled sites")
		}
	}

start():
	done := make(chan, 1)
	go monitor_progress(done)
	crawl(..., done)

*/
