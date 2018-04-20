package apiv1

import (
	"encoding/json"
	"fmt"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/baseController"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/account/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=abcd.com&password=asdf"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var respT baseController.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 0, respT.Code, "登录失败！")
}

func TestRegister(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/account/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=abcd.com&password=asdf"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var respT baseController.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 0, respT.Code, "注册失败！")
}
