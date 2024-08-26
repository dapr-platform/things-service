package entity

import (
	"encoding/json"
	"testing"
)

func TestMirror(t *testing.T) {
	events := make([]DeviceMirrorEvent, 0)
	events = append(events, DeviceMirrorEvent{})
	services := make([]DeviceMirrorService, 0)
	services = append(services, DeviceMirrorService{})
	buf, _ := json.Marshal(events)
	t.Log(string(buf))
	buf, _ = json.Marshal(services)
	t.Log(string(buf))
}
