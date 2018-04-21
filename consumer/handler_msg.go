package comsumer

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
	"github.com/open-fightcoder/oj-dispatcher/dispatcher"
	"github.com/open-fightcoder/oj-dispatcher/judger"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Topic string
}

func (this *Handler) HandleMessage(m *nsq.Message) error {
	log.Infof("HandbleMessage: ", string(m.Body))

	job := new(judger.Job)
	if err := json.Unmarshal(m.Body, job); err != nil {
		log.Errorf("unmarshal JudgerData from NsqMessage failed, err: %v, event:%s", err, m.Body)
		return nil
	}

	log.Infof("consume Message from dispatch: %#v", job)

	dispatcher.AddJob(job)

	// 返回err为nil表示消费成功
	return nil
}
