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
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))
	if err != nil {
		fmt.Println("POST请求失败: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取Response失败: " + err.Error())
	}

	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("获取Body失败: " + err.Error())
	}
	assert.Equal(t, 0, respT.Code, "登录失败！")
}

func TestRegister(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=ssdfr.com&password=ssdfr"))
	if err != nil {
		fmt.Println("POST请求失败: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取Response失败: " + err.Error())
	}

	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("获取Body失败: " + err.Error())
	}
	assert.Equal(t, 0, respT.Code, "注册失败！")
}
