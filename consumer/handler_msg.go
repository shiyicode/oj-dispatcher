package consumer

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
	//log.Infof("HandleMessage: ", string(m.Body))

	job := new(judger.Job)
	if err := json.Unmarshal(m.Body, job); err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"message": m.Body,
		}).Error("unmarshal job from NsqMessage failed")
		return nil
	}

	//log.Infof("consume Message from dispatch: %#v", job)

	dispatcher.AddJob(job)

	// 返回err为nil表示消费成功
	return nil
}
