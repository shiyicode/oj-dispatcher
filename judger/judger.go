package judger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/open-fightcoder/oj-dispatcher/docker"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
	log "github.com/sirupsen/logrus"
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

	if err != nil {
		log.Info("重置容器，错误：", err.Error())
		j.DropDocker()
		j.createDocker()
	}
	//if !j.checkHealth() {
	//	log.Info("重置容器，错误：check失败")
	//	j.DropDocker()
	//	j.createDocker()
	//}
}

func (j *Judger) doDefa(submitId int64) error {
	log.Info("id ", j.id, " submitId ", submitId)
	cli := j.getClient()
	resp, err := cli.Post("http://127.0.0.1:"+strconv.Itoa(8000+j.id)+"/apiv1/judge/default",
		"application/x-www-form-urlencoded",
		strings.NewReader("submit_id="+strconv.FormatInt(submitId, 10)))
	if err != nil {
		log.Error("请求失败:", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("B:", err.Error())
		return err
	}
	log.Info(submitId, " ", string(body))
	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		return err
	}

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
	bindPort := strconv.Itoa(8000 + j.id)
	var err error
	//j.containerId, err = docker.CreateContainer("test", []string{"./oj-judger"}, bindPort)
	j.containerId, err = docker.CreateContainer("shiyicode/oj-judger", []string{}, bindPort)
	if err != nil {
		log.Errorf("create container %s failure: ", j.containerId, err.Error())
		return
	}

	log.Info(j.id, "创建容器", j.containerId)

	err = docker.StartContainer(j.containerId)
	if err != nil {
		log.Errorf("start container %s failure: ", j.containerId, err.Error())
	}

	log.Info(j.id, "启动容器", j.containerId)
}

// 删除容器
func (j *Judger) DropDocker() {
	log.Info(j.id, "删除容器", j.containerId)
	err := docker.KillContainer(j.containerId)
	if err != nil {
		log.Errorf("kill container %s failure: ", j.containerId, err.Error())
	}
}

// 是否健康
func (j *Judger) checkHealth() bool {
	cli := j.getClient()
	resp, err := cli.Get("http://127.0.0.1:" + strconv.Itoa(8000+j.id) + "/apiv1/self/health")
	if err != nil {
		log.Println("check err ", err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("check err", err.Error())
		return false
	}
	if string(body) != "ok" {
		return false
	}

	return true
}

func (j *Judger) getClient() http.Client {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return client
}
