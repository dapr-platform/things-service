package eventsub

import (
	"context"
	"encoding/json"
	mycommon "github.com/dapr-platform/common"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"things-service/service"
)

func NewWebMsgEventHandler(server common.Service, eventPub string, eventTopic string) {
	var sub = &common.Subscription{
		PubsubName: eventPub,
		Topic:      eventTopic,
		Route:      "/WebMessageHandler",
	}

	err := server.AddTopicEventHandler(sub, webMsgHandler)

	if err != nil {
		panic(err)
	}
}

func webMsgHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("eventsub webMsgHandler - PubsubName: %s, Topic: %s, data: %s, \n", e.PubsubName, e.Topic, string(e.RawData))
	var event mycommon.InternalMessage
	err = json.Unmarshal(e.RawData, &event)
	if err != nil {
		log.Println("eventsub - data ", string(e.RawData))
		log.Println("eventsub - unmarshal error: ", err)
	} else {
		err = service.ProcessWebMsg(ctx, event)
		if err != nil {
			log.Println("eventsub - process error: ", err)
		}
	}

	return false, nil
}
