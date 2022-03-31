package main

import (
	"fmt"
	"sort"
	"strings"
)

// SetsOfAnagram function
// the function takes a slice of strings
// returns a map of sets of anagrams
func SetsOfAnagram(words []string) map[string][]string {
	myMap := make(map[string][]string)
	resultMap := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		key := uniqKey(word)

		_, ok := myMap[key]
		if !ok {
			myMap[key] = []string{word}
		} else {
			if !uniqWord(word, myMap[key]) {
				myMap[key] = append(myMap[key], word)
			}
		}
	}

	// sets the key to the first word from the set
	// sorting values of map
	for _, v := range myMap {
		resultMap[v[0]] = v
		sort.Strings(v)
	}

	return resultMap
}

// create uniq key for map
func uniqKey(word string) string {
	sliceLetters := strings.Split(word, "")

	sort.Strings(sliceLetters)

	return strings.Join(sliceLetters, "")
}

// check uniq word in slice of strings
func uniqWord(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	arr := []string{"пятак", "пятак", "пятка", "Яптка", "апятк", "тяпка", "листок", "ислток", "слиток", "столик"}

	result := SetsOfAnagram(arr)

	fmt.Println(result)
}
