package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func parseRaces(data string) []Race {
	groupExpression := regexp.MustCompile(`Time:\s*((\d+\s*)+)\nDistance:\s*((\d+\s*)+)`)
	numberExpression := regexp.MustCompile(`\d{1,}`)
	matches := groupExpression.FindStringSubmatch(data)
	timesString := matches[1]
	distancesString := matches[3]
	times := []int{}
	distances := []int{}
	for i, numbers := range []string{timesString, distancesString} {
		matches = numberExpression.FindAllString(numbers, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			if i == 0 {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}
		}
	}
	if len(times) != len(distances) {
		panic("Times and distances are not the same length.")
	}
	races := []Race{}
	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}
	return races
}

func doRace(race Race) int {
	time := race.time
	distance := race.distance

	waysToWin := 0

	for speed := 0; speed < time; speed++ {
		if speed*(time-speed) > distance {
			waysToWin++
		}
	}

	return waysToWin
}

func doRaces(races []Race) int {
	result := 1
	for _, race := range races {
		result *= doRace(race)
	}
	return result
}

func partTwoParse(data string) Race {
	expression := regexp.MustCompile(`Time:\s*((?:\d\s*)+)\nDistance:\s*((?:\d\s*)+)`)
	result := expression.FindStringSubmatch(data)
	fmt.Println(result)
	time, err := strconv.Atoi(strings.Replace(result[1], " ", "", -1))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Replace(result[2], " ", "", -1))
	if err != nil {
		panic(err)
	}
	return Race{time, distance}
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}
	data := string(content)
	// Part 1
	// races := parseRaces(data)
	// fmt.Printf("Races: %v\n", races)
	// result := doRaces(races)
	// fmt.Println(result)

	// Part 2
	race := partTwoParse(data)
	result := doRace(race)
	fmt.Println(result)
}
