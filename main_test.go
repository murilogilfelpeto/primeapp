package main

import (
	"io"
	"os"
	"testing"
)

//to run test with coverage we can use the following command:
//go test -cover .
// go tool cover -html=coverage.out

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
