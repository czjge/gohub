package bootstrap

import (
	"log"

	"github.com/czjge/gohub/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupRoutingQueue() {

	for i := 0; i < 1; i++ {

		go func() {

			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			logger.LogIf(err)
			defer conn.Close()

			ch, err := conn.Channel()
			logger.LogIf(err)
			defer ch.Close()

			err = ch.ExchangeDeclare(
				"logs_direct", // name
				"direct",      // type
				true,          // durable
				false,         // auto-deleted
				false,         // internal
				false,         // no-wait
				nil,           // arguments
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

			bindings := []string{"info", "warning", "error"}
			for _, s := range bindings {
				log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", s)
				err = ch.QueueBind(
					q.Name,        // queue name
					s,             // routing key
					"logs_direct", // exchange
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
				false,  // no local
				false,  // no wait
				nil,    // args
			)
			logger.LogIf(err)

			var forever chan struct{}

			go func() {
				for d := range msgs {
					log.Printf(" [x] received %s", d.Body)
				}
			}()

			log.Printf(" [*] Waiting for logs...")
			<-forever
		}()
	}
}
