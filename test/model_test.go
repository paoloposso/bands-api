package test

import (
	"testing"

	"bands-api/user"
)

func Test_ShoulFailUserRegisterValidation(t *testing.T) {
	user := user.User {}
	err := user.ValidateRegister()
	if err == nil {
		t.Error("should have failed validation")
	}
	user.Name = "Paolo"
	user.Password = "123"
	err = user.ValidateRegister()
	if err == nil {
		t.Error("should have failed validation")
	}
	user.Email = "asdfg"
	err = user.ValidateRegister()
	if err == nil {
		t.Error("should have failed validation")
	}
}

func Test_ShouldPassUserRegisterValidate(t *testing.T) {
	user := user.User {
		Name: "Test",
		Email: "paolo@test.com",
		Password: "123456789",
	}
	err := user.ValidateRegister()
	if err != nil {
		t.Errorf("should passed validation %s", err)
	}
}