package judger

import (
	"fmt"
	"time"
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
	containerId int //容器ID
}

func NewJudger() *Judger {
	judger := new(Judger)

	judger.CreateDocker()
	return judger
}

func (j *Judger) CreateDocker() {
	fmt.Println("创建docker")
}

// TODO 用API超时来决定是否任务失败
func (j *Judger) Do(job *Job) {
	switch job.SubmitType {
	case SUBMITTYPE_DEFA:
		j.doDefa(job.SubmitId)
	case SUBMITTYPE_SPEC:
		j.doSpec(job.SubmitId)
	case SUBMITTYPE_TEST:
		j.doTest(job.SubmitId)
	default:
		panic("not ")
	}
}

func (j *Judger) doDefa(submitId int64) {
	time.Sleep(10 * time.Second)
	fmt.Println("dodefa")
}

func (j *Judger) doTest(submitId int64) {
	time.Sleep(10 * time.Second)
	fmt.Println("dotest")
}

func (j *Judger) doSpec(submitId int64) {
	time.Sleep(10 * time.Second)
	fmt.Println("dospec")
}

// 删除容器
func (j *Judger) DropDocker() {
	fmt.Println("删除容器")
}

func (j *Judger) checkHealth() bool {
	return true
}

// 销毁当前容器，创建新容器，并且写入任务的结果
func (j *Judger) reCreateDocker() {

}
