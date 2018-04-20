package managers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/open-fightcoder/oj-dispatcher/common/components"
	"github.com/open-fightcoder/oj-dispatcher/models"
)

func AccountLogin(email, password string) (bool, string, string) {
	account := getAccountByEmail(email)
	if account == nil {
		return false, "", "Email is not exist"
	}
	if account.Password != md5Encode(password) {
		return false, "", "Password is wrong"
	} else {
		//TODO
		userId := 1
		if token, err := components.CreateToken(userId); err != nil {
			panic(err.Error())
		} else {
			return true, token, ""
		}
	}
}

func AccountRegister(email, password string) (bool, int64, string) {
	account := getAccountByEmail(email)
	if account != nil {
		return false, 0, "Email is exist"
	}
	account = &models.Account{Email: email, Password: md5Encode(password)}
	if insertId, err := (models.Account{}).Add(account); err != nil {
		panic(err.Error())
	} else {
		return true, insertId, ""
	}
}

func getAccountByEmail(email string) *models.Account {
	account, err := models.Account{}.GetByEmail(email)
	if err != nil {
		panic(err.Error())
	}
	return account
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
