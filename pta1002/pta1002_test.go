package pta1002

import "testing"

func TestPta(t *testing.T) {
	if "yi san wu" != Pta("1234567890987654321123456789") {
		t.Log("failed")
	}
}
