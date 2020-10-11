package pta1001

import "testing"

func TestPta(t *testing.T) {
	if 5 != Pta(3) {
		t.Log("failed")
	}
}
