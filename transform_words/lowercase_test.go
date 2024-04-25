package transformation

import (
	"reflect"
	"testing"
)

func TestLowercase(t *testing.T) {
	sample := []string{"HELLO", "WORLD", "(low)", "talent"}
	expected := []string{"HELLO", "world", "talent"}

	result := Lowercase(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
