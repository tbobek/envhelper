package envhelper

import (
	"testing"
	"os"
)

func TestGetEnvAsInt(t *testing.T) {
	expected := 123
	os.Setenv("TEST_ENV_VAR", fmt.Sprintf("%d",expected))

	actual := GetEnvAsInt("TEST_ENV_VAR", 1)
	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}

}