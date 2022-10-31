package user

import (
	"testing"

	"notify-api/utils/config"
)

func TestInit(t *testing.T) {
	t.Run("test init", func(t *testing.T) {
		config.SetTest()
		Init()
		ret := Is("test")
		if ret != true {
			t.Errorf("Init() not working.")
		}
	})
}
