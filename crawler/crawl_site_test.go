package crawler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kodejuice/localgoogoo-go/util"
)

func TestInitCrawler(t *testing.T) {
	testSiteName := "site_name"
	testSiteURL := "site_url"

	pageCount := 0

	// Start a local HTTP server for the crawler
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		crawlPath := "/php/start_crawler.php"
		monitorProgressPath := "/php/get_pages_count.php"

		if req.URL.Path == crawlPath {
			// Test request parameters
			util.Equals(t, testSiteName, req.PostFormValue("web_name"))
			util.Equals(t, testSiteURL, req.PostFormValue("web_url"))
		} else if req.URL.Path == monitorProgressPath {
			util.Equals(t, testSiteName, req.URL.Query().Get("sitename"))
			util.Equals(t, testSiteURL, req.URL.Query().Get("siteurl"))

			pageCount++
		} else {
			log.Fatalf("Where does this request come from?\nPath:%s\n", req.URL.Path)
		}

		testResp := `1` // dummy response
		_, _ = rw.Write([]byte(testResp))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{Client: server.Client(), BaseURL: server.URL}
	// Set site name&url to crawl
	api.Payload.siteName = testSiteName
	api.Payload.siteURL = testSiteURL

	err := api.InitCrawler(testSiteName, testSiteURL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// the monitor progress path should be hit
	// at least once in this test
	util.Assert(t, pageCount > 0, "'monitor progress' endpoint wasnt hit")
}
