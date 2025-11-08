package config

import "os"

// We return the secret from the env and cast it into a slice of bytes
var JwtSecret = []byte(func() string {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		// Fallback
		s = "dev-secret"
	}
	return s
}())
