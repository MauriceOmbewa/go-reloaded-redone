package transformation

import (
	"reflect"
	"testing"
)

func TestUppercase(t *testing.T) {
	sample := []string{"Hello", "world", "(up,", "2)"}
	expected := []string{"HELLO", "WORLD"}

	result := Uppercase(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
