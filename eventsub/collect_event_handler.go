package eventsub

import (
	"context"
	"encoding/json"
	mycommon "github.com/dapr-platform/common"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"sync"
	"things-service/entity"
	"things-service/service"
)

var subscribeMap = sync.Map{}

func NewDeviceMsgEventHandler(server common.Service, eventPub string, eventTopic string) {
	var sub = &common.Subscription{
		PubsubName: eventPub,
		Topic:      eventTopic,
		Route:      "/DeviceMessageHandler",
	}

	err := server.AddTopicEventHandler(sub, deviceMsgHandler)

	if err != nil {
		panic(err)
	}
}

func deviceMsgHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("eventsub deviceMsgHandler - PubsubName: %s, Topic: %s, msg: %s, \n", e.PubsubName, e.Topic, string(e.RawData))
	var event entity.DeviceInfoMsg
	err = json.Unmarshal(e.RawData, &event)
	if err != nil {
		log.Println("eventsub - data ", string(e.RawData))
		log.Println("eventsub - unmarshal error: ", err)
	} else {
		err = service.ProcessDeviceMsg(ctx, event)
		if err != nil {
			mycommon.Logger.Error("ProcessDeviceMsg error ", err)
		}
	}

	return false, nil
}
