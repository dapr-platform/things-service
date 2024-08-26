package eventsub

import (
	"context"
	"encoding/json"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"things-service/entity"
	"things-service/service"
)

func NewWorkflowEventHandler(server common.Service, eventPub string, eventTopic string) {
	var sub = &common.Subscription{
		PubsubName: eventPub,
		Topic:      eventTopic,
		Route:      "/WorkflowEventHandler",
	}

	err := server.AddTopicEventHandler(sub, workflowEventHandler)

	if err != nil {
		panic(err)
	}
}

func workflowEventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("workflowEventHandler - PubsubName: %s, Topic: %s, msg: %s, \n", e.PubsubName, e.Topic, string(e.RawData))
	var event entity.WorkflowEvent
	err = json.Unmarshal(e.RawData, &event)
	if err != nil {
		log.Println("eventsub - data ", string(e.RawData))
		log.Println("eventsub - unmarshal error: ", err)
	} else {
		switch event.Type {
		case "DeviceAction":
			var actionEvent entity.DeviceActionEvent
			err = json.Unmarshal(e.RawData, &actionEvent)
			if err != nil {
				log.Println("eventsub "+event.Type+" - data unmarshal error: ", err)
			} else {
				err = service.ProcessWorkflowEventDeviceAction(ctx, actionEvent)
			}
			if err != nil {
				log.Println("eventsub "+event.Type+" - process error: ", err)
			}

		}
	}

	return false, nil
}
