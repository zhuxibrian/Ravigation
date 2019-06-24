package utils

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

func Publish(topic string, qos byte, retained bool, payload string)  {
	client.Publish(topic, qos, retained, payload)
}

func Subcsribe(topic string, qos byte, msghandler func(mqtt.Client, mqtt.Message))  {
	if token := client.Subscribe(topic, qos, msghandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func Connect() {
	opts := mqtt.NewClientOptions().AddBroker(Config().GetString("MqttConfig.Url")).SetClientID(Config().GetString("MqttConfig.Name"))

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Disconnect(){
	client.Disconnect(200)
}



