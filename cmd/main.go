package main

import (
	"flag"
	"log"
	"os"

	"haydenheroux.github.io/adapter"
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
)

const (
	APP_NAME          = "tbascraper"
	DEFAULT_YEAR      = 0
	DEFAULT_EVENT     = ""
	DEFAULT_API_KEY   = ""
	DEFAULT_SCOUT_URL = ""
)

var (
	year     int
	eventKey string
	apiKey   string
	scoutURL string
)

func init() {
	flag.IntVar(&year, "year", DEFAULT_YEAR, "Year")
	flag.StringVar(&eventKey, "eventKey", DEFAULT_EVENT, "Event Key")
	flag.StringVar(&apiKey, "apiKey", DEFAULT_API_KEY, "API Key")
	flag.StringVar(&scoutURL, "scoutURL", DEFAULT_SCOUT_URL, "Scout URL")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, APP_NAME+": ", 0)

	api := tba.New(apiKey)

	scout := scout.New(scoutURL)

	event, err := api.GetEvent(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get event: %v\n", err)
	}

	if err := scout.InsertEvent(adapter.ToScoutEvent(event)); err != nil {
		logger.Fatalf("Failed to insert event: %v\n", err)
	}

	teams, err := api.GetTeams(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get teams: %v\n", err)
	}

	for _, team := range teams {
		if err := scout.InsertTeam(adapter.ToScoutTeam(team)); err != nil {
			logger.Fatalf("Failed to insert team: %v\n", err)
		}
	}

	// matches, err := api.GetMatches(eventKey)

	// if err != nil {
	// 	logger.Fatalf("Failed to get matches: %v\n", err)
	// }

	// for _, match := range matches {
	// 	if err := scout.InsertMatch(match); err != nil {
	// 		logger.Fatalf("Failed to insert match: %v\n", err)
	// 	}

	// 	matchKey := match.eventKey

	// 	matchData := api.GetMatchData(matchKey)

	// 	if err := scout.InsertMatchData(matchData); err != nil {
	// 		logger.Fatalf("Failed to insert match data: %v\n", err)
	// 	}
	// }
}
