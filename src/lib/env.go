package lib

import (
	"log"
	"os"
	"strconv"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

// User : return current user's name
var User = os.Getenv("PGUSER")

// Port : returns the port from env or string
func Port() int {
	// Atoi converts string to int, which returns an error as the second value
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Print("Unable to get env PORT, falling back to :8080")
		port = 8080
	}
	return port
}
