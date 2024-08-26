package service

import (
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"

	"things-service/entity"
)

func ProcessWebMsg(ctx context.Context, msg common.InternalMessage) (err error) {

	switch msg.GetType() {
	case common.INTERNAL_MESSAGE_TYPE_WEB_CONNECT:
		mark, exist := msg[common.INTERNAL_MESSAGE_KEY_MARK].(string)
		if !exist {
			err = errors.New("mark not exist")
			return
		}

		switch mark {
		case "device_mirror":
			err = processDeviceMirrorMsg(ctx, msg) // 第一次接到消息，要立刻反馈给web
		default:
			common.Logger.Debugf("don't need process mark type %s", mark)

		}
	default:
		common.Logger.Debugf("don't need process message type %s", msg.GetType())
	}
	return
}

// 第一次接到消息，要立刻发一次
func processDeviceMirrorMsg(ctx context.Context, msg common.InternalMessage) (err error) {
	common.Logger.Info("processDeviceMirrorMsg start", msg)
	deviceIdentifier, exist := msg[common.INTERNAL_MESSAGE_KEY_CONNECT_ID].(string)
	if !exist {
		err = errors.New("connect id not exist")
		return
	}
	if deviceIdentifier == "" {
		err = errors.New("connect id not exist")
		return
	}
	deviceMirror, err := GetDeviceMirror(ctx, deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	if deviceMirror == nil {
		err = errors.New(deviceIdentifier + " device mirror not exist")
		return
	}
	device, err := GetDeviceWithTagByIdentifier(ctx, deviceIdentifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceWithTagByIdentifier")
		return
	}
	deviceProductJson, err := GetDeviceProductModelJsonString(ctx, deviceIdentifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceProductModel")
		return
	}
	err = SendDeviceMirrorToMessageChannel(ctx, deviceIdentifier, device, deviceMirror, deviceProductJson)
	return
}
func SendDeviceMirrorToMessageChannel(ctx context.Context, deviceIdentifier string, device *entity.DeviceInfo, deviceMirror *entity.DeviceMirror, productJson string) (err error) {
	buf, err := json.Marshal(deviceMirror)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	var msg common.CommonMessage
	if productJson != "" {
		msg = common.CommonMessage{
			common.COMMON_MESSAGE_KEY_MARK:       "device_mirror",
			common.COMMON_MESSAGE_KEY_CONNECT_ID: deviceIdentifier,
			common.COMMON_MESSAGE_KEY_MESSAGE:    string(buf),
			"message_ext":                        productJson,
		}
	} else {
		msg = common.CommonMessage{
			common.COMMON_MESSAGE_KEY_MARK:       "device_mirror",
			common.COMMON_MESSAGE_KEY_CONNECT_ID: deviceIdentifier,
			common.COMMON_MESSAGE_KEY_MESSAGE:    string(buf),
		}
	}

	err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, common.WEB_MESSAGE_TOPIC_NAME, msg)
	if err != nil {
		common.Logger.Error(err.Error())
	}
	//for workflow

	if device != nil {
		mapMsg := make(map[string]any)
		buf1, _ := json.Marshal(device)
		_ = json.Unmarshal(buf1, &mapMsg)
		delete(mapMsg, "json_data") //减少数据大小
		for k, v := range deviceMirror.State.Reported {
			mapMsg[k] = v
		}

		err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, "device", mapMsg)
	}

	return
}
