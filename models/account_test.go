package models

import (
	"testing"
)

func TestAccountAdd(t *testing.T) {
	InitAllInTest()

	account := &Account{Email: "fffff@qq.com", Password: "123", Phone: "1234", QqId: "1"}
	if _, err := AccountAdd(account); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}
func TestAccountUpdate(t *testing.T) {
	InitAllInTest()

	account := &Account{1, "qaqq@qq.com", "88", "10086", "2222", "33", "asd"}
	if err := AccountUpdate(account); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}
func TestAccountRemove(t *testing.T) {
	InitAllInTest()

	if err := AccountRemove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}
func TestAccountGetById(t *testing.T) {
	InitAllInTest()

	account := &Account{Email: "bbb@qq.com", Password: "123", Phone: "1234", QqId: "1"}
	AccountAdd(account)

	getAccount, err := AccountGetById(account.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getAccount != *account {
		t.Error("GetById() failed:", account, "!=", getAccount)
	}
}

func TestAccountGetByEmail(t *testing.T) {
	InitAllInTest()

	account := &Account{Email: "xxx@qq.com", Password: "123", Phone: "1234", QqId: "1"}
	AccountAdd(account)

	getAccount, err := AccountGetByEmail(account.Email)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getAccount != *account {
		t.Error("GetById() failed:", account, "!=", getAccount)
	}

}
