package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/term"
)

type automaton struct {
	ruleSet    map[[3]int]int
	currentGen []int
}

func newAutomaton(rule string) *automaton {
	ruleSet := calculateRuleset(rule)
	firstGen := firstGeneration()
	a := automaton{ruleSet, firstGen}
	return &a
}

func (a *automaton) Next() {
	extGen := append(append(a.currentGen[0:1], a.currentGen...), a.currentGen[len(a.currentGen)-1:]...)
	result := make([]int, len(a.currentGen))
	for i := 0; i < len(a.currentGen); i++ {
		slice := extGen[i : i+3]
		var key [3]int
		copy(key[:], slice[:])
		result[i] = a.ruleSet[key]
	}
	fmt.Print(result)
	a.currentGen = result
	fmt.Print(a.currentGen)
}

func (a *automaton) Show() {
	for i := 0; i < len(a.currentGen); i++ {
		n := a.currentGen[i]
		if n == 0 {
			fmt.Print("\u25a1")
			//fmt.Print(" ")
		} else {
			fmt.Print("\u25a3")
			//fmt.Print("\u25af")
			//fmt.Print("\u2588")
		}
	}
	fmt.Print("\n")
}

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("You have to choose a rule number between 1 and 255")
	}

	automaton := newAutomaton(args[1])

	//ruleset := calculateRuleset(args[1])

	//fmt.Println(ruleset)

	//current := firstGeneration()
	//printGeneration(current)

	for true {
		automaton.Next()
		automaton.Show()
		//current = nextGeneration(current[:], ruleset)
		//printGeneration(current)
		time.Sleep(50 * time.Millisecond)
	}
}

func firstGeneration() []int {
	width, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	gen := make([]int, width)
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
