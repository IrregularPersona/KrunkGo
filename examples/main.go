package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/IrregularPersona/KrunkGo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, falling back to system environment variables")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Error: API_KEY is not set in your environment or .env file")
	}

	// debug key
	// fmt.Println(apiKey)

	client := krunkgo.NewClient(apiKey)

	regions := []struct {
		id   int32
		name string
	}{
		{2, "Asia"},
		{3, "Europe"},
		{4, "North America"},
	}

	var allEntries []krunkgo.LeaderboardEntry

	for _, region := range regions {
		response, err := client.GetLeaderboard(region.id, 1)
		if err != nil {
			fmt.Printf("Failed to fetch %s leaderboard: %v\n", region.name, err)
			continue
		}

		if response != nil && len(response.LREntries) > 0 {
			allEntries = append(allEntries, response.LREntries...)
		}
	}

	if len(allEntries) == 0 {
		fmt.Println("No leaderboard entries found.")
		return
	}

	// Sort by MMR descending
	slices.SortFunc(allEntries, func(a, b krunkgo.LeaderboardEntry) int {
		if b.LEMMR != a.LEMMR {
			return int(b.LEMMR - a.LEMMR)
		}
		return 0
	})

	// Dedup for multiple entries of the same player name
	if len(allEntries) > 0 {
		keepIdx := 0
		for i := 1; i < len(allEntries); i++ {
			if allEntries[i].LEPlayerName != allEntries[keepIdx].LEPlayerName {
				keepIdx++
				allEntries[keepIdx] = allEntries[i]
			}
		}
		allEntries = allEntries[:keepIdx+1]
	}

	limit := 10
	if len(allEntries) < limit {
		limit = len(allEntries)
	}

	topEntries := allEntries[:limit]

	fmt.Println("=== Top 10 Players Overall (By MMR) ===")
	for i, entry := range topEntries {
		fmt.Printf(
			"%2d. %-20s | MMR: %-5d | Win/Loss: %d/%d\n",
			i+1,
			entry.LEPlayerName,
			entry.LEMMR,
			entry.LEWins,
			entry.LELosses,
		)
	}
}
