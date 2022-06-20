package utils

import (
	"testing"
)

func TestIsTestInstance(t *testing.T) {
	t.Run("test is test instance", func(t *testing.T) {
		if got := IsTestInstance(); got != true {
			t.Errorf("IsTestInstance() not working.")
		}
	})
}
