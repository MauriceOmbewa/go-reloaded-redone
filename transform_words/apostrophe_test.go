package transformation

import (
	"reflect"
	"testing"
)

func TestApostrophe(t *testing.T) {
	sample := []string{"Recode", "your", "'", "world", "with", "'", "your", "talent"}
	expected := []string{"Recode", "your", "'world", "with'", "your", "talent"}

	result := Apostrophe(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
