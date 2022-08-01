// The guess program requires the player to guess random numbers
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I chose a random number between 1 and 100")
	fmt.Println("You can guess it")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Number of tries remaining: ", 10-guesses, ".")

		fmt.Print("Guess the number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops, this value is too LOWER")
		} else if guess > target {
			fmt.Println("Oops, the value given is too HIGH")
		} else {
			fmt.Println("Great! You played ")
			break
		}

		if !success {
			fmt.Println("Unfortunately, you did not guess. The number you are looking for is:", ".")
		}
	}

}
