package judger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTimeoutClient(t *testing.T) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post("http://127.0.0.1:"+strconv.Itoa(9001)+"/apiv1/judge/default",
		"application/x-www-form-urlencoded",
		strings.NewReader("submit_id=1"))
	if err != nil {
		fmt.Println("timeout", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil err", err.Error())
		return
	}
	fmt.Println(string(body))
	if string(body) != "ok" {
		return
	}
}
