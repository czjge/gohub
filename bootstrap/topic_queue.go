package bootstrap

import (
	"log"

	"github.com/czjge/gohub/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupTopicQueue() {

	for i := 0; i < 1; i++ {

		go func() {

			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			logger.LogIf(err)
			defer conn.Close()

			ch, err := conn.Channel()
			logger.LogIf(err)
			defer ch.Close()

			err = ch.ExchangeDeclare(
				"logs_topic", // name
				"topic",      // type
				true,         // durable
				false,        // auto-deleted
				false,        // internal
				false,        // no-wait
				nil,          // arguments
			)
			logger.LogIf(err)

			q, err := ch.QueueDeclare(
				"",    // name
				false, // durable
				false, // delete when unused
				true,  // exclusive
				false, // no-wait
				nil,   // arguments
			)
			logger.LogIf(err)

			bindings := []string{"kern.*", "*.critical"}
			for _, s := range bindings {
				log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_topic", s)
				err = ch.QueueBind(
					q.Name,       // queue name
					s,            // routing key
					"logs_topic", // exchange
					false,
					nil,
				)
				logger.LogIf(err)
			}

			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				true,   // auto-ack
				false,  // exclusive
				false,  // nolocal
				false,  // no wait
				nil,    // args
			)
			logger.LogIf(err)

			var forever chan struct{}

			go func() {
				for d := range msgs {
					log.Printf(" [x] Received message %s", d.Body)
				}
			}()

			log.Printf(" [*] Waiting for logs...")
			<-forever
		}()
	}
}
