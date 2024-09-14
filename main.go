package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/a-h/templ"
)

const dataFile = "data.json"

type Ranking struct {
	Id     int         `json:"id"`
	Name   interface{} `json:"name"`
	Points int         `json:"points"`
}

type RankingData struct {
	Total        int       `json:"total"`
	Page         *int      `json:"page"`
	Data         []Ranking `json:"data"`
	ItemsPerPage *int      `json:"items_per_page"`
}

type PointsTable struct {
	pointsDiffLowerBound int
	values               []int
}

var winPointsTable = []PointsTable{
	{pointsDiffLowerBound: 500, values: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	{pointsDiffLowerBound: 400, values: []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3}},
	{pointsDiffLowerBound: 300, values: []int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6}},
	{pointsDiffLowerBound: 200, values: []int{2, 2, 3, 4, 5, 5, 6, 7, 8, 9}},
	{pointsDiffLowerBound: 150, values: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12}},
	{pointsDiffLowerBound: 100, values: []int{3, 4, 5, 6, 8, 9, 10, 11, 13, 15}},
	{pointsDiffLowerBound: 50, values: []int{3, 5, 6, 8, 9, 11, 12, 14, 15, 18}},
	{pointsDiffLowerBound: 25, values: []int{4, 6, 7, 9, 11, 12, 14, 16, 18, 21}},
	{pointsDiffLowerBound: -24, values: []int{4, 6, 8, 10, 12, 14, 16, 18, 20, 24}},
	{pointsDiffLowerBound: -49, values: []int{5, 7, 9, 11, 14, 16, 18, 20, 23, 27}},
	{pointsDiffLowerBound: -99, values: []int{6, 9, 11, 14, 17, 19, 22, 25, 28, 33}},
	{pointsDiffLowerBound: -149, values: []int{7, 11, 14, 18, 21, 25, 28, 32, 35, 42}},
	{pointsDiffLowerBound: -199, values: []int{9, 13, 17, 21, 26, 30, 34, 38, 43, 51}},
	{pointsDiffLowerBound: -299, values: []int{11, 17, 22, 28, 33, 39, 44, 50, 55, 66}},
	{pointsDiffLowerBound: -399, values: []int{15, 23, 30, 38, 45, 53, 60, 68, 75, 90}},
	{pointsDiffLowerBound: -499, values: []int{20, 30, 40, 50, 60, 70, 80, 90, 100, 120}},
	{pointsDiffLowerBound: -100000, values: []int{25, 38, 50, 63, 75, 88, 100, 113, 125, 150}},
}

var losePointsTable = []PointsTable{
	{pointsDiffLowerBound: 500, values: []int{-13, -20, -26, -33, -39}},
	{pointsDiffLowerBound: 400, values: []int{-10, -15, -20, -25, -30}},
	{pointsDiffLowerBound: 300, values: []int{-8, -12, -16, -20, -24}},
	{pointsDiffLowerBound: 200, values: []int{-6, -9, -12, -15, -18}},
	{pointsDiffLowerBound: 150, values: []int{-5, -8, -10, -13, -15}},
	{pointsDiffLowerBound: 100, values: []int{-4, -6, -8, -10, -12}},
	{pointsDiffLowerBound: 50, values: []int{-3, -5, -6, -8, -9}},
	{pointsDiffLowerBound: 25, values: []int{-2, -3, -4, -5, -8}},
	{pointsDiffLowerBound: -24, values: []int{-2, -3, -4, -5, -6}},
	{pointsDiffLowerBound: -49, values: []int{-2, -3, -4, -5, -5}},
	{pointsDiffLowerBound: -99, values: []int{-2, -3, -3, -4, -5}},
	{pointsDiffLowerBound: -149, values: []int{-1, -2, -3, -4, -4}},
	{pointsDiffLowerBound: -199, values: []int{-1, -2, -2, -3, -3}},
	{pointsDiffLowerBound: -299, values: []int{-1, -2, -2, -2, -2}},
	{pointsDiffLowerBound: -399, values: []int{0, -1, -1, -1, -1}},
	{pointsDiffLowerBound: -10000, values: []int{0, 0, 0, 0, 0}},
}

func getValIndexFromMultiplier(multiplier float32) int {
	switch multiplier {
	case 0.5:
		return 0
	case 0.75:
		return 1
	case 1:
		return 2
	case 1.25:
		return 3
	case 1.5:
		return 4
	case 1.75:
		return 5
	case 2:
		return 6
	case 2.25:
		return 7
	case 2.50:
		return 8
	case 3.0:
		return 9
	}

	return -1
}

type Result struct {
	opponentId int
	result     string
}

type TournamentOption struct {
	Id         string
	Name       string
	Multiplier float32
}

var tournamentOptions = map[string]TournamentOption{
	"olympics":                        {Id: "olympics", Name: "Olympic Games", Multiplier: 3.0},
	"worlds":                          {Id: "worlds", Name: "World Championships", Multiplier: 3.0},
	"europe-champs-senior":            {Id: "europe-champs-senior", Name: "European Championships (Senior)", Multiplier: 2.75},
	"europe-games":                    {Id: "europe-games", Name: "European Games", Multiplier: 2.5},
	"olympic-qualification":           {Id: "olympic-qualification", Name: "Olympic Qualification Tournaments", Multiplier: 2.5},
	"commonwealth":                    {Id: "commonwealth", Name: "Commonwealth Games", Multiplier: 2.5},
	"euro-asia":                       {Id: "euro-asia", Name: "Euro-Asia Matches", Multiplier: 2.5},
	"euro-16":                         {Id: "euro-16", Name: "European Top 16", Multiplier: 2.5},
	"euro-10-junior":                  {Id: "euro-10-junior", Name: "European Top 10 (Juniors)", Multiplier: 2.5},
	"wtt-senior":                      {Id: "wtt-senior", Name: "WTT Senior Events", Multiplier: 2.5},
	"ittf-world-team":                 {Id: "ittf-world-team", Name: "ITTF World Team Cup", Multiplier: 2.5},
	"ittf-world":                      {Id: "ittf-world", Name: "ITTF World Cup", Multiplier: 2.5},
	"euro-champs-youth":               {Id: "euro-champs-youth", Name: "European Championships (Youth)", Multiplier: 2.5},
	"commonwealth-championships":      {Id: "commonwealth-championships", Name: "Commonwealth Championships", Multiplier: 2.25},
	"euro-champs-quali":               {Id: "euro-champs-quali", Name: "European Championships Qualification Matches", Multiplier: 2.25},
	"wtt-feeder":                      {Id: "wtt-feeder", Name: "WTT Feeder Series", Multiplier: 2.25},
	"english-senior-champs":           {Id: "english-senior-champs", Name: "English Senior National Championships", Multiplier: 2.25},
	"english-age-champs":              {Id: "english-age-champs", Name: "English Age-Group National Championships", Multiplier: 2},
	"foreign-open":                    {Id: "foreign-open", Name: "Other Foreign Open Championships", Multiplier: 2},
	"wtt-youth":                       {Id: "wtt-youth", Name: "WTT Youth Events", Multiplier: 2},
	"cadet-six":                       {Id: "cadet-six", Name: "Cadet Six Nations", Multiplier: 2},
	"senior-champs-quali":             {Id: "senior-champs-quali", Name: "Senior National Championships Qualifier", Multiplier: 2},
	"4-star":                          {Id: "4-star", Name: "4 Star Open Tournaments (Senior)", Multiplier: 2},
	"national-open":                   {Id: "national-open", Name: "National Open Tournaments (Senior)", Multiplier: 2},
	"4-star-non-senior":               {Id: "4-star-non-senior", Name: "4 Star Open Tournaments (Non-Senior)", Multiplier: 1.75},
	"national-open-non-senior":        {Id: "national-open-non-senior", Name: "National Open Tournaments (Non-Senior)", Multiplier: 1.75},
	"satelite-gp":                     {Id: "satelite-gp", Name: "Satellite Grand Prix", Multiplier: 1.75},
	"13-15-19-national-series":        {Id: "13-15-19-national-series", Name: "U13, U15, U19 National Series", Multiplier: 1.75},
	"21-national-cup":                 {Id: "21-national-cup", Name: "U21 National Cup", Multiplier: 1.75},
	"international-friendly":          {Id: "international-friendly", Name: "Friendly International Matches", Multiplier: 1.5},
	"3-star":                          {Id: "3-star", Name: "3 Star Open Tournaments", Multiplier: 1.5},
	"2-star":                          {Id: "2-star", Name: "2 Star Open Tournaments", Multiplier: 1.5},
	"zonal-opens":                     {Id: "zonal-opens", Name: "Zonal Open Tournaments", Multiplier: 1.5},
	"county-premier":                  {Id: "county-premier", Name: "County Championships (Premier)", Multiplier: 1.5},
	"vetts-regional":                  {Id: "vetts-regional", Name: "VETTS Regional Tournaments", Multiplier: 1.5},
	"elcc-finals":                     {Id: "elcc-finals", Name: "English Leagues Cup Competition - Finals", Multiplier: 1.5},
	"vetts-rating":                    {Id: "vetts-rating", Name: "VETTS Rating Tournaments", Multiplier: 1.5},
	"home-nations-cadet-junior":       {Id: "home-nations-cadet-junior", Name: "Home Nations Cadet & Junior", Multiplier: 1.5},
	"youth-vetts-womens-bcl-prem":     {Id: "youth-vetts-womens-bcl-prem", Name: "Youth BCL, Vetts BCL & Women's BCL Premier Divisions", Multiplier: 1.5},
	"senior-bcl-champs":               {Id: "senior-bcl-champs", Name: "Senior BCL Championship Division", Multiplier: 1.5},
	"county-champs":                   {Id: "county-champs", Name: "County Championships", Multiplier: 1.25},
	"local-open":                      {Id: "local-open", Name: "Local Open Tournaments", Multiplier: 1.25},
	"youth-vetts-womens-bcl-non-prem": {Id: "youth-vetts-womens-bcl-non-prem", Name: "Youth BCL, Vetts BCL, Women's BCL", Multiplier: 1.25},
	"senior-bcl":                      {Id: "senior-bcl", Name: "Senior BCL", Multiplier: 1.25},
	"elcc":                            {Id: "elcc", Name: "English Leagues Cup Competition", Multiplier: 1},
	"jbl":                             {Id: "jbl", Name: "Junior British Clubs League", Multiplier: 1},
	"cbl":                             {Id: "cbl", Name: "Cadet British Clubs League", Multiplier: 1},
	"1-star":                          {Id: "1-star", Name: "1 Star Open Tournaments", Multiplier: 0.75},
}

func parseName(name interface{}) string {
	stringName, ok := name.(string)
	if !ok {
		return ""
	}
	nameSplitByBracket := strings.Split(stringName, "(")
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

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", handler(rankingData))
	http.Handle("/calculate", calculateHandler(rankingData))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

type CalcResponse struct {
	Points int `json:"points"`
}

func calculateHandler(rankingData RankingData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		results := []Result{}
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for i := 0; true; i++ {
			opponentKey := fmt.Sprintf("opponent%d", i)
			winKey := fmt.Sprintf("win%d", i)

			opponent := r.Form.Get(opponentKey)
			if opponent == "" {
				break
			}

			opponentId := -1
			for _, ranking := range rankingData.Data {
				if parseName(ranking.Name) == opponent {
					opponentId = ranking.Id
					break
				}
			}
			if opponentId == -1 {
				continue
			}

			win := r.PostForm.Get(winKey)
			winOrLoss := "loss"
			if win != "" {
				winOrLoss = "win"
			}

			results = append(results, Result{opponentId: opponentId, result: winOrLoss})
		}

		me := r.PostForm.Get("me")
		mePoints := 0
		for _, ranking := range rankingData.Data {
			if parseName(ranking.Name) == me {
				mePoints = ranking.Points
				break
			}
		}

		tournamentKey := r.PostForm.Get("tournament")
		valueIndex := getValIndexFromMultiplier(tournamentOptions[tournamentKey].Multiplier)

		totalPoints := 0
		for _, result := range results {
			for _, ranking := range rankingData.Data {
				if ranking.Id != result.opponentId {
					continue
				}

				opponentsPoints := ranking.Points
				pointsDiff := mePoints - opponentsPoints

				var pointsTable []PointsTable
				if result.result == "win" {
					pointsTable = winPointsTable
				} else {
					pointsTable = losePointsTable
				}

				for _, pointsTableRow := range pointsTable {
					if pointsDiff >= pointsTableRow.pointsDiffLowerBound {
						appropriateValIndex := valueIndex
						if len(pointsTableRow.values) <= valueIndex {
							appropriateValIndex = len(pointsTableRow.values) - 1
						}
						pointsEarned := pointsTableRow.values[appropriateValIndex]
						totalPoints += pointsEarned
						break
					}
				}
			}
		}
		data := CalcResponse{
			Points: totalPoints,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}

func handler(rankingData RankingData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sort.Slice(rankingData.Data, func(i, j int) bool {
			firstName := strings.ToLower(parseName(rankingData.Data[i].Name))
			secondName := strings.ToLower(parseName(rankingData.Data[j].Name))
			return firstName < secondName
		})

		filteredData := []Ranking{}
		for _, ranking := range rankingData.Data {
			stringName, ok := ranking.Name.(string)
			if !ok || stringName == "" {
				continue
			}

			filteredData = append(filteredData, ranking)
		}
		rankingData.Data = filteredData

		tournaments := []TournamentOption{}
		for _, tournament := range tournamentOptions {
			tournaments = append(tournaments, tournament)
		}
		sort.Slice(tournaments, func(i, j int) bool {
			firstTournament := tournaments[i]
			secondTournament := tournaments[j]
			if firstTournament.Multiplier < secondTournament.Multiplier {
				return true
			}

			if firstTournament.Multiplier > secondTournament.Multiplier {
				return false
			}

			return strings.Compare(firstTournament.Name, secondTournament.Name) < 0
		})

		component := page(rankingData.Data, tournaments)
		templ.Handler(component).ServeHTTP(w, r)
	}
}
