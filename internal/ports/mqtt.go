package ports

type MQTTPort interface {
	Publish(topic string, payload any)
	Subscribe(topic string)
}
