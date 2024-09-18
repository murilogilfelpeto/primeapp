package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)
	go readUserInput(doneChan)
	<-doneChan
	close(doneChan)

	fmt.Println("Goodbye!")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	input := scanner.Text()

	if strings.EqualFold(input, "q") {
		return "", true
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return "Please enter a valid number", false
	}

	_, msg := isPrime(num)
	return msg, false
}

func intro() {
	fmt.Println("Welcome to the prime number checker!")
	fmt.Println("Enter a number to check if it is a prime number, or type 'q' to exit.")
	prompt()
}

func prompt() {
	fmt.Print("Enter a number: ")
}

func isPrime(n int) (bool, string) {
	if n < 2 {
		return false, fmt.Sprintf("%d is not a prime number, by definition", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
