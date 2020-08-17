package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(DigitalRoot(493193))
}

func DigitalRoot(n int) int {
	stringConvert := strconv.Itoa(n)

	var numberArray = strings.Split(stringConvert, "")
	var sum = 0
	for i := 0; i < len(numberArray); i++ {
		if s, err := strconv.ParseInt(numberArray[i], 10, 32); err == nil {
			sum = sum + int(s)
		}

	}
	if sum >= 10 {
		return DigitalRoot(sum)
	}
	return sum
}
