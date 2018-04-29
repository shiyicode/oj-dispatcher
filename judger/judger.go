package judger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"strconv"

	"log"

	"github.com/open-fightcoder/oj-dispatcher/docker"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
)

const (
	SUBMITTYPE_TEST = "test"
	SUBMITTYPE_DEFA = "default"
	SUBMITTYPE_SPEC = "special"
)

type Job struct {
	SubmitType string `json:"submit_type"`
	SubmitId   int64  `json:"submit_id"`
}

type Judger struct {
	containerId string // 容器ID
	id          int    // judger编号
}

func NewJudger(id int) *Judger {
	judger := new(Judger)
	judger.id = id
	judger.createDocker()
	return judger
}

// TODO 用API超时来决定是否任务失败
func (j *Judger) Do(job *Job) {
	var err error

	switch job.SubmitType {
	case SUBMITTYPE_DEFA:
		err = j.doDefa(job.SubmitId)
	case SUBMITTYPE_SPEC:
		err = j.doSpec(job.SubmitId)
	case SUBMITTYPE_TEST:
		err = j.doTest(job.SubmitId)
	default:
		panic("not has this submit type %s" + job.SubmitType)
	}

	if err != nil || j.checkHealth() {
		j.DropDocker()
		j.createDocker()
	}
}

const (
	JudgerUrl = "http://127.0.0.1:9001"
)

func (j *Judger) doDefa(submitId int64) error {
	cli := j.getClient()
	resp, err := cli.Post("http://127.0.0.1:"+strconv.Itoa(8000+j.id)+"/apiv1/judge/default",
		"application/x-www-form-urlencoded",
		strings.NewReader("submit_id="+strconv.FormatInt(submitId, 10)))
	if err != nil {
		fmt.Println("A:", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("B:", err.Error())
		return err
	}
	fmt.Println(string(body))
	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("C:", err.Error())
		return err
	}
	fmt.Printf("%s\n", respT)

	return nil
}

func (j *Judger) doTest(submitId int64) error {
	time.Sleep(10 * time.Second)
	fmt.Println("dotest")
	return nil
}

func (j *Judger) doSpec(submitId int64) error {
	time.Sleep(10 * time.Second)
	fmt.Println("dospec")
	return nil
}

func (j *Judger) createDocker() {
	fmt.Println("创建docker")
	bindPort := strconv.Itoa(8000 + j.id)
	var err error
	j.containerId, err = docker.CreateContainer("shiyicode/judger", []string{}, bindPort)
	if err != nil {
		log.Panicf("create container %s failure: ", j.containerId, err.Error())
	}

	err = docker.StartContainer(j.containerId)
	if err != nil {
		log.Panicf("start container %s failure: ", j.containerId, err.Error())
	}
}

// 删除容器
func (j *Judger) DropDocker() {
	fmt.Println("删除容器")
	err := docker.KillContainer(j.containerId)
	if err != nil {
		log.Panicf("kill container %s failure: ", j.containerId, err.Error())
	}
}

// 是否健康
func (j *Judger) checkHealth() bool {
	cli := j.getClient()
	resp, err := cli.Get(JudgerUrl + "/apiv1/self/health")
	if err != nil {
		fmt.Println("check err ", err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil err", err.Error())
		return false
	}
	fmt.Println(string(body))
	if string(body) != "ok" {
		return false
	}

	return true
}

// 销毁当前容器，创建新容器，并且写入任务的结果
func (j *Judger) reCreateDocker() {

}

func (j *Judger) getClient() http.Client {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	return client
}
