package helpers

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	slice := []string{"apple", "banana", "orange", "grape", "apple"}

	// Test existing value
	index := IndexOf(slice, "orange")
	if index != 2 {
		t.Errorf("Expected index 2 for 'orange', but got %d", index)
	}

	// Test non-existing value
	index = IndexOf(slice, "pear")
	if index != -1 {
		t.Errorf("Expected index -1 for 'pear', but got %d", index)
	}
}
