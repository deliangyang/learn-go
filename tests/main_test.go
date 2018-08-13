package tests

import "testing"

func TestTrue(t *testing.T)  {
	//assert.True(t, 5 == 5)
	var a int = 5
	if a != 5 {
		t.Error("a is not 5")
	}
}
