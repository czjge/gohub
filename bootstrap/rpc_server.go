package bootstrap

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/czjge/gohub/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func SetupRPCServer() {

	for i := 0; i < 1; i++ {

		go func() {

			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			logger.LogIf(err)
			defer conn.Close()

			ch, err := conn.Channel()
			logger.LogIf(err)
			defer ch.Close()

			q, err := ch.QueueDeclare(
				"rpc_queue", // name
				false,       // durable
				false,       // delete when unused
				false,       // exclusive
				false,       // no-wait
				nil,         // arguments
			)
			logger.LogIf(err)

			err = ch.Qos(
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

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				for d := range msgs {
					n, err := strconv.Atoi(string(d.Body))
					logger.LogIf(err)

					log.Printf(" [.] fib(%d)", n)
					response := fib(n)

					err = ch.PublishWithContext(ctx,
						"",        // exchange
						d.ReplyTo, // routing key
						false,     // mandatory
						false,     // immediate
						amqp.Publishing{
							ContentType:   "text/plain",
							CorrelationId: d.CorrelationId,
							Body:          []byte(strconv.Itoa(response)),
						},
					)
					logger.LogIf(err)

					d.Ack(false)
				}
			}()

			log.Printf(" [*] Waiting for RPC requests...")
			<-forever
		}()
	}
}
