package config

import "os"

var MQTT_GATEWAY_TOPIC = "/gateway/#"
var PLATFORM_WORKFLOW_DATA_TOPIC = "workflow"
var MQTT_QOS = 0

var MQTT_BROKER = "emqx:1883"
var MQTT_USER = "admin"
var MQTT_PASSWORD = "things2023"
var MQTT_CLIENT_ID = "things-service"

var CACHE_DEVICE_PRODUCT_MODEL_KEY_PREFIX = "dpm:"
var CACHE_DEVICE_MIRROR_KEY_PREFIX = "dm:"
var MONITOR_ENABLED = false

func init() {
	if val, ok := os.LookupEnv("MQTT_BROKER"); ok {
		MQTT_BROKER = val
	}
	if val, ok := os.LookupEnv("MQTT_USER"); ok {
		MQTT_USER = val
	}
	if val, ok := os.LookupEnv("MQTT_PASSWORD"); ok {
		MQTT_PASSWORD = val
	}
	if val, ok := os.LookupEnv("MQTT_CLIENT_ID"); ok {
		MQTT_CLIENT_ID = val
	}

}
