package transformation

import (
	"reflect"
	"testing"
)

func TestVowels(t *testing.T) {
	sample := []string{"This", "is", "a", "umbrella"}
	expected := []string{"This", "is", "an", "umbrella"}

	result := Vowels(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
