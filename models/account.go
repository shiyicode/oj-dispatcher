package models

import (
	. "github.com/open-fightcoder/oj-dispatcher/common/store"
)

type Account struct {
	Id        int64
	Email     string //邮箱
	Password  string //密码
	Phone     string //手机号
	QqId      string //用于QQ第三方登录
	GithubId  string //Github第三方登录
	WeichatId string //weichat第三方登录
}

//增加
func (this Account) Add(account *Account) (int64, error) {
	_, err := OrmWeb.Insert(account)
	if err != nil {
		return 0, err
	}
	return account.Id, nil
}

//删除
func (this Account) Remove(id int64) error {
	account := &Account{}
	_, err := OrmWeb.Id(id).Delete(account)
	return err
}

//修改
func (this Account) Update(account *Account) error {
	_, err := OrmWeb.AllCols().ID(account.Id).Update(account)
	return err
}

//查询
func (this Account) GetById(id int64) (*Account, error) {
	account := new(Account)

	has, err := OrmWeb.Id(id).Get(account)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return account, nil
}

func (this Account) GetByEmail(email string) (*Account, error) {
	account := new(Account)
	has, err := OrmWeb.Where("email=?", email).Get(account)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return account, nil
}
