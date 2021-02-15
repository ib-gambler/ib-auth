package pkg

import "testing"

func TestValidateSSO(t *testing.T) {
	err := ValidateSSO()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAuthenticationStatus(t *testing.T) {
	err := GetAuthenticationStatus()
	if err != nil {
		t.Fatal(err)
	}
}