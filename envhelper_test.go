package envhelper

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnvAsInt11(t *testing.T) {
	expected := 123
	os.Setenv("TEST_ENV_VAR", fmt.Sprintf("%d", expected))

	actual := GetEnvAsInt("TEST_ENV_VAR", 1)
	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	testEnvVar := "TEST_ENV_VAR1"
	os.Setenv(testEnvVar, "5")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key      string
		fallback int
		want     int
	}{
		{
			name:     "TEST1",
			key:      "TEST_ENV_VAR1",
			fallback: 2,
			want:     5,
		},
		{
			name:     "Test for env var not set",
			key:      "TEST_ENV_VAR_NOT_SET",
			fallback: 2,
			want:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetEnvAsInt(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("GetEnvAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	testEnvVar2 := "TEST_ENV_VAR2"
	os.Setenv(testEnvVar2, "value123")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key      string
		fallback string
		want     string
	}{
		{
			name:     "Test for existing env var",
			key:      testEnvVar2,
			fallback: "value456",
			want:     "value123",
		},
		{
			name:     "Test for non-existing env var",
			key:      "TEST_ENV_VAR_NOT_EXIST",
			fallback: "value456",
			want:     "value456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetEnv(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
