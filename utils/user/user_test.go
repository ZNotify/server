package user

import "testing"

func TestInit(t *testing.T) {
	t.Run("test init", func(t *testing.T) {
		Init()
		ret := Is("test")
		if ret != true {
			t.Errorf("Init() not working.")
		}
	})
}
