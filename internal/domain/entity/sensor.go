package entity

import "time"

type Sensor struct {
	Id          string    `json:"id"`
	Tipo        string    `json:"tipo"`
	Temperatura int       `json:"temperatura"`
	Timestamp   time.Time `json:"timestamp"`
	Callback    func() int
}

func (s *Sensor) SetTimeNow() {
	s.Timestamp = time.Now()
	return
}

func (s *Sensor) GenerateReading() {
	s.Temperatura = s.Callback()
}
