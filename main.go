package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func generate_ruleset(rules []string) map[int]string {
	ruleset := make(map[int]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		index, _ := strconv.Atoi(parts[0])
		ruleset[index] = ruleset[index] + parts[1] + ","
	}
	return ruleset
}

func main() {
	total := 0
	TotalforDisordered := 0
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	parts := strings.Split(string(content), "\n\n")

	segment1 := parts[0]
	segment2 := parts[1]
	rules := strings.Split(segment1, "\n")
	inputs := strings.Split(segment2, "\n")
	ruleset := generate_ruleset(rules)
	for i := 0; i < len(inputs); i++ {
		orderedInput := strings.Split(inputs[i], ",")
		sort.Slice(orderedInput, func(j, k int) bool {
			index, _ := strconv.Atoi(orderedInput[j])
			if strings.Contains(ruleset[index], orderedInput[k]) {
				return true
			}
			return false

		})
		verificationString := ""

		for i := 0; i < len(orderedInput); i++ {
			if i == len(orderedInput)-1 {
				verificationString += orderedInput[i]
			} else {
				verificationString += orderedInput[i] + ","
			}
		}

		if inputs[i] == verificationString {
			index := len(strings.Split(inputs[i], ","))
			toRound := index / 2
			middle := math.Round(float64(toRound))
			amount, _ := strconv.Atoi(string(orderedInput[int(middle)]))
			total += amount

		} else {
			index := len(strings.Split(inputs[i], ","))
			toRound := index / 2
			middle := math.Round(float64(toRound))
			amount, _ := strconv.Atoi(string(orderedInput[int(middle)]))
			TotalforDisordered += amount
		}

	}
	fmt.Println("Total sum of middle values for Ordered lists",total)
	fmt.Println("Total sum of middle values for not ordered lists:",TotalforDisordered)
}
