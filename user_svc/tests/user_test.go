package tests

import (
	"testing"
	"user_svc/pkg/model"
)

func TestValidateUserStruct(t *testing.T) {
	u := &model.User{
		Email: "msms",
	}
	err := u.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
