package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dataFile = "data.json"

type Ranking struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type RankingData struct {
	Total        int       `json:"total"`
	Page         *int      `json:"page"`
	Data         []Ranking `json:"data"`
	ItemsPerPage *int      `json:"items_per_page"`
}

func parseName(name string) string {
	nameSplitByBracket := strings.Split(name, "(")
	trimmedName := strings.TrimSpace(nameSplitByBracket[0])
	return trimmedName
}

func main() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		err := fmt.Errorf("failed to open data file: %w", err)
		panic(err)
	}

	var rankingData RankingData
	err = json.Unmarshal(file, &rankingData)
	if err != nil {
		err := fmt.Errorf("failed to unmarshal JSON: %w", err)
		panic(err)
	}

	for _, data := range rankingData.Data {
		fmt.Printf("Name: %s, Points: %d\n", parseName(data.Name), data.Points)
	}
}
