package model

import "testing"

func TestGormInit(t *testing.T) {

	if err := GormInit(); err != nil {
		t.Fatal(err)
	}
}

func TestGormClose(t *testing.T) {

	if err := GormClose(); err != nil {
		t.Error(err)
	}

}
