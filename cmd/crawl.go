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
	"net/url"

	"github.com/kodejuice/localgoogoo-go/crawler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl <site_name> <site_url>",
	Short: "Crawl specified wesite",
	Long:  `Use this command to crawl a website, which indexes the pages of the website and stores them in the localGoogoo database`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Requires at least two arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			siteName = args[0]
			siteURL  = args[1]
		)

		// validate URL
		_, err := url.Parse(siteURL)
		if err != nil {
			fmt.Printf("Error: %s (%s)", err.Error(), siteURL)
		}

		crawl(siteName, siteURL)
	},
}

// start crawling
func crawl(siteName, siteURL string) {
	// get host from config file
	host := viper.Get("HOST").(string)

	api := crawler.API{Client: http.DefaultClient, BaseURL: host}

	err := api.InitCrawler(siteName, siteURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
