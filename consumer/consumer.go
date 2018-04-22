package consumer

import (
	"github.com/nsqio/go-nsq"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	NsqConsumer *nsq.Consumer
	NsqCfg      *nsq.Config
	Topic       string
	Channel     string
}

var consumers []*Consumer

func Start() {
	consumer, err := newConsumer()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("consumer start failure")

	}
	consumers = append(consumers, consumer)
}

func Stop() {
	for _, c := range consumers {
		c.NsqConsumer.Stop()
	}
}

func newConsumer() (*Consumer, error) {
	consumer := new(Consumer)

	consumer.NsqCfg = nsq.NewConfig()
	consumer.NsqCfg.MaxInFlight = g.Conf().Nsq.MaxInFlight
	consumer.Topic = g.Conf().Nsq.JudgeTopic
	consumer.Channel = g.Conf().Nsq.JudgeChannel

	var err error
	consumer.NsqConsumer, err = nsq.NewConsumer(consumer.Topic, consumer.Channel, consumer.NsqCfg)
	if err != nil {
		return nil, err
	}
	consumer.NsqConsumer.AddHandler(&Handler{Topic: consumer.Topic})

	err = consumer.NsqConsumer.ConnectToNSQLookupds(g.Conf().Nsq.Lookupds)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
