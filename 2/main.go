package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	_, err := strconv.Atoi(str)
	if err == nil {
		return "", errors.New("некорректная строка")
	}

	runes := []rune(str)

	mainStr := ""
	prev := runes[0]

	for i := 0; i < len(runes); i++ {
		num, err := strconv.Atoi(string(runes[i]))

		if err == nil {
			mainStr += strings.Repeat(string(prev), num-1)
			prev = runes[i+1]
		} else {
			mainStr += string(runes[i])
			prev = runes[i]
		}
	}

	return mainStr, nil
}
func main() {
	str1, _ := unpack("a4bc2d5e")
	fmt.Printf("a4bc2d5e = %s\n\n", str1)

	str2, _ := unpack("abcd")
	fmt.Printf("abcd = %s\n\n", str2)

	str3, err := unpack("45")
	fmt.Printf("45 = %s\n", str3)
	fmt.Printf("%s\n\n", err)

	str4, _ := unpack("")
	fmt.Printf("''= %s\n\n", str4)

	str5, _ := unpack("da4bc2d5e")
	fmt.Printf("da4bc2d5e = %s\n\n", str5)

	str6, _ := unpack(" ")
	fmt.Printf("' ' = %s\n", str6)

}
