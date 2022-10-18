// Package env is still under development.
package env

import (
	"os"
	"strconv"
	"time"
)

// env represents an environment that maintains variables as key-value pairs of
// strings. Its API is not yet stable therefore it is not public.
type env struct {
	lookupFn func(key string) (string, bool)
}

// osEnv is the default env, managing environment variables using functions of
// the os package of the standard library. The top-level functions such as
// String, StringVar, and so on are wrappers for the methods of osEnv.
//
// osEnv is not public, yet.
var osEnv = &env{lookupFn: os.LookupEnv}

// String retrieves the value of the environment variable named by the key. If
// the variable is present in the environment, its value (which may be empty) is
// returned, otherwise fallback is returned.
func (e *env) String(key string, fallback string) string {
	value, ok := e.lookupFn(key)
	if !ok {
		return fallback
	}

	return value
}

// Bool retrieves the value of the environment variable named by the key, parses
// the value as a boolean, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func (e *env) Bool(key string, fallback bool) bool {
	value, ok := e.lookupFn(key)
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
func (e *env) Int(key string, fallback int) int {
	value, ok := e.lookupFn(key)
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
func (e *env) Int64(key string, fallback int64) int64 {
	value, ok := e.lookupFn(key)
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
func (e *env) Uint(key string, fallback uint) uint {
	value, ok := e.lookupFn(key)
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
func (e *env) Uint64(key string, fallback uint64) uint64 {
	value, ok := e.lookupFn(key)
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
func (e *env) Float32(key string, fallback float32) float32 {
	value, ok := e.lookupFn(key)
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
func (e *env) Float64(key string, fallback float64) float64 {
	value, ok := e.lookupFn(key)
	if !ok {
		return fallback
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}

	return res
}

// Duration retrieves the value of the environment variable named by the key,
// parses the value as time.Duration, and returns the result. If the variable is
// not present or its value cannot be parsed, fallback is returned.
func (e *env) Duration(key string, fallback time.Duration) time.Duration {
	value, ok := e.lookupFn(key)
	if !ok {
		return fallback
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return res
}

// String retrieves the value of the environment variable named by the key. If
// the variable is present in the environment, its value (which may be empty) is
// returned, otherwise fallback is returned.
func String(key string, fallback string) string {
	return osEnv.String(key, fallback)
}

// Bool retrieves the value of the environment variable named by the key, parses
// the value as a boolean, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Bool(key string, fallback bool) bool {
	return osEnv.Bool(key, fallback)
}

// Int retrieves the value of the environment variable named by the key, parses
// the value as an integer, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Int(key string, fallback int) int {
	return osEnv.Int(key, fallback)
}

// Int64 retrieves the value of the environment variable named by the key, parses
// the value as a 64-bit integer, and returns the result. If the variable is not
// present or its value cannot be parsed, fallback is returned.
func Int64(key string, fallback int64) int64 {
	return osEnv.Int64(key, fallback)
}

// Uint retrieves the value of the environment variable named by the key, parses
// the value as an unsigned integer, and returns the result. If the variable is
// not present or its value cannot be parsed, fallback is returned.
func Uint(key string, fallback uint) uint {
	return osEnv.Uint(key, fallback)
}

// Uint64 retrieves the value of the environment variable named by the key,
// parses the value as an unsigned 64-bit integer, and returns the result. If the
// variable is not present or its value cannot be parsed, fallback is returned.
func Uint64(key string, fallback uint64) uint64 {
	return osEnv.Uint64(key, fallback)
}

// Float32 retrieves the value of the environment variable named by the key,
// parses the value as a floating-point number, and returns the result. If the
// variable is not present or its value cannot be parsed, fallback is returned.
func Float32(key string, fallback float32) float32 {
	return osEnv.Float32(key, fallback)
}

// Float64 retrieves the value of the environment variable named by the key,
// parses the value as a 64-bit floating-point number, and returns the result. If
// the variable is not present or its value cannot be parsed, fallback is
// returned.
func Float64(key string, fallback float64) float64 {
	return osEnv.Float64(key, fallback)
}

// Duration retrieves the value of the environment variable named by the key,
// parses the value as time.Duration, and returns the result. If the variable is
// not present or its value cannot be parsed, fallback is returned.
func Duration(key string, fallback time.Duration) time.Duration {
	return osEnv.Duration(key, fallback)
}

// StringVar retrieves the value of the environment variable named by the key,
// and stores the result into the variable pointed by p.
func StringVar(p *string, key string, fallback string) {
	*p = osEnv.String(key, fallback)
}

// BoolVar retrieves the value of the environment variable named by the key,
// parses the value as a boolean, and stores the result into the variable pointed
// by p.
func BoolVar(p *bool, key string, fallback bool) {
	*p = osEnv.Bool(key, fallback)
}

// IntVar retrieves the value of the environment variable named by the key,
// parses the value as an integer, and stores the result into the variable
// pointed by p.
func IntVar(p *int, key string, fallback int) {
	*p = osEnv.Int(key, fallback)
}

// Int64Var retrieves the value of the environment variable named by the key,
// parses the value as a 64-bit integer, and stores the result into the variable
// pointed by p.
func Int64Var(p *int64, key string, fallback int64) {
	*p = osEnv.Int64(key, fallback)
}

// UintVar retrieves the value of the environment variable named by the key,
// parses the value as an unsigned integer, and stores the result into the
// variable pointed by p.
func UintVar(p *uint, key string, fallback uint) {
	*p = osEnv.Uint(key, fallback)
}

// Uint64Var retrieves the value of the environment variable named by the key,
// parses the value as an unsigned 64-bit integer, and stores the result into the
// variable pointed by p.
func Uint64Var(p *uint64, key string, fallback uint64) {
	*p = osEnv.Uint64(key, fallback)
}

// Float32Var retrieves the value of the environment variable named by the key,
// parses the value as a floating-point number, and stores the result into the
// variable pointed by p.
func Float32Var(p *float32, key string, fallback float32) {
	*p = osEnv.Float32(key, fallback)
}

// Float64Var retrieves the value of the environment variable named by the key,
// parses the value as a 64-bit floating-point number, and stores the result into
// the variable pointed by p.
func Float64Var(p *float64, key string, fallback float64) {
	*p = osEnv.Float64(key, fallback)
}

// DurationVar retrieves the value of the environment variable named by the key,
// parses the value as time.Duration, and stores the result into the variable
// pointed by p.
func DurationVar(p *time.Duration, key string, fallback time.Duration) {
	*p = osEnv.Duration(key, fallback)
}
