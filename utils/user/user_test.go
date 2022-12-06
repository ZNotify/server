package user

import (
	"notify-api/utils/config"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("test init", func(t *testing.T) {
		config.Load("test_config")
		Init()
		ret := Is("test")
		if ret != true {
			t.Errorf("Init() not working.")
		}
	})
}
