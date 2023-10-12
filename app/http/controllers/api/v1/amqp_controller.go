package v1

import (
	"context"
	"log"
	"time"

	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/response"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPController struct {
	BaseAPIControler
}

func (ctrl *AMQPController) WorkSend(c *gin.Context) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	logger.LogIf(err)
	defer conn.Close()

	ch, err := conn.Channel()
	logger.LogIf(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	logger.LogIf(err)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := c.PostForm("message")
	err = ch.PublishWithContext(ctx,
		// Here we use the default or nameless exchange: messages are routed to the queue with the name specified by routing_key parameter, if it exists.
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	logger.LogIf(err)
	log.Printf(" [x] Sent %s\n", body)

	response.Success(c)
}

func (ctrl *AMQPController) PubsubSend(c *gin.Context) {

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
		false,    // no-wati
		nil,      // arguments
	)
	logger.LogIf(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := c.PostForm("message")
	err = ch.PublishWithContext(ctx,
		"logs", // exchange
		// its value is ignored since the exchange type is fanout
		"",    // routing_key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	logger.LogIf(err)

	log.Printf(" [x] Sent %s\n", body)

	response.Success(c)
}

func (ctrl *AMQPController) RoutingSend(c *gin.Context) {

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := c.PostForm("message")
	severity := c.PostForm("severity")
	err = ch.PublishWithContext(ctx,
		"logs_direct", // exchange
		severity,      // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	logger.LogIf(err)

	log.Printf(" [x] Sent %s", body)

	response.Success(c)
}
