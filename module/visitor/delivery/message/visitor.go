package message

import (
	"bigproject/module/visitor"
	"bigproject/util/config"
	"bigproject/util/messaging"
	"github.com/nsqio/go-nsq"
)

type VisitorMessagingHandler struct {
	uv *visitor.VisitorModule
}

func (vm *VisitorMessagingHandler) ConsumeIncrement(msg *nsq.Message) error {
	vm.uv.IncrementCount()
}

func RegisterVisitorMQ(uv *visitor.VisitorModule){
	handler := &VisitorMessagingHandler{
		uv:uv,
	}
	conf := config.GetConfig()

	// create messaging options
	messagingOptions := messaging.Options{
		LookupAddress:  []string{"localhost:4161"},
		PublishAddress: "localhost:4150",
		Prefix:         "bigproject_",
	}
	consumerEngine := messaging.NewConsumer(messagingOptions)

	// create consumer's option
	consumerOption := messaging.ConsumerOptions{
		Topic:       "bigproject",
		Channel:     "dhanarjkusuma", //TODO : change this with your channel name
		Handler:     handler.ConsumeIncrement,
		MaxAttempts: uint16(conf.MQDefaultConsumerMaxAttempts),
		MaxInFlight: conf.MQDefaultConsumerMaxInFlight,
	}
	consumerEngine.RegisterConsumer(consumerOption)
}