package transformation

import (
	"reflect"
	"testing"
)

func TestBinary(t *testing.T) {
	sample := []string{"10", "(bin)", "your", "talent"}
	expected := []string{"2", "your", "talent"}

	result := Binary(sample)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Apostrophe failed. Expected: %v. Got: %v", expected, result)
	}
}
