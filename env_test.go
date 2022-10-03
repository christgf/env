package env_test

import (
	"testing"
	"time"

	"github.com/christgf/env"
)

func TestString(t *testing.T) {
	var key string

	key = "TEST_STRING_VAR"
	if got, want := env.String(key, "bar"), "bar"; got != want {
		t.Errorf("String(%q): got %q, want %q", key, got, want)
	}

	key = "TEST_STRING_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.String(key, "bar"), "foo"; got != want {
		t.Errorf("String(%q): got %q, want %q", key, got, want)
	}
}

func TestBool(t *testing.T) {
	var key string

	key = "TEST_BOOL_VAR"
	if got := env.Bool(key, true); got != true {
		t.Errorf("Bool(%q): got %v", key, got)
	}

	key = "TEST_BOOL_VAR_TRUE"
	t.Setenv(key, "true")
	if got := env.Bool(key, false); got != true {
		t.Errorf("Bool(%q): got %v", key, got)
	}

	key = "TEST_BOOL_VAR_ONE"
	t.Setenv(key, "1")
	if got := env.Bool(key, false); got != true {
		t.Errorf("Bool(%q): got %v", key, got)
	}

	key = "TEST_BOOL_VAR_FALSE"
	t.Setenv(key, "false")
	if got := env.Bool(key, true); got != false {
		t.Errorf("Bool(%q): got %v", key, got)
	}

	key = "TEST_BOOL_VAR_ZERO"
	t.Setenv(key, "0")
	if got := env.Bool(key, true); got != false {
		t.Errorf("Bool(%q): got %v", key, got)
	}

	key = "TEST_BOOL_VAR_FOO"
	t.Setenv(key, "foo")
	if got := env.Bool(key, true); got != true {
		t.Errorf("Bool(%q): got %v", key, got)
	}
}

func TestInt(t *testing.T) {
	var key string

	key = "TEST_INT_VAR"
	if got, want := env.Int(key, 7), 7; got != want {
		t.Errorf("Int(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_INT_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Int(key, 7), 9; got != want {
		t.Errorf("Int(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_INT_VAR_NEGATIVE_NINE"
	t.Setenv(key, "-9")
	if got, want := env.Int(key, 7), -9; got != want {
		t.Errorf("Int(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_INT_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Int(key, 5), 5; got != want {
		t.Errorf("Int(%q): got %d, want %d", key, got, want)
	}
}

func TestInt64(t *testing.T) {
	var key string

	key = "TEST_INT64_VAR"
	if got, want := env.Int64(key, 7), 7; got != int64(want) {
		t.Errorf("Int64(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_INT64_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Int64(key, 7), 9; got != int64(want) {
		t.Errorf("Int64(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_INT64_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Int64(key, 5), 5; got != int64(want) {
		t.Errorf("Int64(%q): got %d, want %d", key, got, want)
	}
}

func TestUint(t *testing.T) {
	var key string

	key = "TEST_UINT_VAR"
	if got, want := env.Uint(key, 7), uint(7); got != want {
		t.Errorf("Uint(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Uint(key, 7), uint(9); got != want {
		t.Errorf("Uint(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT_VAR_NEGATIVE_NINE"
	t.Setenv(key, "-9")
	if got, want := env.Uint(key, 7), uint(7); got != want {
		t.Errorf("Uint(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Uint(key, 5), uint(5); got != want {
		t.Errorf("Uint(%q): got %d, want %d", key, got, want)
	}
}

func TestUint64(t *testing.T) {
	var key string

	key = "TEST_UINT64_VAR"
	if got, want := env.Uint64(key, 7), 7; got != uint64(want) {
		t.Errorf("Uint64(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT64_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Uint64(key, 7), 9; got != uint64(want) {
		t.Errorf("Uint64(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT64_VAR_NEGATIVE_NINE"
	t.Setenv(key, "-9")
	if got, want := env.Uint64(key, 7), 7; got != uint64(want) {
		t.Errorf("Uint64(%q): got %d, want %d", key, got, want)
	}

	key = "TEST_UINT64_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Uint64(key, 5), 5; got != uint64(want) {
		t.Errorf("Uint64(%q): got %d, want %d", key, got, want)
	}
}

func TestFloat32(t *testing.T) {
	var key string

	key = "TEST_FLOAT32_VAR"
	if got, want := env.Float32(key, 4.25), float32(4.25); got != want {
		t.Errorf("Float32(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT32_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Float32(key, 7), float32(9.0); got != want {
		t.Errorf("Float32(%q): got %.3f, want %.3f", key, got, want)
	}

	key = "TEST_FLOAT32_VAR_TWO_POINT_NINE"
	t.Setenv(key, "2.9")
	if got, want := env.Float32(key, 7), float32(2.9); got != want {
		t.Errorf("Float32(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT32_VAR_NEGATIVE_POINT_NINE"
	t.Setenv(key, "-0.9")
	if got, want := env.Float32(key, 7), float32(-0.9); got != want {
		t.Errorf("Float32(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT32_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Float32(key, 5.25), float32(5.25); got != want {
		t.Errorf("Float32(%q): got %.2f, want %.2f", key, got, want)
	}
}

func TestFloat64(t *testing.T) {
	var key string

	key = "TEST_FLOAT64_VAR"
	if got, want := env.Float64(key, 4.25), 4.25; got != want {
		t.Errorf("Float64(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT64_VAR_NINE"
	t.Setenv(key, "9")
	if got, want := env.Float64(key, 7), 9.0; got != want {
		t.Errorf("Float64(%q): got %.3f, want %.3f", key, got, want)
	}

	key = "TEST_FLOAT64_VAR_TWO_POINT_NINE"
	t.Setenv(key, "2.9")
	if got, want := env.Float64(key, 7), 2.9; got != want {
		t.Errorf("Float64(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT64_VAR_NEGATIVE_POINT_NINE"
	t.Setenv(key, "-0.9")
	if got, want := env.Float64(key, 7), -0.9; got != want {
		t.Errorf("Float64(%q): got %.2f, want %.2f", key, got, want)
	}

	key = "TEST_FLOAT64_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Float64(key, 5.25), 5.25; got != want {
		t.Errorf("Float64(%q): got %.2f, want %.2f", key, got, want)
	}
}

func TestDuration(t *testing.T) {
	var key string

	key = "TEST_DURATION_VAR"
	if got, want := env.Duration(key, 2*time.Hour), 2*time.Hour; got != want {
		t.Errorf("Duration(%q): got %v, want %v", key, got, want)
	}

	key = "TEST_DURATION_VAR_TWO_HOURS"
	t.Setenv(key, "2h")
	if got, want := env.Duration(key, 3*time.Hour), 2*time.Hour; got != want {
		t.Errorf("Duration(%q): got %v, want %v", key, got, want)
	}

	key = "TEST_DURATION_VAR_TWO_AND_A_HALF_HOURS"
	t.Setenv(key, "2h30m")
	if got, want := env.Duration(key, 2*time.Hour), 150*time.Minute; got != want {
		t.Errorf("Duration(%q): got %v, want %v", key, got, want)
	}

	key = "TEST_DURATION_VAR_NEGATIVE_UNITS"
	t.Setenv(key, "-20ms")
	if got, want := env.Duration(key, 20*time.Millisecond), -20*time.Millisecond; got != want {
		t.Errorf("Duration(%q): got %v, want %v", key, got, want)
	}

	key = "TEST_DURATION_VAR_FOO"
	t.Setenv(key, "foo")
	if got, want := env.Duration(key, 180*time.Second), 3*time.Minute; got != want {
		t.Errorf("Duration(%q): got %v, want %v", key, got, want)
	}
}
