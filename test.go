package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected to MQTT broker")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func main1() {
	// Define the MQTT broker options
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// Create and start an MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Channel to receive messages from MQTT
	messageChan := make(chan string)

	// Subscribe to a topic and send received messages to the channel
	go func() {
		if token := client.Subscribe("orders/processed", 1, func(client MQTT.Client, msg MQTT.Message) {
			messageChan <- string(msg.Payload())
		}); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return
		}
	}()

	// Goroutine to publish messages to the topic
	go func() {
		for {
			// Publish a message
			text := fmt.Sprintf("Hello MQTT %s", time.Now().String())
			token := client.Publish("orders/processed", 0, false, text)
			token.Wait()

			// Sleep for a while before publishing the next message
			time.Sleep(2 * time.Second)
		}
	}()

	// Main goroutine to handle incoming messages from the channel
	go func() {
		for {
			select {
			case msg := <-messageChan:
				fmt.Printf("Received from channel: %s\n", msg)
			}
		}
	}()

	// Keep the main function running
	select {}
}
