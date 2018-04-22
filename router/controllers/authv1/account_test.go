package authv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"

	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Post("http://127.0.0.1:8000/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))

	resp, err = client.Post("http://127.0.0.1:8000/authv1/quit",
		"application/x-www-form-urlencoded", strings.NewReader(""))
	if err != nil {
		fmt.Println("POST请求失败: " + err.Error())
	}
	defer resp.Body.Close()
	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取Response失败: " + err.Error())
		}
		var respT base.HttpResponse
		if err := json.Unmarshal(body, &respT); err != nil {

			fmt.Println("获取Body失败: " + err.Error())
		}
		assert.Equal(t, 0, respT.Code, "退出失败！")
	}
}
