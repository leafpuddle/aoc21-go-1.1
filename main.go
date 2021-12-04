package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(inputf)

	inputs := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	result := 0

	for i := 0; i < len(inputs); i++ {
		if i == 0 {
			continue
		}

		a, err := strconv.ParseInt(inputs[i], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.ParseInt(inputs[i-1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if a > b {
			result++
		}
	}

	fmt.Print(result)
}
