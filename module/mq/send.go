package mq

import (
        "github.com/streadway/amqp"

        "fmt"
        "log"
)

type Send struct {
}

func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
                panic(fmt.Sprintf("%s: %s", msg, err))
        }
}

// new Send Object
func NewSend() *Send {
        return &Send{}
}

//MQ Send object
func (s *Send) Send() {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:6666/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

        q, err := ch.QueueDeclare(
                "hello", // name
                false,   // durable
                false,   // delete when unused
                false,   // exclusive
                false,   // no-wait
                nil,     // arguments
        )
        failOnError(err, "Failed to declare a queue")

        body := "hello test"
        err = ch.Publish(
                "",     // exchange
                q.Name, // routing key
                false,  // mandatory
                false,  // immediate
                amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte(body),
                })
        log.Printf(" [x] Sent %s", body)
        failOnError(err, "Failed to publish a message")
}
