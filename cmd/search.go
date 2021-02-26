/*
Copyright © 2020 Sochima Biereagu <sochima.agu1@gmail.com>
This file is part of the CLI application localgoogoo.
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/kodejuice/localgoogoo-go/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search the localgoogoo database",
	Long:  `Makes an HTTP request to the localgoogoo installed on your system, rendering the search result on your terminal`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a query argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args[0:], " ")
		var result search.Result = searchResult(query)

		printResult(result)
	},
}

func searchResult(q string) search.Result {
	// get host from config file
	host := viper.Get("HOST").(string)
	// should never return nil, default already set in root.go

	return search.Search(host, q)
}

// output search result to stdout
func printResult(r search.Result) {
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()

	fmt.Printf("%d Result(s) (%.2f seconds)", r.Total, r.QueryTime)

	ln := len(r.Results)
	for i := ln - 1; i >= 0; i-- {
		item := r.Results[i]
		fmt.Printf("\n\n%s %s\n%s\n%s", cyan(i+1), green(item.Title), yellow(item.URL), item.Content)
	}

	// print new line
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
