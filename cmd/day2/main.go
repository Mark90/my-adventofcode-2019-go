package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetIntegersFromFirstLine opens the given file and returns integers from the first
func GetIntegersFromFirstLine(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	var numbers []int
	for _, val := range strings.Split(input, ",") {
		newv, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, newv)
	}
	return numbers
}

func algorithm(values []int, replace1, replace2 int) int {
	values[1] = replace1
	values[2] = replace2

	var op, pos int
	for {
		op = values[pos]
		if op == 99 {
			return values[0]
		}
		if values[pos] == 1 {
			values[values[pos+3]] = values[values[pos+1]] + values[values[pos+2]]
		} else {
			values[values[pos+3]] = values[values[pos+1]] * values[values[pos+2]]
		}
		pos += 4
	}
}

func getCopy(numbers []int) []int {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	return numbersCopy
}

func part1(numbers []int) {
	fmt.Println("[part1] Index 0:", algorithm(numbers, 12, 2))
}

func part2(numbers []int) {
	targetOutput := 19690720

	baseInput1, baseInput2 := 0, 0
	baseOutput := algorithm(getCopy(numbers), baseInput1, baseInput2)
	totalDelta := targetOutput - baseOutput

	deltaInput1 := algorithm(getCopy(numbers), baseInput1+1, baseInput2) - baseOutput
	deltaInput2 := algorithm(getCopy(numbers), baseInput1, baseInput2+1) - baseOutput

	multDelta1 := totalDelta / deltaInput1
	multDelta2 := int((totalDelta % deltaInput1) / deltaInput2)

	noun := baseInput1 + multDelta1
	verb := baseInput2 + multDelta2
	if algorithm(getCopy(numbers), noun, verb) != targetOutput {
		panic("aaaa")
	}
	fmt.Println("[part2] 100 * noun + verb =", 100*noun+verb)
}

func main() {
	numbers := GetIntegersFromFirstLine("input/day2.txt")
	part1(getCopy(numbers))
	part2(getCopy(numbers))
}
