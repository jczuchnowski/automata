package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/term"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("You have to choose a rule number between 1 and 255")
	}
	println("----")
	println(int(os.Stdout.Fd()))
	println("----")

	ruleset := calculateRuleset(args[1])

	fmt.Println(ruleset)

	//current := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	current := firstGeneration()
	printGeneration(current)

	n := 100
	count := 0

	for count < n {
		current = nextGeneration(current[:], ruleset)
		count = count + 1
		printGeneration(current)
	}
}

func firstGeneration() []int {
	width, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	gen := make([]int, width)
	a := width / 2
	fmt.Println("index %d", a)
	gen[width/2] = 1
	return gen
}

func calculateRuleset(ruleNumber string) map[[3]int]int {
	rule, err := strconv.Atoi(ruleNumber)
	if err != nil || rule < 1 || rule > 255 {
		panic("The rule number has to be number between 1 and 255")
	}

	ruleBinaryStr := fmt.Sprintf("%08b\n", rule)

	chars := []rune(ruleBinaryStr)
	var ints [8]int
	fmt.Println(len(chars))
	for i := len(ints) - 1; i >= 0; i-- {
		fmt.Printf("%d - %d\n", i, chars[i])
		ints[i] = int(chars[i] - '0')
	}

	return map[[3]int]int{
		{1, 1, 1}: ints[0],
		{1, 1, 0}: ints[1],
		{1, 0, 1}: ints[2],
		{1, 0, 0}: ints[3],
		{0, 1, 1}: ints[4],
		{0, 1, 0}: ints[5],
		{0, 0, 1}: ints[6],
		{0, 0, 0}: ints[7],
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

func printGeneration(gen []int) {
	for i := 0; i < len(gen); i++ {
		n := gen[i]
		if n == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("\u2588")
		}
	}
	fmt.Print("\n")
}
