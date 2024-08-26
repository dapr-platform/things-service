package service

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strings"
	"things-service/entity"
)

func ProcessWorkflowEventDeviceAction(ctx context.Context, event entity.DeviceActionEvent) (err error) {
	sourceDevice := cast.ToStringMap(event.SourceDevice)
	sourceTags := cast.ToStringSlice(sourceDevice["tags"])
	queryTag := make(map[string]bool, 0)
	for _, matchTag := range event.MatchTags {
		for _, stag := range sourceTags {
			if strings.Index(stag, matchTag) == 0 {
				queryTag[stag] = true
			}
		}
	}
	var devices = make([]entity.DeviceInfo, 0)
	switch event.DataType {
	case "tag":
		tagStr := strings.Join(event.Tags, ",")
		devices, err = QueryDeviceByTagsAndProductId(ctx, tagStr, event.ProductId)
		if err != nil {
			err = errors.Wrap(err, "QueryDeviceByTagsAndProductId")
			common.Logger.Error(err)
			return
		}
	case "device":
		for _, deviceId := range event.DeviceIds {
			deviceWithTag, err1 := GetDeviceWithTagById(ctx, deviceId)
			if err1 != nil {
				err = errors.Wrap(err1, "GetDeviceWithTagById")
				common.Logger.Error(err)
				continue
			}
			if deviceWithTag == nil {
				err = errors.New("device " + deviceId + " not found")
				continue
			}
			devices = append(devices, *deviceWithTag)

		}
	case "product":
		var tagArr = make([]string, 0)
		for k, _ := range queryTag {
			tagArr = append(tagArr, k)
		}
		tagStr := strings.Join(tagArr, ",")
		devices, err = QueryDeviceByTagsAndProductId(ctx, tagStr, event.ProductId)
		if err != nil {
			err = errors.Wrap(err, "QueryDeviceByTagsAndProductId")
			common.Logger.Error(err)
			return
		}
	}
	for _, device := range devices {
		process := false
		if len(queryTag) == 0 {
			process = true
		} else {
			for _, tag := range device.Tags {
				if _, ok := queryTag[tag]; ok {
					process = true
					break
				}
			}
		}

		if process {
			req := entity.PropertySetReq{
				DeviceIdentifier:   device.Identifier,
				PropertyIdentifier: event.PropertyIdentifier,
				Value:              event.Value,
			}
			err = ProcessPropertySet(ctx, req)
			if err != nil {
				err = errors.Wrap(err, "ProcessPropertySet")
				common.Logger.Error(err)
				continue
			}
		}

	}
	return
}
