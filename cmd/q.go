package cmd

/*
Copyright © 2020 Sochima Biereagu <sochima.agu1@gmail.com>
This file is part of the CLI application localgoogoo.
*/

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/kodejuice/localgoogoo-go/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flag variables
var resultCount *int64
var reverseResults bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "q <query>",
	Short: "Search the localgoogoo database",
	Long:  `Makes an HTTP request to the localgoogoo installed on your system, rendering the search result on your terminal`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a query argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args[:], " ")
		var result search.Result = searchResult(query)

		printResult(result)
	},
}

func searchResult(q string) search.Result {
	// get host from config file
	// should never return nil, default already set in root.go
	host := viper.Get("HOST").(string)

	api := search.API{Client: http.DefaultClient, BaseURL: host}

	resp, err := api.QueryDB(q)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return resp
}

// output search result to stdout
func printResult(r search.Result) {
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()

	fmt.Printf("%d Result(s) (%.2f seconds)", r.Total, r.QueryTime)

	// how many items are we printing?
	ln := int64(len(r.Results))
	if v := *resultCount; v < ln {
		ln = v
	}

	// loop over results and print

	var currIndex int64 = 0
	if reverseResults {
		currIndex = ln - 1
	}

	for {
		if currIndex == -1 || currIndex == ln {
			break
		}

		item := r.Results[currIndex]
		fmt.Printf("\n\n%s %s\n%s\n%s", cyan(currIndex+1), green(item.Title), yellow(item.URL), item.Content)

		currIndex += advance(reverseResults)
	}

	// print new line
	fmt.Println("")
}

// Used to advance currIndex to next item while printing results.
// if we're printing in reversed order go back `-1`
// else go forward `1`
func advance(reversed bool) int64 {
	if reversed {
		return -1
	}
	return 1
}

func init() {
	resultCount = searchCmd.PersistentFlags().Int64P("count", "n", 10, "number of results to display")
	searchCmd.PersistentFlags().BoolVarP(&reverseResults, "reverse", "r", false, "display results in reversed order")

	rootCmd.AddCommand(searchCmd)
}
