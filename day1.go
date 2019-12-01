package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetLinesAsIntegers opens the given file and returns lines as integers
func GetLinesAsIntegers(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var values []int
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, input)
	}
	return values
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func part1(modules []int) {
	totalfuel := 0
	for _, mass := range modules {
		totalfuel += calculateFuel(mass)
	}
	fmt.Println("[part1] Total fuel:", totalfuel)
}

func part2(modules []int) {
	totalfuel := 0
	for _, mass := range modules {
		for fuel := calculateFuel(mass); fuel > 0; fuel = calculateFuel(fuel) {
			totalfuel += fuel
		}
	}
	fmt.Println("[part2] Total fuel:", totalfuel)
}

func main() {
	values := GetLinesAsIntegers("input/day1.txt")
	part1(values)
	part2(values)
}
