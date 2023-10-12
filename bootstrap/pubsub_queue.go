package bootstrap

import (
	"log"

	"github.com/czjge/gohub/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupPubsubQueue() {

	for i := 0; i < 2; i++ {

		go func(n int) {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			logger.LogIf(err)
			defer conn.Close()

			ch, err := conn.Channel()
			logger.LogIf(err)
			defer ch.Close()

			err = ch.ExchangeDeclare(
				"logs",   // name
				"fanout", // type
				true,     // durable
				false,    // auto-deleted
				false,    // internal
				false,    // no-wait
				nil,      // arguments
			)
			logger.LogIf(err)

			q, err := ch.QueueDeclare(
				"",    // name
				false, // durable
				false, // delete when unused
				true,  // exclusive
				false, // no-wait
				nil,   //arguments
			)
			logger.LogIf(err)

			err = ch.QueueBind(
				q.Name, // queue name
				// its value is ignored since the exchange type is fanout
				"",     // routing key
				"logs", // exchange
				false,
				nil,
			)
			logger.LogIf(err)

			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				true,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // arguments
			)
			logger.LogIf(err)

			var forever chan struct{}

			go func() {
				for d := range msgs {
					log.Printf(" [x] %s receive a message: %s", q.Name, d.Body)
				}
			}()

			log.Printf(" [*] %s waiting for logs...", q.Name)
			<-forever
		}(i)
	}
}
