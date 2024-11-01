package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id             int
	total          int
	winningNumbers []int
	cardNumbers    []int
}

func stringsToInts(s []string) []int {
	i := []int{}
	for _, x := range s {
		y, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		i = append(i, y)
	}
	return i
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(content), "\n")
	numberRe := regexp.MustCompile(`\d{1,2}`)
	scratchCards := []ScratchCard{}

	for i, card := range lines {
		winningNumbersString := card[strings.Index(card, ":"):strings.Index(card, "|")]
		cardNumbersString := card[strings.Index(card, "|"):]
		winningNumbers := stringsToInts(numberRe.FindAllString(winningNumbersString, -1))
		cardNumbers := stringsToInts(numberRe.FindAllString(cardNumbersString, -1))

		slices.Sort(winningNumbers)
		slices.Sort(cardNumbers)

		sc := ScratchCard{
			i + 1,
			1,
			winningNumbers,
			cardNumbers,
		}

		scratchCards = append(scratchCards, sc)
	}

	cardsProcessed := 0
	for i, card := range scratchCards {
		for range card.total {
			cardsProcessed += 1
			matches := 0
		winning:
			for _, w := range card.winningNumbers {
				for _, c := range card.cardNumbers {
					if c > w {
						continue winning
					}
					if w == c {
						matches++
						continue winning
					}
				}
			}

			for j := range matches {
				scratchCards[i+1+j].total += 1
			}
			// if matches < 3 {
			// 	points += matches
			// } else {
			// 	newPoints := int(math.Pow(2, float64(matches)-1))
			// 	points += newPoints
			// }
		}
	}

	fmt.Println("Cards Processed: ", cardsProcessed)
}
