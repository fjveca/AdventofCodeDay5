package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// creates a map where the index of the map is the page number and the content for that index is a string with all the page numbers that
// should go after the pagenumber used for the index of the MAP, thus creating a ruleset for all the pages
func generate_ruleset(rules []string) map[int]string {
	ruleset := make(map[int]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		index, _ := strconv.Atoi(parts[0])
		ruleset[index] = ruleset[index] + parts[1] + ","
	}
	return ruleset
}

func add_middle_page_num(update string, ordered []string, sum int) int { //returns the value of the middle page of the update
	index := len(strings.Split(update, ","))
	toRound := index / 2
	middle := math.Round(float64(toRound))
	amount, _ := strconv.Atoi(string(ordered[int(middle)]))
	return amount
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
		orderedInput := strings.Split(inputs[i], ",")  //create a version of the list that can be updated into an ordered form of the list
		sort.Slice(orderedInput, func(j, k int) bool { //orders the update according to the ruleset
			index, _ := strconv.Atoi(orderedInput[j])
			if strings.Contains(ruleset[index], orderedInput[k]) { // if the next item is found on the ruleset of the current item, then they are in order
				return true
			}
			return false // if the next item is not found on the ruleset of the current item, let the sorting algorithm work

		})
		verificationString := ""

		for i := 0; i < len(orderedInput); i++ { //create a string with the sorted slice of the update
			if i == len(orderedInput)-1 {
				verificationString += orderedInput[i]
			} else {
				verificationString += orderedInput[i] + ","
			}
		}

		if inputs[i] == verificationString { //if the inputted update is equal to the string of the ordered slice then add the page number to the sum of the ordered updates
			total += add_middle_page_num(inputs[i], orderedInput, total)

		} else { //add the sum of the middle pages ordered version of the disordered updates
			TotalforDisordered += add_middle_page_num(inputs[i], orderedInput, TotalforDisordered)
		}

	}
	fmt.Println("Total sum of middle values for Ordered lists", total)
	fmt.Println("Total sum of middle values for not ordered lists:", TotalforDisordered)
}
