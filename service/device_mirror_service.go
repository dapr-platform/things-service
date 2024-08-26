package service

import (
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
	"sync"

	"things-service/entity"
	"things-service/model"
	"time"
)

var STATE_DEVICE_MIRROR_KEY = "device_mirror:"

// var persistLock = sync.Mutex{} //每个设备一个锁？
var deviceMirrorMapLock = sync.Mutex{}
var DeviceMirrorLockMap = make(map[string]*sync.Mutex)

func init() {
	common.RegisterDeleteBeforeHook("Device_mirror", DeleteDeviceMirrorHook)
}
func DeleteDeviceMirrorHook(r *http.Request, in any) (out any, err error) {
	id := cast.ToString(in)
	err = DeleteDeviceMirror(r.Context(), id)
	return
}

func GetDeviceMirrorLock(deviceIdentifier string) *sync.Mutex {
	deviceMirrorMapLock.Lock()
	defer deviceMirrorMapLock.Unlock()
	lock, ok := DeviceMirrorLockMap[deviceIdentifier]
	if ok {
		return lock
	}
	lock = &sync.Mutex{}
	DeviceMirrorLockMap[deviceIdentifier] = lock
	return lock
}

func NewDeviceMirror() *entity.DeviceMirror {
	t := time.Now()
	return &entity.DeviceMirror{
		State:                 entity.DeviceMirrorState{Reported: make(map[string]interface{}, 0), Desired: make(map[string]interface{}, 0)},
		Metadata:              entity.DeviceMirrorMetadata{Reported: make(map[string]interface{}, 0), Desired: make(map[string]interface{}, 0)},
		Timestamp:             t.UnixMilli(),
		CurrentKpiDatas:       make([]entity.CurrentKpiData, 0),
		Recent_InvokeServices: make([]entity.DeviceMirrorService, 0),
		RecentEvents:          make([]entity.DeviceMirrorEvent, 0),
		Alerts:                make([]map[string]any, 0),
		TimestampStr:          t.Format("2006-01-02 15:04:05"),
	}

}

func DeviceMirroerMergeKpiData(ctx context.Context, deviceIdentifier string, kpiDatas []entity.CurrentKpiData) (err error) {

	deviceMirror, err := GetDeviceMirror(ctx, deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "DeviceMirroerMergeKpiData")
		return
	}
	if deviceMirror == nil {
		common.Logger.Error(deviceIdentifier + " deviceMirror is nil")
		return
	}
	deviceMirror.CurrentKpiDatas = make([]entity.CurrentKpiData, 0)
	for _, kpiData := range kpiDatas {
		deviceMirror.CurrentKpiDatas = append(deviceMirror.CurrentKpiDatas, kpiData)
	}
	err = persistDeviceMirror(ctx, deviceIdentifier, deviceMirror)
	return
}
func MergeDeviceMirrorDesired(ctx context.Context, deviceIdentifier string, propertyIdentitfier string, value any) (err error) {

	deviceMirror, err := GetDeviceMirror(ctx, deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "MergeDeviceMirrorDesired")
		return
	}
	if deviceMirror == nil {
		common.Logger.Warning(deviceIdentifier + " deviceMirror is nil")
		deviceMirror = NewDeviceMirror()
	}
	deviceMirror.State.Desired[propertyIdentitfier] = value
	err = persistDeviceMirror(ctx, deviceIdentifier, deviceMirror)
	return
}

func PersistDeviceMirror(ctx context.Context, deviceIdentifier string, deviceMirror *entity.DeviceMirror) (err error) {
	err = persistDeviceMirror(ctx, deviceIdentifier, deviceMirror)
	return
}
func DeleteDeviceMirror(ctx context.Context, deviceIdentifier string) (err error) {
	err = common.DeleteKeyInStateStore(ctx, common.GetDaprClient(), common.GLOBAL_STATESTOR_NAME, STATE_DEVICE_MIRROR_KEY+deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
	}
	err = common.DbDelete(ctx, common.GetDaprClient(), model.Device_mirrorTableInfo.Name, model.Device_mirror_FIELD_NAME_id, deviceIdentifier)
	return
}
func persistDeviceMirror(ctx context.Context, deviceIdentifier string, deviceMirror *entity.DeviceMirror) (err error) {
	buf, err := json.Marshal(deviceMirror)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "PersistDeviceMirror")
		return
	}
	err = common.SaveInStateStore(ctx, common.GetDaprClient(), common.GLOBAL_STATESTOR_NAME, STATE_DEVICE_MIRROR_KEY+deviceIdentifier, buf, true, time.Hour*24)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "PersistDeviceMirror")
		return
	}
	common.Logger.Debugf("Save %s in state store %s ", deviceIdentifier, common.GLOBAL_STATESTOR_NAME)
	dbMirror := model.Device_mirror{
		ID:          deviceIdentifier,
		UpdatedTime: common.LocalTime(time.UnixMilli(deviceMirror.Timestamp)),
		JSONData:    string(buf),
	}
	err = common.DbUpsert[model.Device_mirror](ctx, common.GetDaprClient(), dbMirror, model.Device_mirrorTableInfo.Name, model.Device_mirror_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "PersistDeviceMirror")
	}
	return
}

func GetDeviceMirror(ctx context.Context, deviceIdentifier string) (deviceMirror *entity.DeviceMirror, err error) {
	buf, err := common.GetInStateStore(ctx, common.GetDaprClient(), common.GLOBAL_STATESTOR_NAME, STATE_DEVICE_MIRROR_KEY+deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "GetDeviceMirror")
	} else {
		if len(buf) == 0 {
			common.Logger.Error(deviceIdentifier + " GetDeviceMirror from StateStore: empty")
		} else {
			deviceMirror = new(entity.DeviceMirror)
			err = json.Unmarshal(buf, deviceMirror)
			if err != nil {
				common.Logger.Error(err.Error())
				err = errors.Wrap(err, "GetDeviceMirror")
				return
			}
			return
		}

	}
	dbMirror, err := common.DbGetOne[model.Device_mirror](ctx, common.GetDaprClient(), model.Device_mirrorTableInfo.Name, model.Device_mirror_FIELD_NAME_id+"="+deviceIdentifier)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "GetDeviceMirror")
		return
	}
	if dbMirror == nil {
		return
	}

	deviceMirror = new(entity.DeviceMirror)
	err = json.Unmarshal([]byte(dbMirror.JSONData), deviceMirror)
	if err != nil {
		common.Logger.Error(err.Error())
		err = errors.Wrap(err, "GetDeviceMirror")
		return
	}
	err = PersistDeviceMirror(ctx, deviceIdentifier, deviceMirror)
	return
}
