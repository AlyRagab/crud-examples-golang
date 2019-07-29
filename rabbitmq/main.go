package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var err error

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Establish the Connection
	conn, err := amqp.Dial("amqp://localhost:5672/")
	check(err)
	fmt.Println("Connected Successfully !!")
	defer conn.Close()

	// Create a Channel
	ch, err := conn.Channel()
	check(err)
	defer ch.Close()
	fmt.Println("Channel is Created Successfully !!")

	// Create a Queue
	q, err := ch.QueueDeclare("Hello", false, false, false, false, nil)
	check(err)
	fmt.Println("Queue is Created Successfully ")

	body := "Hello-World"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("Message: %s", body)

}
