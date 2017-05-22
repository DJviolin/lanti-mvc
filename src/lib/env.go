package lib

import (
	"os"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

// User : return current user's name
var User = os.Getenv("PGUSER")
