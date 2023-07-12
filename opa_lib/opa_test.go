package opalib

import "testing"

func TestCheck(t *testing.T) {
	err := Check()
	if err != nil {
		t.Fatal(err)
	}
}
