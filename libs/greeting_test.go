package libs

import "testing"

func TestLogin(t *testing.T) {
	Greeting()
	if Login("admin", "234") != false {
		t.Error("error test fail")
	}

	if Login("admin", "234") != false {
		t.Error("error test pass")
	}
}
