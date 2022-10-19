package env_test

import (
	"testing"
	"time"

	"github.com/christgf/env"
)

func TestString(t *testing.T) {
	const envKey = "ENV_TEST_STRING"

	if got, want := env.String(envKey, "fallback"), "fallback"; got != want {
		t.Errorf("String(%q): got %q, want %q", envKey, got, want)
	}

	t.Setenv(envKey, "foo")
	if got, want := env.String(envKey, "fallback"), "foo"; got != want {
		t.Errorf("String(%q): got %q, want %q", envKey, got, want)
	}

	var p string
	env.StringVar(&p, envKey, "fallback")
	if got, want := p, "foo"; got != want {
		t.Errorf("StringVar(%q): got %q, want %q", envKey, got, want)
	}

	prefix, key := "ENV_", "TEST_STRING"
	env.SetPrefix(prefix)
	if got, want := env.String(key, "fallback"), "foo"; got != want {
		t.Errorf("String(Prefix=%q, Key=%q): got %q, want %q", prefix, key, got, want)
	}

	env.SetPrefix("")
	if got, want := env.String(key, "fallback"), "fallback"; got != want {
		t.Errorf("String(Prefix=%q, Key=%q): got %q, want %q", prefix, key, got, want)
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
			name:      "value is empty",
			envValue:  "",
			fallback:  true,
			wantValue: true,
		},
		{
			name:      "value set to true",
			envValue:  "true",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "value set to false",
			envValue:  "false",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "value set to TRUE",
			envValue:  "TRUE",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "value set to FALSE",
			envValue:  "FALSE",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "value set to T",
			envValue:  "T",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "value set to F",
			envValue:  "F",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "value set to 1",
			envValue:  "1",
			fallback:  false,
			wantValue: true,
		},
		{
			name:      "value set to 0",
			envValue:  "0",
			fallback:  true,
			wantValue: false,
		},
		{
			name:      "value is foobar",
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

			var p bool
			env.BoolVar(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("BoolVar(%q): got %v", envKey, got)
			}

			prefix, key := "ENV_", "TEST_BOOL"
			env.SetPrefix(prefix)
			if got, want := env.Bool(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Bool(Prefix=%q, Key=%q): got %v", prefix, key, got)
			}

			env.SetPrefix("")
			if got, want := env.Bool(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Bool(Prefix=%q, Key=%q): got %v", prefix, key, got)
			}
		})
	}
}

func TestInt(t *testing.T) {
	const envKey = "ENV_TEST_INT"

	if got, want := env.Int(envKey, 42), 42; got != want {
		t.Errorf("Int(%q): got %d, want %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  int
		wantValue int
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: -42,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Int(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int(%q): got %d, want %d", envKey, got, want)
			}

			var p int
			env.IntVar(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("IntVar(%q): got %d, want %d", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_INT"
			env.SetPrefix(prefix)
			if got, want := env.Int(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Int(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Int(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	const envKey = "ENV_TEST_INT64"

	if got, want := env.Int64(envKey, 42), int64(42); got != want {
		t.Errorf("Int64(%q): got %d, want %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  int64
		wantValue int64
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: -42,
		},
		{
			name:      "value is 64-bit positive",
			envValue:  "9223372036854775807",
			fallback:  42,
			wantValue: 9223372036854775807,
		},
		{
			name:      "value is 64-bit negative",
			envValue:  "-9223372036854775808",
			fallback:  42,
			wantValue: -9223372036854775808,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Int64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int64(%q): got %d, want %d", envKey, got, want)
			}

			var p int64
			env.Int64Var(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("Int64Var(%q): got %d, want %d", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_INT64"
			env.SetPrefix(prefix)
			if got, want := env.Int64(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Int64(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Int64(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Int64(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	const envKey = "ENV_TEST_UINT"

	if got, want := env.Uint(envKey, 42), uint(42); got != want {
		t.Errorf("Uint(%q): got %d, want %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  uint
		wantValue uint
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Uint(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint(%q): got %d, want %d", envKey, got, want)
			}

			var p uint
			env.UintVar(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("UintVar(%q): got %d, want %d", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_UINT"
			env.SetPrefix(prefix)
			if got, want := env.Uint(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Uint(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Uint(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	const envKey = "ENV_TEST_UINT64"

	if got, want := env.Uint64(envKey, 42), uint64(42); got != want {
		t.Errorf("Uint64(%q): got %d, want %d", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  uint64
		wantValue uint64
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  42,
			wantValue: 0,
		},
		{
			name:      "value is positive",
			envValue:  "42",
			fallback:  43,
			wantValue: 42,
		},
		{
			name:      "value is negative",
			envValue:  "-42",
			fallback:  42,
			wantValue: 42,
		},
		{
			name:      "value is 64-bit positive",
			envValue:  "9223372036854775807",
			fallback:  42,
			wantValue: 9223372036854775807,
		},
		{
			name:      "value is 64-bit negative",
			envValue:  "-9223372036854775808",
			fallback:  9223372036854775807,
			wantValue: 9223372036854775807,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  42,
			wantValue: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Uint64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint64(%q): got %d, want %d", envKey, got, want)
			}

			var p uint64
			env.Uint64Var(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("Uint64Var(%q): got %d, want %d", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_UINT64"
			env.SetPrefix(prefix)
			if got, want := env.Uint64(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Uint64(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Uint64(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Uint64(Prefix=%q, Key=%q): got %d, want %d", prefix, key, got, want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	const envKey = "ENV_TEST_FLOAT32"

	if got, want := env.Float32(envKey, 42), float32(42); got != want {
		t.Errorf("Float32(%q): got %.2f, want %.2f", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  float32
		wantValue float32
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  4.2,
			wantValue: 4.2,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "value is zero zero",
			envValue:  "0.0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "value is fraction",
			envValue:  "0.2",
			fallback:  0.3,
			wantValue: 0.2,
		},
		{
			name:      "value is negative fraction",
			envValue:  "-0.2",
			fallback:  0.3,
			wantValue: -0.2,
		},
		{
			name:      "value is negative",
			envValue:  "-4.2",
			fallback:  4.2,
			wantValue: -4.2,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  4.2,
			wantValue: 4.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Float32(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float32(%q): got %.2f, want %.2f", envKey, got, want)
			}

			var p float32
			env.Float32Var(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("Float32Var(%q): got %.2f, want %.2f", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_FLOAT32"
			env.SetPrefix(prefix)
			if got, want := env.Float32(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float32(Prefix=%q, Key=%q): got %.2f, want %.2f", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Float32(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Float32(Prefix=%q, Key=%q): got %.2f, want %.2f", prefix, key, got, want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	const envKey = "ENV_TEST_FLOAT64"

	if got, want := env.Float64(envKey, 42), float64(42); got != want {
		t.Errorf("Float64(%q): got %.2f, want %.2f", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  float64
		wantValue float64
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  4.2,
			wantValue: 4.2,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "value is zero zero",
			envValue:  "0.0",
			fallback:  4.2,
			wantValue: 0,
		},
		{
			name:      "value is fraction",
			envValue:  "0.2",
			fallback:  0.3,
			wantValue: 0.2,
		},
		{
			name:      "value is negative fraction",
			envValue:  "-0.2",
			fallback:  0.3,
			wantValue: -0.2,
		},
		{
			name:      "value is negative",
			envValue:  "-4.2",
			fallback:  4.2,
			wantValue: -4.2,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  4.2,
			wantValue: 4.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Float64(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float64(%q): got %.2f, want %.2f", envKey, got, want)
			}

			var p float64
			env.Float64Var(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("Float64Var(%q): got %.2f, want %.2f", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_FLOAT64"
			env.SetPrefix(prefix)
			if got, want := env.Float64(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Float64(Prefix=%q, Key=%q): got %.2f, want %.2f", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Float64(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Float64(Prefix=%q, Key=%q): got %.2f, want %.2f", prefix, key, got, want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	const envKey = "ENV_TEST_DURATION"

	if got, want := env.Duration(envKey, 2*time.Second), 2*time.Second; got != want {
		t.Errorf("Duration(%q): got %v, want %v", envKey, got, want)
	}

	tests := []struct {
		name      string
		envValue  string
		fallback  time.Duration
		wantValue time.Duration
	}{
		{
			name:      "value is empty",
			envValue:  "",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Second,
		},
		{
			name:      "value is zero",
			envValue:  "0",
			fallback:  2 * time.Second,
			wantValue: 0,
		},
		{
			name:      "value is nano",
			envValue:  "2ns",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Nanosecond,
		},
		{
			name:      "value is micro",
			envValue:  "2μs",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Microsecond,
		},
		{
			name:      "value is millis",
			envValue:  "2ms",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Millisecond,
		},
		{
			name:      "value is composite",
			envValue:  "2h30m45s",
			fallback:  2 * time.Hour,
			wantValue: (2 * time.Hour) + (30 * time.Minute) + (45 * time.Second),
		},
		{
			name:      "value is negative",
			envValue:  "-2h30m",
			fallback:  150 * time.Minute,
			wantValue: -150 * time.Minute,
		},
		{
			name:      "value is foobar",
			envValue:  "foobar",
			fallback:  2 * time.Second,
			wantValue: 2 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.envValue)
			if got, want := env.Duration(envKey, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Duration(%q): got %v, want %v", envKey, got, want)
			}

			var p time.Duration
			env.DurationVar(&p, envKey, tt.fallback)
			if got, want := p, tt.wantValue; got != want {
				t.Errorf("DurationVar(%q): got %v, want %v", envKey, got, want)
			}

			prefix, key := "ENV_", "TEST_DURATION"
			env.SetPrefix(prefix)
			if got, want := env.Duration(key, tt.fallback), tt.wantValue; got != want {
				t.Errorf("Duration(Prefix=%q, Key=%q): got %v, want %v", prefix, key, got, want)
			}

			env.SetPrefix("")
			if got, want := env.Duration(key, tt.fallback), tt.fallback; got != want {
				t.Errorf("Duration(Prefix=%q, Key=%q): got %v, want %v", prefix, key, got, want)
			}
		})
	}
}

func TestSetPrefix(t *testing.T) {
	if got := env.Prefix(); got != "" {
		t.Fatalf("Prefix(): got %q", got)
	}

	env.SetPrefix("foo")
	if got, want := env.Prefix(), "foo"; got != want {
		t.Fatalf("Prefix(): got %q, want %q", got, want)
	}

	env.SetPrefix("")
	if got := env.Prefix(); got != "" {
		t.Fatalf("Prefix(): got %q, want empty", got)
	}
}
