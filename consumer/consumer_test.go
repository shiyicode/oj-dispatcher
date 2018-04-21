package comsumer

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"github.com/nsqio/go-nsq"
)

func TestSendMessCpp(t *testing.T) {
	Nsq{}.send("realJudge", &SendMess{"submit", 1, "problem", 1})
}

func TestSendMessPy(t *testing.T) {
	Nsq{}.send("realJudge", &SendMess{"submit", 2, "problem", 1})
}

func TestSendMessC(t *testing.T) {
	Nsq{}.send("realJudge", &SendMess{"submit", 3, "problem", 1})
}

func TestSendMessGo(t *testing.T) {
	Nsq{}.send("realJudge", &SendMess{"submit", 4, "problem", 1})
}

func TestSendMessJava(t *testing.T) {
	Nsq{}.send("realJudge", &SendMess{"submit", 5, "problem", 1})
}

type Nsq struct{}

type SendMess struct {
	SubmitType  string `json:"submit_type"`  //提交类型
	SubmitId    int64  `json:"submit_id"`    //提交id
	ProblemType string `json:"problem_type"` //题库类型
	ProblemId   int64  `json:"problem_id"`   //题目Id
}

func (this Nsq) send(topic string, sendMess *SendMess) {
	if topic != "realJudge" && topic != "virtualJudge" {
		err := errors.New("topic is false!")
		panic(err.Error())
	}
	adds := [1]string{"xupt2.fightcoder.com:9002"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(adds))
	msg, err := json.Marshal(sendMess)
	if err != nil {
		fmt.Println(err)
	}
	postNsq(adds[num], topic, msg)
}
func postNsq(address, topic string, msg []byte) {
	config := nsq.NewConfig()
	if w, err := nsq.NewProducer(address, config); err != nil {
		panic(err)
	} else {
		w.Publish(topic, msg)
	}
}
