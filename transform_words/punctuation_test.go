package transformation

import (
	"reflect"
	"testing"
)

func TestPunctuation(t *testing.T) {
	sample := []string{"Hello", "world", "!", "talent"}
	expected := []string{"Hello", "world!", "talent"}

	result := Punctuation(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
