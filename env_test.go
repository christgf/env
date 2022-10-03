package env_test

import (
	"testing"
	"time"

	"github.com/christgf/env"
)

func TestString(t *testing.T) {
	const envKey = "ENV_TEST_STRING"

	if got, want := env.String(envKey, "foo"), "foo"; got != want {
		t.Errorf("String(%q): got %v", envKey, got)
	}

	t.Setenv(envKey, "bar")
	if got, want := env.String(envKey, "foo"), "bar"; got != want {
		t.Errorf("String(%q): got %q, want %q", envKey, got, want)
	}
}

func TestBool(t *testing.T) {
	const envKey = "ENV_TEST_BOOL"

	if got := env.Bool(envKey, true); got != true {
		t.Errorf("Bool(%q): got %v", envKey, got)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  bool
		wantValue bool
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  true,
			wantValue: true,
		},
		{
			name:      "Value set to true",
			envValue:  "true",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "Value set to false",
			envValue:  "false",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "Value set to TRUE",
			envValue:  "TRUE",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "Value set to FALSE",
			envValue:  "FALSE",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "Value set to T",
			envValue:  "T",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "Value set to F",
			envValue:  "F",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "Value set to 1",
			envValue:  "1",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "Value set to 0",
			envValue:  "0",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  true,
			wantValue: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Bool(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Bool(%q): got %v", envKey, got)
			}
		})
	}
}

func TestInt(t *testing.T) {
	const envKey = "ENV_TEST_INT"

	if got, want := env.Int(envKey, 42), 42; got != want {
		t.Errorf("Int(%q): got %d, want: %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  int
		wantValue int
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "Value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "Value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: -42,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Int(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int(%q): got %d, want: %d", envKey, got, want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	const envKey = "ENV_TEST_INT64"

	if got, want := env.Int64(envKey, 42), int64(42); got != want {
		t.Errorf("Int64(%q): got %d, want: %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  int64
		wantValue int64
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "Value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "Value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: -42,
		},
		{
			name:      "Value is 64-bit positive",
			envValue:  "9223372036854775807",
			fallback:  42,
			wantValue: 9223372036854775807,
		},
		{
			name:      "Value is 64-bit negative",
			envValue:  "-9223372036854775808",
			fallback:  42,
			wantValue: -9223372036854775808,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Int64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int64(%q): got %d, want: %d", envKey, got, want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	const envKey = "ENV_TEST_UINT"

	if got, want := env.Uint(envKey, 42), uint(42); got != want {
		t.Errorf("Uint(%q): got %d, want: %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  uint
		wantValue uint
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "Value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "Value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Uint(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint(%q): got %d, want: %d", envKey, got, want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	const envKey = "ENV_TEST_UINT64"

	if got, want := env.Uint64(envKey, 42), uint64(42); got != want {
		t.Errorf("Uint64(%q): got %d, want: %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  uint64
		wantValue uint64
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "Value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "Value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "Value is 64-bit positive",
			envValue:  "9223372036854775807",
			fallback:  42,
			wantValue: 9223372036854775807,
		},
		{
			name:      "Value is 64-bit negative",
			envValue:  "-9223372036854775808",
			fallback:  9223372036854775807,
			wantValue: 9223372036854775807,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Uint64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint64(%q): got %d, want: %d", envKey, got, want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	const envKey = "ENV_TEST_FLOAT32"

	if got, want := env.Float32(envKey, 42), float32(42); got != want {
		t.Errorf("Float32(%q): got %.2f, want: %.2f", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  float32
		wantValue float32
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  4.2,
			wantValue: 4.2,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "Value is zero zero",
			envValue:  "0.0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "Value is fraction",
			envValue:  "0.2",
			fallback:  0.3,
			wantValue: 0.2,
		},
		{
			name:      "Value is negative fraction",
			envValue:  "-0.2",
			fallback:  0.3,
			wantValue: -0.2,
		},
		{
			name:      "Value is negative",
			envValue:  "-4.2",
			fallback:  4.2,
			wantValue: -4.2,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  4.2,
			wantValue: 4.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Float32(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float32(%q): got %.2f, want: %.2f", envKey, got, want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	const envKey = "ENV_TEST_FLOAT64"

	if got, want := env.Float64(envKey, 42), float64(42); got != want {
		t.Errorf("Float64(%q): got %.2f, want: %.2f", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  float64
		wantValue float64
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  4.2,
			wantValue: 4.2,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "Value is zero zero",
			envValue:  "0.0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "Value is fraction",
			envValue:  "0.2",
			fallback:  0.3,
			wantValue: 0.2,
		},
		{
			name:      "Value is negative fraction",
			envValue:  "-0.2",
			fallback:  0.3,
			wantValue: -0.2,
		},
		{
			name:      "Value is negative",
			envValue:  "-4.2",
			fallback:  4.2,
			wantValue: -4.2,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  4.2,
			wantValue: 4.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Float64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float64(%q): got %.2f, want: %.2f", envKey, got, want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	const envKey = "ENV_TEST_FLOAT64"

	if got, want := env.Duration(envKey, 2*time.Second), 2*time.Second; got != want {
		t.Errorf("Duration(%q): got %v, want: %v", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  time.Duration
		wantValue time.Duration
	}{
		{
			name:      "Value is empty",
			envValue:  "",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Second,
		},
		{
			name:      "Value is zero",
			envValue:  "0",
			fallback:  2 * time.Second,
			wantValue: 0,
		},
		{
			name:      "Value is nano",
			envValue:  "2ns",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Nanosecond,
		},
		{
			name:      "Value is micro",
			envValue:  "2μs",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Microsecond,
		},
		{
			name:      "Value is millis",
			envValue:  "2ms",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Millisecond,
		},
		{
			name:      "Value is composite",
			envValue:  "2h30m45s",
			fallback:  2 * time.Hour,
			wantValue: (2 * time.Hour) + (30 * time.Minute) + (45 * time.Second),
		},
		{
			name:      "Value is negative",
			envValue:  "-2h30m",
			fallback:  150 * time.Minute,
			wantValue: -150 * time.Minute,
		},
		{
			name:      "Value is foobar",
			envValue:  "foobar",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Duration(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Duration(%q): got %v, want: %v", envKey, got, want)
			}
		})
	}
}
