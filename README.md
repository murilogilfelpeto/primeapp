# Prime Number Checker

This is a simple Go application that checks if a given number is a prime number. The application reads user input from the command line and provides feedback on whether the entered number is prime. The user can exit the application by typing 'q'.

## Features

- Check if a number is prime
- Exit the application by typing 'q'

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. Follow the prompts to enter a number or type 'q' to exit.

## Testing

To run the tests for this project, use the following commands:

1. Run all tests:
    ```sh
    go test ./...
    ```

2. Run tests with coverage:
    ```sh
    go test -cover .
    ```

3. Generate a coverage report:
    ```sh
    go test -coverprofile=coverage.out
    go tool cover -html=coverage.out
    ```

## Project Structure

- `main.go`: Contains the main application logic.
- `main_test.go`: Contains the test cases for the application.

## Dependencies

This project uses the Go standard library and does not require any external dependencies.
