package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sourceToDestination(source int, destinations [][]int) int {
	// MAP ARRAY
	// ARR[0] to ARR[0] + ARR[2]  = DESTINATION RANGE
	// ARR [1] to ARR[1] + ARR[2] = SOURCE RANGE
	// IF SEED OR START POS < ARR[0] THEN DESTINATION = SEED OR START POS
	// START WITH SEED VALUE, CHECK FOR SOURCE RANGE THEN DEPOSIT AT
	// SAME RELATIVE POSITION FROM THE DESTINATION RANGE
	var n int
	nset := false
	for i := 0; i < len(destinations); i++ {
		// log.Println("source is", source, ", destination is:", destinations[i][1], "to", destinations[i][1]+destinations[i][2])
		if source >= destinations[i][1] && source < destinations[i][1]+destinations[i][2] {
			n = source - destinations[i][1]
			n = destinations[i][0] + n
			nset = true
		}
	}
	if !nset {
		n = source
	}
	return n
}

func stringToArray(str string) []int {
	numstr := strings.Split(str, " ")
	var nums []int
	for _, num := range numstr {
		t, err := strconv.Atoi(num)
		if err == nil {
			nums = append(nums, t)
		}
	}
	return nums
}

func LinesToArray(searchstr string, lines []string) [][]int {
	start := slices.Index(lines, searchstr) + 1
	var tempArray [][]int
	for _, line := range lines[start:] {
		if strings.TrimSpace(line) == "" {
			break
		}
		nums := stringToArray(line)
		tempArray = append(tempArray, nums)
	}
	return tempArray
}

func partOne(lines []string) {
	seeds := stringToArray(lines[0][7:])
	seedsToSoilMapArray := LinesToArray("seed-to-soil map:", lines)
	soilToFertiliserMapArray := LinesToArray("soil-to-fertilizer map:", lines)
	fertiliserToWaterMapArray := LinesToArray("fertilizer-to-water map:", lines)
	waterToLightMapArray := LinesToArray("water-to-light map:", lines)
	lightToTemperatureMapArray := LinesToArray("light-to-temperature map:", lines)
	temperatureToHumidityMapArray := LinesToArray("temperature-to-humidity map:", lines)
	humidityToLocationMapArray := LinesToArray("humidity-to-location map:", lines)
	// log.Println("seeds:", seeds)
	// log.Println("seed-to-soil map:", seedsToSoilMapArray)
	// log.Println("fertilizer-to-water map:", soilToFertiliserMapArray)
	// log.Println("water-to-light map:", fertiliserToWaterMapArray)
	// log.Println("light-to-temperature map:", waterToLightMapArray)
	// log.Println("temperature-to-humidity map:", lightToTemperatureMapArray)
	// log.Println("temperature-to-humidity map:", temperatureToHumidityMapArray)
	// log.Println("humidity-to-location map:", humidityToLocationMapArray)
	var locations []int
	for _, seed := range seeds {
		// fmt.Printf("Seed: %d", seed)
		n := sourceToDestination(seed, seedsToSoilMapArray)
		// fmt.Printf(", soil: %d", n)
		n = sourceToDestination(n, soilToFertiliserMapArray)
		// fmt.Printf(", fertilizer: %d", n)
		n = sourceToDestination(n, fertiliserToWaterMapArray)
		// fmt.Printf(", water: %d", n)
		n = sourceToDestination(n, waterToLightMapArray)
		// fmt.Printf(", light: %d", n)
		n = sourceToDestination(n, lightToTemperatureMapArray)
		// fmt.Printf(", temperature: %d", n)
		n = sourceToDestination(n, temperatureToHumidityMapArray)
		// fmt.Printf(", humidity: %d", n)
		n = sourceToDestination(n, humidityToLocationMapArray)
		// fmt.Printf(", location: %d", n)
		// fmt.Println()
		locations = append(locations, n)
	}
	log.Println("Lowest location: ", slices.Min(locations))
}

func ReadMaps(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var lines []string
	for fScanner.Scan() {
		line := fScanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	// filePath := "./d05test.txt"
	filePath := "./d05input.txt"

	lines := ReadMaps(filePath)
	partOne(lines)
}
