package transformation

import (
	"reflect"
	"testing"
)

func TestHexadecimal(t *testing.T) {
	sample := []string{"1E", "(hex)", "your", "talent"}
	expected := []string{"30", "your", "talent"}

	result := Hexadecimal(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
