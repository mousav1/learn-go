package main

import (
	"errors"
	"fmt"
)

// Custom error type
type DatabaseError struct {
	Code    int
	Message string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("DatabaseError: code=%d message=%s", e.Code, e.Message)
}

// Another error type
type NetworkError struct {
	Status int
	URL    string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("NetworkError: status=%d url=%s", e.Status, e.URL)
}

// Simulate database operation
func queryDatabase() error {
	return &DatabaseError{
		Code:    500,
		Message: "connection failed",
	}
}

// Simulate network operation
func callAPI() error {
	return &NetworkError{
		Status: 404,
		URL:    "https://example.com",
	}
}

// Modern Go 1.26 approach
func handleWithAsType(err error) {

	fmt.Println("\n=== Using errors.AsType (Go 1.26) ===")

	if dbErr, ok := errors.AsType[*DatabaseError](err); ok {

		fmt.Println("Database error detected")
		fmt.Println("Code:", dbErr.Code)
		fmt.Println("Message:", dbErr.Message)

		return
	}

	if netErr, ok := errors.AsType[*NetworkError](err); ok {

		fmt.Println("Network error detected")
		fmt.Println("Status:", netErr.Status)
		fmt.Println("URL:", netErr.URL)

		return
	}

	fmt.Println("Unknown error:", err)
}

// Old approach
func handleWithAs(err error) {

	fmt.Println("\n=== Using errors.As (Old way) ===")

	var dbErr *DatabaseError

	if errors.As(err, &dbErr) {

		fmt.Println("Database error detected")
		fmt.Println("Code:", dbErr.Code)
		fmt.Println("Message:", dbErr.Message)

		return
	}

	var netErr *NetworkError

	if errors.As(err, &netErr) {

		fmt.Println("Network error detected")
		fmt.Println("Status:", netErr.Status)
		fmt.Println("URL:", netErr.URL)

		return
	}

	fmt.Println("Unknown error:", err)
}

func main() {

	fmt.Println("=== Database Example ===")

	err := queryDatabase()

	handleWithAsType(err)

	handleWithAs(err)

	fmt.Println("\n=== Network Example ===")

	err = callAPI()

	handleWithAsType(err)

	handleWithAs(err)
}
