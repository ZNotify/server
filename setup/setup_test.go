package setup

import (
	"testing"

	"notify-api/setup/misc"
)

func Test_checkInternetConnection(t *testing.T) {
	t.Run("test check internet connection", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Expected no panic, got %v\n", err)
				t.Errorf("You should keep global network access.")
			}
		}()
		misc.RequireNetwork()
	})
}
