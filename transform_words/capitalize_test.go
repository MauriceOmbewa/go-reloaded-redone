package transformation

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	sample := []string{"hello", "world", "(cap,", "2)"}
	expected := []string{"Hello", "World"}

	result := Capitalize(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
