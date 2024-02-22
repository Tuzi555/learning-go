package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []string
}
type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	if firstTeamScore == secondTeamScore {
		return
	}
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeamName] += 1
	} else {
		l.Wins[secondTeamName] += 1
	}
}

func (l League) Ranking() []string {
	keys := make([]string, 0, len(l.Wins))
	for key := range l.Wins {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return l.Wins[keys[i]] > l.Wins[keys[j]]
	})
	return keys
}

func RankPrinter(ranker Ranker, writer io.Writer) error {
	for _, ranking := range ranker.Ranking() {
		_, err := io.WriteString(writer, ranking+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)
	err := RankPrinter(l, os.Stdout)
	if err != nil {
		return
	}
}
