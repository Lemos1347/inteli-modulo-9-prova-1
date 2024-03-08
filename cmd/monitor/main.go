package main

import (
	"encoding/json"
	"fmt"

	"github.com/Lemos1347/inteli-modulo-9-prova-1/internal/domain/entity"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(_ MQTT.Client, msg MQTT.Message) {
	data := entity.Sensor{}
	alert := ""

	json.Unmarshal(msg.Payload(), &data)

	if data.Tipo == "freezer" {
		if data.Temperatura > -15 {
			alert = "[ALERTA: Temperatura ALTA]"
		}
		if data.Temperatura < -25 {
			alert = "[ALERTA: Temperatura BAIXA]"
		}
	}

	if data.Tipo == "geladeira" {
		if data.Temperatura > 10 {
			alert = "[ALERTA: Temperatura ALTA]"
		}
		if data.Temperatura < 2 {
			alert = "[ALERTA: Temperatura BAIXA]"
		}
	}

	loja := string(data.Id[3])
	id := string(data.Id[6])

	fmt.Printf("Lj %s: %s %s | %d %s\n", loja, data.Tipo, id, data.Temperatura, alert)
}

func runSub(topic string, callback MQTT.MessageHandler) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(callback)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {}
}

func main() {
	runSub("sensors/data", messagePubHandler)
}
