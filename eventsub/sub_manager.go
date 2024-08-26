package eventsub

import (
	mycommon "github.com/dapr-platform/common"
	"github.com/dapr/go-sdk/service/common"
	"things-service/config"
)

func Sub(s common.Service) {
	NewDeviceMsgEventHandler(s, mycommon.PUBSUB_NAME, mycommon.DEVICE_DATA_TOPIC)
	NewWebMsgEventHandler(s, mycommon.PUBSUB_NAME, mycommon.INTERNAL_MESSAGE_TOPIC_NAME)
	NewWorkflowEventHandler(s, mycommon.PUBSUB_NAME, config.PLATFORM_WORKFLOW_DATA_TOPIC)
}
