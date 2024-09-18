package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

//to run test with coverage we can use the following command:
//go test -cover .
//go test -coverprofile=coverage.out && go tool cover -html=coverage.out

//Running all tests: go test -v .

//Running a single test: go test -run Test_isPrime

//Running groups test (test suite): go test -run <Test_Suite>

// Test must start as Test_ followed by what we are testing
func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("isPrime(0) = %t; want false", result)
	}

	if msg != "0 is not a prime number, by definition" {
		t.Errorf("isPrime(0) = %s; want 0 is not a prime number, by definition", msg)
	}
}

func Test_isPrimeTableTest(t *testing.T) {
	testCases := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{
			name:     "prime number",
			testNum:  7,
			expected: true,
			msg:      "7 is a prime number",
		},
		{
			name:     "not prime number",
			testNum:  10,
			expected: false,
			msg:      "10 is not a prime number because it is divisible by 2",
		},
		{
			name:     "number lower than 2",
			testNum:  1,
			expected: false,
			msg:      "1 is not a prime number, by definition",
		},
		{
			name:     "number lower than 0",
			testNum:  -1,
			expected: false,
			msg:      "-1 is not a prime number, by definition",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, msg := isPrime(tc.testNum)
			if result != tc.expected {
				t.Errorf("isPrime(%d) = %t; want %t", tc.testNum, result, tc.expected)
			}

			if msg != tc.msg {
				t.Errorf("isPrime(%d) = %s; want %s", tc.testNum, msg, tc.msg)
			}
		})
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	//create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	//close write pipe
	_ = w.Close()

	os.Stdout = oldOut

	//read from read pipe
	out, _ := io.ReadAll(r)

	if string(out) != "Enter a number: " {
		t.Errorf("prompt() = %s; want Enter a number: ", out)
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	//create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	//close write pipe
	_ = w.Close()

	os.Stdout = oldOut

	//read from read pipe
	out, _ := io.ReadAll(r)

	outString := string(out)
	if !strings.Contains(outString, "Welcome to the prime number checker!") {
		t.Errorf("intro() = %s; want Welcome to the prime number checker!", out)
	}

	if !strings.Contains(outString, "Enter a number to check if it is a prime number, or type 'q' to exit.") {
		t.Errorf("intro() = %s; want Enter a number to check if it is a prime number, or type 'q' to exit.", out)
	}
}

func Test_checkNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		done     bool
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "Please enter a valid number",
			done:     false,
		},
		{
			name:     "not a number",
			input:    "a",
			expected: "Please enter a valid number",
			done:     false,
		},
		{
			name:     "prime number",
			input:    "7",
			expected: "7 is a prime number",
			done:     false,
		},
		{
			name:     "quit application upper case",
			input:    "Q",
			expected: "",
			done:     true,
		},
		{
			name:     "quit application",
			input:    "q",
			expected: "",
			done:     true,
		},
		{
			name:     "decimal number",
			input:    "3.14",
			expected: "Please enter a valid number",
			done:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			reader := bufio.NewScanner(input)
			res, done := checkNumbers(reader)

			if !strings.EqualFold(res, tc.expected) {
				t.Errorf("checkNumbers(%s) = %s; want %s", tc.input, res, tc.expected)
			}

			if done != tc.done {
				t.Errorf("checkNumbers(%s) = %t; want %t", tc.input, done, tc.done)
			}
		})
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer
	stdin.Write([]byte("7\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

	//no need to test the output, we are testing the input
}
