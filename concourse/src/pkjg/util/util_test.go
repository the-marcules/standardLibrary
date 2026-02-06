package util

import (
	"reflect"
	"testing"
)

func TestWorker(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry"}
		expected := []string{"0. item apple", "1. item banana", "2. item cherry"}

		result := Worker(items)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Worker() = %v, want %v", result, expected)
		}

	})
}
