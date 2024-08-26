package service

import (
	"fmt"
	"github.com/dapr-platform/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"things-service/config"
	"time"
)

var running bool
var _client mqtt.Client

func MqttListen(broker, username, password, clientid string, deviceMsgProc func(c mqtt.Client, m mqtt.Message)) (err error) {
	running = false
	if _client == nil {
		// 掉线重连
		var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
			fmt.Printf("Mqtt Connect lost: %v", err)
			i := 0
			for {
				time.Sleep(5 * time.Second)
				if !_client.IsConnectionOpen() {
					i++
					fmt.Println("MQTT掉线重连...", i)
				} else {
					subscribe(deviceMsgProc)
					break
				}
			}
		}
		opts := mqtt.NewClientOptions()
		fmt.Println(broker + " " + username + " " + password + " " + clientid)
		opts.SetUsername(username)
		opts.SetPassword(password)
		opts.SetClientID(clientid)
		opts.AddBroker(broker)
		opts.SetAutoReconnect(true)
		opts.SetOrderMatters(false)
		opts.OnConnectionLost = connectLostHandler
		opts.SetOnConnectHandler(func(c mqtt.Client) {
			if !running {
				fmt.Println("MQTT CONNECT SUCCESS -- ", broker)
			}
			running = true
		})

		_client = mqtt.NewClient(opts)
		reconnec_number := 0
		for { // 失败重连
			if token := _client.Connect(); token.Wait() && token.Error() != nil {
				reconnec_number++
				fmt.Println("MQTT连接失败...重试", reconnec_number)
			} else {
				break
			}
			time.Sleep(5 * time.Second)
		}
		subscribe(deviceMsgProc)

	}
	return
}

// mqtt订阅
func subscribe(deviceMsgProc func(c mqtt.Client, m mqtt.Message)) {
	// 订阅直连设备
	if token := _client.Subscribe(common.DEVICE_DATA_TOPIC, byte(config.MQTT_QOS), func(c mqtt.Client, m mqtt.Message) {
		deviceMsgProc(c, m)
	}); token.Wait() &&
		token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}
