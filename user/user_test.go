package user

import "testing"

func TestInit(t *testing.T) {
	t.Run("test init", func(t *testing.T) {
		Controller.Init()
		ret := Controller.Is("test")
		if ret != true {
			t.Errorf("Init() not working.")
		}
	})
}
