package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	// tokenRe := regexp.MustCompile(`(one)*(two)*(three)*(four)*(five)*(six)*(seven)*(eight)*(nine)*([\d])?`)
	tokenRe := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)

	for _, line := range lines {
		leftMatchNumber := []int{1000, 1000}
		rightMatchNumber := []int{-1, -1}

		for i := 0; i < len(line); i++ {
			match := tokenRe.FindStringIndex(line[i:])
			if len(match) > 0 && match[0] != -1 {
				var num int
				strNum := line[match[0]+i : match[1]+i]
				if match[1]-match[0] > 1 {
					num = numbers[strNum]
				} else {
					num, err = strconv.Atoi(strNum)
				}
				if err != nil {
					panic(err)
				}
				if leftMatchNumber[0] > match[0]+i {
					leftMatchNumber = []int{match[0], num}
				}
				if rightMatchNumber[0] < match[0]+i {
					rightMatchNumber = []int{match[0], num}
				}
			}
		}
		if leftMatchNumber[1] > 0 && leftMatchNumber[1] < 10 && rightMatchNumber[1] > 0 && rightMatchNumber[1] < 10 {
			// fmt.Printf("%d%d %d\n", leftMatchNumber[1], rightMatchNumber[1], 10*leftMatchNumber[1]+rightMatchNumber[1])
			sum += 10*leftMatchNumber[1] + rightMatchNumber[1]
		}
	}

	fmt.Println("Sum: ", sum)
}

// func checkIndex(i int, indicies map[int]int, greater bool) bool {
// 	x := true
// 	for index := range indicies {
// 		if greater {
// 			if i > index {
// 				x = false
// 				break
// 			}
// 		} else {
// 			if i < index {
// 				x = false
// 				break
// 			}
// 		}
// 	}
// 	return x
// }
