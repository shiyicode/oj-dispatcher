package judger

import (
	"fmt"
)

const (
	SUBMITTYPE_TEST = "test"
	SUBMITTYPE_DEFA = "default"
	SUBMITTYPE_SPEC = "special"
)

type Job struct {
	SubmitType string
	SubmitId   int64
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
	//time.Sleep(6 * time.Second)
}

// TODO 用API超时来决定是否任务失败
func (j *Judger) Do(job *Job) {
	fmt.Println("Do Judger ", job.SubmitType)
	switch job.SubmitType {
	case SUBMITTYPE_DEFA:

	case SUBMITTYPE_SPEC:
	case SUBMITTYPE_TEST:
	default:
		panic("not ")
	}
}

func doReal(submitId int64) {

}

func doTest(submitId int64) {

}

func doSpec(submitId int64) {

}

// 删除容器
func (j *Judger) DropDocker() {

}

func (j *Judger) checkHealth() bool {
	return true
}

// 销毁当前容器，创建新容器，并且写入任务的结果
func (j *Judger) reCreateDocker() {

}
