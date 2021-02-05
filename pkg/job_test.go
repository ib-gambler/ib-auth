package pkg

import "testing"

func TestValidateSSO(t *testing.T) {
	err := ValidateSSO()
	if err != nil {
		t.Fatal(err)
	}
}
