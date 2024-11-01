package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ColorMax = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	content, err := os.ReadFile("data")
	if err != nil {
		log.Fatalln(err)
	}

	games := strings.Split(string(content), "\n")
	legitSum := 0
	powerSum := 0

	for i, game := range games {
		gameSets := strings.Split(strings.Trim(strings.Split(string(game), ":")[1], " "), ";")
		isLegit, gamePower := ProcessGame(gameSets)
		powerSum += gamePower
		if isLegit {
			legitSum += (i + 1)
		}
	}
	fmt.Println(legitSum)
	fmt.Println(powerSum)
}

func ProcessGame(gameSets []string) (bool, int) {
	isLegit := true
	colorMin := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range gameSets {
		set = strings.Trim(set, " ")
		r := regexp.MustCompile(`((?P<num>\d*) (?P<color>red|green|blue)+),? ?`)
		matches := r.FindAllStringSubmatch(set, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalln(err)
			}
			color := match[3]

			if num > colorMin[color] {
				colorMin[color] = num
			}

			if num > ColorMax[color] {
				isLegit = false
			}
		}
	}
	return isLegit, colorMin["red"] * colorMin["green"] * colorMin["blue"]
}
