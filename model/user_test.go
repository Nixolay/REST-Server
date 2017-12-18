package model

import (
	"testing"
	"fmt"
)

func TestUser_Get(t *testing.T) {
	//t.Skip()

	GormInit()
	defer GormClose()

	us := User{}

	if err := us.Get("test_user", "test_pass"); err != nil {
		t.Fatal(err)
	}
	fmt.Println("----",us)

}

func TestUser_Save(t *testing.T) {
	//t.Skip()

	GormInit()
	defer GormClose()

	us := User{1, "test_user", "test_pass", 1}

	if err := us.Save(); err != nil {
		t.Error(err)
	}

}
