// service/mqtt_service.go
package service

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	client mqtt.Client
}

func NewMQTTService(broker string, clientID string) *MQTTService {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &MQTTService{client: client}
}

func (s *MQTTService) Publish(topic string, payload string) error {
	token := s.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}
