package rabbit

import (
	"github.com/chetinchog/feedbackratingms/tools/env"
	"github.com/streadway/amqp"
)

type productParams struct {
	ProductId string `json:"productId" binding:"required"`
}

/**
 * @api {direct} cart/article-exist Product Validation
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Sending a validation request for a product.
 *
 * @apiSuccessExample {json} Message
 *     {
			"type": "article-exist",
			"queue": "cart",
			"exchange": "cart",
			"message" : {
				"referenceId": "{cartId}",
 *             	"articleId": "{articleId}",
			}
		}
*/
func ProductValidation(productId string, cartID string) error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		return err
	}
	defer chn.Close()

	queue, err := chn.QueueDeclare(
		"catalog", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msg := productParams{}
	msg.ProductId = productId

	err = chn.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(`msg`),
		},
	)

	if err != nil {
		return err
	}

	return err
}

/**
 * @api {fanout} feedback/ Send Feedback
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Sending new feedback.
 *
 * @apiSuccessExample {json} Message
 *     {
			"type": "article-exist",
			"queue": "feedback",
			"exchange": "feedback",
			"message" : {
				"articleId": "{articleId}",
			}
		}
*/
/*
func sendFeedback(feedback string) error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"feedback_topic", // name
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return err
	}

	err = chn.Publish(
		"feedback_topic", // exchange
		"feedback",       // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(feedback),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
*/
