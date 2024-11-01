package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type PartNumber struct {
	id     int
	row    int
	start  int
	end    int
	number int
}

var symbols = ""
var partNumbers = []*PartNumber{}

func newPart(row, start, end int, line string) *PartNumber {
	number, err := strconv.Atoi(line[start : end+1])
	if err != nil {
		panic(err)
	}
	return &PartNumber{
		id:     (row * 3) + (start * 2),
		row:    row,
		start:  start,
		end:    end,
		number: number,
	}
}

func checkPart(row, col int) *PartNumber {
	for _, part := range partNumbers {
		if part.row == row && col >= part.start && col <= part.end {
			return part
		}
	}
	return nil
}

func isSymbol(r rune) bool {
	for _, x := range symbols {
		if r == x {
			return true
		}
	}

	return false
}

func makeRange(start, end int) []int {
	x := []int{}
	for i := range end - start {
		x = append(x, i+start)
	}
	return x
}

func getGearRatio(lines []string) int64 {
	escaped := regexp.QuoteMeta(symbols)
	escaped = strings.Replace(escaped, "-", "", -1) + "-"
	symbolsRe := regexp.MustCompile("[" + escaped + "]")

	ratioSum := int64(0)

	for row, line := range lines {
		symbolIndicies := symbolsRe.FindAllStringIndex(line, -1)
		rangeDelta := makeRange(-1, 2)

		for _, i := range symbolIndicies {
			symbolNeighbors := []*PartNumber{}
			for _, r := range rangeDelta {
				for _, c := range rangeDelta {
					part := checkPart(row+r, i[0]+c)
					if part != nil && !slices.Contains(symbolNeighbors, part) {
						symbolNeighbors = append(symbolNeighbors, part)
					}
				}
			}

			if len(symbolNeighbors) == 2 {
				ratioSum += int64(symbolNeighbors[0].number) * int64(symbolNeighbors[1].number)
			}

			// if row > 0 && unicode.IsDigit(rune(lines[row-1][i[0]])) {
			// 	symbolNeighbors = append(symbolNeighbors, )
			// } else if row < len(lines)-2 && unicode.IsDigit(rune(lines[row+1][i[0]])) {
			// 	symbolNeighbors += 1
			// } else if i[0] > 0 && unicode.IsDigit(rune(lines[row][i[0]-1])) {
			// 	symbolNeighbors += 1
			// } else if i[0] < len(line)-2 && unicode.IsDigit(rune(lines[row][i[0]+1])) {
			// 	symbolNeighbors += 1
			// }
		}
	}
	return ratioSum
}

func getSum(lines []string) int {
	numbersRe := regexp.MustCompile(`\d+`)

	sum := 0
	numbers := []int{}

	for row, line := range lines {
		numIndicies := numbersRe.FindAllStringIndex(line, -1)
		for _, iRange := range numIndicies {
			addNumber := false
			number := ""

			r := []int{}
			if iRange[0] > 0 && iRange[1] < len(line)-1 {
				r = makeRange(iRange[0]-1, iRange[1]+1)
			} else if iRange[0] > 0 {
				r = makeRange(iRange[0]-1, iRange[1])
			} else if iRange[1] < len(line)-1 {
				r = makeRange(iRange[0], iRange[1]+1)
			} else {
				r = makeRange(iRange[0], iRange[1])
			}

			for _, col := range r {
				if col >= iRange[0] && col <= iRange[1]-1 {
					number += string(line[col])
				}
				if (col == iRange[0] && col > 0 && isSymbol(rune(lines[row][col-1]))) ||
					(col == iRange[1]-1 && col < len(line)-2 && isSymbol(rune(lines[row][col+1]))) ||
					(row > 0 && isSymbol(rune(lines[row-1][col]))) ||
					(row < len(lines)-2 && isSymbol(rune(lines[row+1][col]))) {
					if !addNumber {
						partNumbers = append(partNumbers, newPart(row, iRange[0], iRange[1]-1, line))
					}
					addNumber = true
				}

			}

			if addNumber {
				n, err := strconv.Atoi(number)
				numbers = append(numbers, n)
				if err != nil {
					panic(err)
				}
				sum += n
			}
		}

	}
	return sum
}

func main() {
	lines := ParseInput()
	fmt.Println("Sum of connected numbers: ", getSum(lines))
	fmt.Println("Gear ratio: ", getGearRatio(lines))
}

func ParseInput() []string {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		for _, r := range line {
			if !unicode.IsDigit(r) && r != '.' && !strings.Contains(symbols, string(r)) {
				symbols += string(r)
			}
		}
	}

	fmt.Println("Symbols: ", symbols)

	return lines
}
