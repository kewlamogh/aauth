package aauth

import "testing"

func TestGenToken(t *testing.T) {
	want := "7063cdd951fa1311939d12698440ee25816e1264e107e2e865200e1f28170d33"
	got := GenToken("user1", "pass1")

	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}