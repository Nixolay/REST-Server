package model

import (
	"errors"
	"fmt"
)

type User struct {
	ID         int64  `gorm:primary key;not_nil`
	Login      string `gorm:not_nil`
	Pass       string `gorm:not_nil`
	WorkNumber int32
}

func (u *User) Get(login string, pass string) (err error) {
	defer func() {
		if er := recover(); er!=nil{
			err = errors.New(fmt.Sprintf("%-v",er))
		}
	}()

	println(login, pass)
	return  DBConn.Where("login = $1 AND pass = $2", login, pass).First(&u).Error
}

func (u *User) Save() error {
	return DBConn.Save(u).Error
}

