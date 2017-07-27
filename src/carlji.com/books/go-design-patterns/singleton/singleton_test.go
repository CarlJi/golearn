package singleton

import "testing"

func TestSingleton_GetInstance(t *testing.T) {

	instance := GetInstance()

	if instance == nil {
		t.Errorf("instance was nil")
	}

	another := GetInstance()
	if instance != another {
		t.Error("not equals")
	}

	res := another.AddOne()
	if res != 1 {
		t.Errorf("failed")
	}
}
