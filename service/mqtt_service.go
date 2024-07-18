package service

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	Client      MQTT.Client
	MessageChan chan string
	LastMessage string
}

func NewMQTTService(client MQTT.Client) *MQTTService {
	return &MQTTService{
		Client:      client,
		MessageChan: make(chan string),
	}
}

func (s *MQTTService) Connect() {
	if token := s.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (s *MQTTService) Subscribe(topic string, qos byte, callback MQTT.MessageHandler) {
	if token := s.Client.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (s *MQTTService) Publish(topic string, qos byte, retained bool, payload string) {
	token := s.Client.Publish(topic, qos, retained, payload)
	token.Wait()
	if token.Error() != nil {
		fmt.Println("Error publishing message:", token.Error())
	}
}

// OnMessageReceived handles incoming MQTT messages
func (s *MQTTService) OnMessageReceived(client MQTT.Client, msg MQTT.Message) {
	s.LastMessage = string(msg.Payload())
	s.MessageChan <- string(msg.Payload())
	// Optionally process or handle the received message further
	// fmt.Printf("Received message: %s from topic: %s\n", s.LastMessage, msg.Topic())
}
