package aauth

import "testing"

func TestLogin(t *testing.T) {
	Signup("bob", "j")
	err := Login("bob", "j")
	if err.Error() == "wrong password" || err.Error() == "invalid user" {
		t.Error("Login threw innacurate error")
	}
}