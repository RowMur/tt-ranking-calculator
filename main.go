package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

const dataFile = "data.json"

type Ranking struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
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

type Result struct {
	opponentId int
	result     string
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

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", handler(rankingData))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(rankingData RankingData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results := []Result{}
		for i := 0; true; i++ {
			opponentKey := fmt.Sprintf("opponent%d", i)
			winKey := fmt.Sprintf("win%d", i)

			opponent := r.URL.Query().Get(opponentKey)
			if opponent == "" {
				break
			}
			opponentId, err := strconv.Atoi(opponent)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("failed to parse opponent ID: %v", err)))
				return
			}

			win := r.URL.Query().Get(winKey)
			winOrLoss := "loss"
			if win != "" {
				winOrLoss = "win"
			}

			results = append(results, Result{opponentId: opponentId, result: winOrLoss})
		}

		meId := r.URL.Query().Get("me")
		mePoints := 0
		for _, ranking := range rankingData.Data {
			if strconv.Itoa(ranking.Id) == meId {
				mePoints = ranking.Points
				break
			}
		}

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
						pointsEarned := pointsTableRow.values[3]
						totalPoints += pointsEarned
						break
					}
				}
			}
		}

		sort.Slice(rankingData.Data, func(i, j int) bool {
			firstName := strings.ToLower(rankingData.Data[i].Name)
			secondName := strings.ToLower(rankingData.Data[j].Name)
			return firstName < secondName
		})

		component := page(rankingData.Data, totalPoints, &meId)
		templ.Handler(component).ServeHTTP(w, r)
	}
}
