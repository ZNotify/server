package utils

import (
	"testing"
)

func TestCheckInternetConnection(t *testing.T) {
	t.Run("test check internet connection", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Expected no panic, got %v", err)
				t.Errorf("You should keep global network access.")
			}
		}()
		CheckInternetConnection()
	})
}

func TestIsTestInstance(t *testing.T) {
	t.Run("test is test instance", func(t *testing.T) {
		if got := IsTestInstance(); got != true {
			t.Errorf("IsTestInstance() not working.")
		}
	})
}
