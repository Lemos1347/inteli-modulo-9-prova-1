package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-prova-1/internal/domain/entity"
	"github.com/Lemos1347/inteli-modulo-9-prova-1/internal/infra"
	"github.com/Lemos1347/inteli-modulo-9-prova-1/internal/ports"
)

func CreateSensor(id string, tipo string, callback func() int) *entity.Sensor {
	return &entity.Sensor{
		Id:       id,
		Tipo:     tipo,
		Callback: callback,
	}
}

type Emulator struct {
	sensor     *entity.Sensor
	mqttClient ports.MQTTPort
}

func (s *Emulator) Start() {
	for {
		time.Sleep(time.Second * 1)
		s.sensor.GenerateReading()
		s.sensor.SetTimeNow()

		temp := struct {
			Id          string    `json:"id"`
			Tipo        string    `json:"tipo"`
			Temperatura int       `json:"temperatura"`
			Timestamp   time.Time `json:"timestamp"`
		}{
			Id:          s.sensor.Id,
			Tipo:        s.sensor.Tipo,
			Temperatura: s.sensor.Temperatura,
			Timestamp:   s.sensor.Timestamp,
		}

		payload, err := json.Marshal(temp)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(payload)

		s.mqttClient.Publish("sensors/data", payload)
	}
}

func NewEmulator(sensor *entity.Sensor) *Emulator {
	mqttconnection := infra.NewMQTTConnection(sensor.Id)

	return &Emulator{
		sensor:     sensor,
		mqttClient: mqttconnection,
	}
}
