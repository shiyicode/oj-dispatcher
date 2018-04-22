package managers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/open-fightcoder/oj-dispatcher/common/components"
	"github.com/open-fightcoder/oj-dispatcher/models"
	"github.com/pkg/errors"
)

func AccountLogin(email string, password string) (string, error) {
	account, err := models.AccountGetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("get account failure : %s ", err.Error())
	}
	if account == nil {
		return "", errors.New("Email is not exist")
	}
	if account.Password != md5Encode(password) {
		return "", errors.New("Password is wrong")
	}

	userId := 1
	token, err := components.CreateToken(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func AccountRegister(email, password string) (int64, error) {
	account, err := models.AccountGetByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("get account failure : %s ", err.Error())
	}
	if account != nil {
		return 0, errors.New("Email is exist")
	}
	account = &models.Account{Email: email, Password: md5Encode(password)}
	insertId, err := models.AccountAdd(account)
	if err != nil {
		return 0, fmt.Errorf("add account failure : %s ", err.Error())
	}
	return insertId, nil
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
