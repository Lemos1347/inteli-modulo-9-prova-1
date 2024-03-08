package infra

import (
	"fmt"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConnection struct {
	clientId string
	client   MQTT.Client
}

func (s *MQTTConnection) connect(callback ...MQTT.MessageHandler) {
	wg := sync.WaitGroup{}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID(s.clientId)

	if len(callback) > 0 {
		opts.SetDefaultPublishHandler(callback[0])
	}

	opts.OnConnect = func(client MQTT.Client) {
		fmt.Printf("-> Client %s connected successfully!\n", s.clientId)
		defer wg.Done()
	}

	opts.OnConnectionLost = func(client MQTT.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %s\n", s.clientId, err.Error())
		wg.Add(1)
	}

	wg.Add(1)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	wg.Wait()

	s.client = client

	return
}

func (s *MQTTConnection) Subscribe(topic string) {
	if token := s.client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Unable to subscribe: %s\n", token.Error()))
	}

	return
}

func (s *MQTTConnection) Publish(topic string, payload any) {
	if token := s.client.Publish(topic, 1, false, payload); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Unable to publish message: %s\n", token.Error()))
	}
}

func NewMQTTConnection(clientId string, callback ...MQTT.MessageHandler) *MQTTConnection {
	temp := MQTTConnection{
		clientId: clientId,
	}
	temp.connect(callback...)

	return &temp
}
