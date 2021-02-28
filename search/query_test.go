package search

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kodejuice/localgoogoo-go/util"
)

func TestQueryDB(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		path := "/php/search/api.search.php"
		query := req.URL.Query().Get("q")
		// Test request parameters
		util.Equals(t, req.URL.Path, path)
		util.Equals(t, "hello world", query)

		// Send response to be tested
		testResp := `{"queryTime":0,"total":1,"results":[{"url":"test_url","title":"test_title","content":"test_content"}]}`
		_, _ = rw.Write([]byte(testResp))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL}
	resp, err := api.QueryDB("hello world")
	if err != nil {
		log.Fatalf(err.Error())
	}

	util.Ok(t, err)

	// test search response
	util.Equals(t, int(1), len(resp.Results))
	util.Equals(t, int64(1), resp.Total)
	util.Equals(t, float64(0), resp.QueryTime)
	util.Equals(t, "test_url", resp.Results[0].URL)
	util.Equals(t, "test_title", resp.Results[0].Title)
	util.Equals(t, "test_content", resp.Results[0].Content)
}
