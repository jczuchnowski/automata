package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("You have to choose a rule number between 1 and 255")
	}

	rule, err := strconv.Atoi(os.Args[1])
	if err != nil || rule < 1 || rule > 255 {
		panic("The rule number has to be number between 1 and 255")
	}

	ruleBinaryStr := fmt.Sprintf("%08b\n", rule)

	fmt.Println(ruleBinaryStr)

	chars := []rune(ruleBinaryStr)
	var ints [8]int
	fmt.Println(len(chars))
	for i := len(ints) - 1; i >= 0; i-- {
		fmt.Printf("%d - %d\n", i, chars[i])
		ints[i] = int(chars[i] - '0')
	}

	ruleMapping := map[[3]int]int{
		{1, 1, 1}: ints[0],
		{1, 1, 0}: ints[1],
		{1, 0, 1}: ints[2],
		{1, 0, 0}: ints[3],
		{0, 1, 1}: ints[4],
		{0, 1, 0}: ints[5],
		{0, 0, 1}: ints[6],
		{0, 0, 0}: ints[7],
	}
	fmt.Println(ruleMapping)

	current := []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
	fmt.Println(current)

	// zero1 := append(append(zero[0:1], zero[0:]...), zero[len(zero)-1:]...)
	// for i := 0; i < len(zero1); i++ {
	// 	slice := zero1[i : i+3]
	// 	fmt.Println(slice)
	// }

	//next := nextGeneration(zero[:], ruleMapping)

	n := 100
	count := 0

	for count < n {
		current = nextGeneration(current[:], ruleMapping)
		count = count + 1
		fmt.Println(current)
	}
}

func nextGeneration(gen []int, ruleset map[[3]int]int) []int {
	extGen := append(append(gen[0:1], gen...), gen[len(gen)-1:]...)
	result := make([]int, len(gen))
	for i := 0; i < len(gen); i++ {
		slice := extGen[i : i+3]
		var key [3]int
		copy(key[:], slice[:])
		result[i] = ruleset[key]
	}
	return result
}
