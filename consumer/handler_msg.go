package comsumer

import (
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Topic string
}

func (this *Handler) HandleMessage(m *nsq.Message) error {
	log.Infof("HandbleMessage: ", string(m.Body))

	//judgerData := new(judger.Judger)
	//if err := json.Unmarshal(m.Body, judgerData); err != nil {
	//	log.Errorf("unmarshal JudgerData from NsqMessage failed, err: %v, event:%s", err, m.Body)
	//	return nil
	//}
	//
	//log.Infof("consume Message from dispatch: %#v", judgerData)
	//
	//handlerCount <- 1
	//go this.doJudge(judgerData)

	// 返回err为nil表示消费成功
	return nil
}
