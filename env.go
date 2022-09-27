package env

import (
	"os"
	"strconv"
)

// String retrieves the value of the environment variable named by the key. If
// the variable is present in the environment, its value (which may be empty) is
// returned, otherwise fallback is returned.
func String(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

// Bool retrieves the value of the environment variable named by the key, parses
// the value as a boolean, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Bool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return res
}

// Int retrieves the value of the environment variable named by the key, parses
// the value as an integer, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Int(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseInt(value, 10, strconv.IntSize)
	if err != nil {
		return fallback
	}

	return int(res)
}

// Int64 retrieves the value of the environment variable named by the key, parses
// the value as a 64-bit integer, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Int64(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}

	return res
}

// Uint retrieves the value of the environment variable named by the key, parses
// the value as an unsigned integer, and returns the result. If the variable is
// not present or its value cannot be parsed, fallback is returned.
func Uint(key string, fallback uint) uint {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseUint(value, 10, strconv.IntSize)
	if err != nil {
		return fallback
	}

	return uint(res)
}

// Uint64 retrieves the value of the environment variable named by the key,
// parses the value as an unsigned 64-bit integer, and returns the result. If the
// variable is not present or its value cannot be parsed, fallback is returned.
func Uint64(key string, fallback uint64) uint64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fallback
	}

	return res
}

// Float32 retrieves the value of the environment variable named by the key,
// parses the value as a floating-point number, and returns the result. If the
// variable is not present or its value cannot be parsed, fallback is returned.
func Float32(key string, fallback float32) float32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return fallback
	}

	return float32(res)
}

// Float64 retrieves the value of the environment variable named by the key,
// parses the value as a 64-bit floating-point number, and returns the result. If
// the variable is not present or its value cannot be parsed, fallback is
// returned.
func Float64(key string, fallback float64) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}

	return res
}
