package main

import (
	"reflect"
	"testing"
)

func TestNextGeneration(t *testing.T) {
	//given
	testRuleSet := map[[3]int]int{
		{1, 1, 1}: 0,
		{1, 1, 0}: 0,
		{1, 0, 1}: 0,
		{1, 0, 0}: 0,
		{0, 1, 1}: 1,
		{0, 1, 0}: 0,
		{0, 0, 1}: 0,
		{0, 0, 0}: 0,
	}

	testGeneration := [5]int{0, 0, 1, 0, 0}

	//when
	result := nextGeneration(testGeneration[:], testRuleSet)

	//then
	expected := [5]int{0, 0, 0, 0, 0}
	if !reflect.DeepEqual(result, expected[:]) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}
}
