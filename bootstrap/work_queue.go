package bootstrap

import (
	"bytes"
	"log"
	"time"

	"github.com/czjge/gohub/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupWorkQueue() {

	for i := 0; i < 2; i++ {

		go func(n int) {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			logger.LogIf(err)
			defer conn.Close()

			ch, err := conn.Channel()
			logger.LogIf(err)
			defer ch.Close()

			q, err := ch.QueueDeclare(
				// we needed to point the workers to the same queue by name
				"task_queue", // name
				true,         // durable
				false,        // delete when unused
				false,        // exclusive
				false,        // no-wait
				nil,          // arguments
			)
			logger.LogIf(err)

			err = ch.Qos(
				// don't dispatch a new message to a worker until it has processed and acknowledged the previous one
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			logger.LogIf(err)

			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			logger.LogIf(err)

			var forever chan struct{}

			go func() {
				for d := range msgs {
					log.Printf("Worker [%d] receive a message: %s", n, d.Body)
					dotCount := bytes.Count(d.Body, []byte("."))
					t := time.Duration(dotCount)
					time.Sleep(t * time.Second)
					log.Printf("Done")
					d.Ack(false) // this acknowledges a single delivery
				}
			}()

			log.Printf(" [*] Worker [%d] waiting for messages...", n)
			<-forever
		}(i)
	}
}
