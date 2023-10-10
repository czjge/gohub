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
