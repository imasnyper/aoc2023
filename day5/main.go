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

var Seeds []int64
var SeedToSoil map[int]int64
var SoilToFertilizer map[int64]int64
var FertilizerToWater map[int64]int64
var WaterToLight map[int64]int64
var LightToTemperature map[int64]int64
var TemperatureToHumidity map[int64]int64
var HumidityToLocation map[int64]int64

func stringsToInts(s []string) []int64 {
	i := []int64{}
	for _, x := range s {
		y, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			panic(err)
		}
		i = append(i, y)
	}
	return i
}

func initMap(max int64) map[int64]int64 {
	m := map[int64]int64{}
	for i := int64(0); i < max; i++ {
		m[i] = i
	}

	return m
}

func initValues(maxNum int64, seeds, soil, fertilizer, water, temperature, humidity, location []int64) {
	Seeds = append(Seeds, seeds...)
}

func parseInput(input []byte) {
	s := string(input)
	seedsLabel := "seeds: "
	seeds2SoilLabel := "seed-to-soil map:"
	soil2FertLabel := "soil-to-fertilizer map:"
	fert2WaterLabel := "fertilizer-to-water map:"
	water2LightLabel := "water-to-light map:"
	light2TempLabel := "light-to-temperature map:"
	temp2HumLabel := "temperature-to-humidity map:"
	hum2LocLabel := "humidity-to-location map:"
	numbersRe := regexp.MustCompile(`\d{1,}`)
	numbers := []int64{}

	seedsIndex := strings.Index(s, seedsLabel)
	seeds2SoilIndex := strings.Index(s, seeds2SoilLabel)
	soil2FertIndex := strings.Index(s, soil2FertLabel)
	fert2WaterIndex := strings.Index(s, fert2WaterLabel)
	water2lightIndex := strings.Index(s, water2LightLabel)
	light2TempIndex := strings.Index(s, light2TempLabel)
	temp2HumIndex := strings.Index(s, temp2HumLabel)
	hum2LocIndex := strings.Index(s, hum2LocLabel)

	seedsString := strings.Trim(s[seedsIndex+len(seedsLabel):seeds2SoilIndex], "\n")
	soilString := strings.Trim(s[seeds2SoilIndex+len(seeds2SoilLabel):soil2FertIndex], "\n")
	fertString := strings.Trim(s[soil2FertIndex+len(soil2FertLabel):fert2WaterIndex], "\n")
	waterString := strings.Trim(s[fert2WaterIndex+len(fert2WaterLabel):water2lightIndex], "\n")
	tempString := strings.Trim(s[water2lightIndex+len(water2LightLabel):light2TempIndex], "\n")
	humString := strings.Trim(s[light2TempIndex+len(light2TempLabel):temp2HumIndex], "\n")
	locString := strings.Trim(s[temp2HumIndex+len(temp2HumLabel):hum2LocIndex], "\n")

	seedsNumbers := stringsToInts(numbersRe.FindAllString(seedsString, -1))
	fertNumbers := stringsToInts(numbersRe.FindAllString(fertString, -1))
	soilNumbers := stringsToInts(numbersRe.FindAllString(soilString, -1))
	waterNumbers := stringsToInts(numbersRe.FindAllString(waterString, -1))
	tempNumbers := stringsToInts(numbersRe.FindAllString(tempString, -1))
	humNumbers := stringsToInts(numbersRe.FindAllString(humString, -1))
	locNumbers := stringsToInts(numbersRe.FindAllString(locString, -1))

	numbers = append(seedsNumbers, fertNumbers...)
	numbers = append(numbers, soilNumbers...)
	numbers = append(numbers, waterNumbers...)
	numbers = append(numbers, tempNumbers...)
	numbers = append(numbers, humNumbers...)
	numbers = append(numbers, locNumbers...)

	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Max(numbers))

	fmt.Println(len(numbers))
	fmt.Println(numbers)

	// fmt.Println(seedsString)
	// fmt.Println(soilString)
	// fmt.Println(fertString)
	// fmt.Println(waterString)
	// fmt.Println(tempString)
	// fmt.Println(humString)
	// fmt.Println(locString)
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}
	parseInput(content)
}
