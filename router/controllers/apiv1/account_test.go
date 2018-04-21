package apiv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/account/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 0, respT.Code, "登录失败！")
}

func TestRegister(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/account/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 0, respT.Code, "注册失败！")
}
