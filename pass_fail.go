// The pass_fail program tells you if the given result means pass
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the result: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passed"
	} else {
		status = "failing to pass"
	}
	fmt.Println("Result", grade, "means", "status", ".")
}
