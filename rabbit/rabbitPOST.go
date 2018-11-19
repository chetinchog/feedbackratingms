package rabbit

import (
	"encoding/json"
	"fmt"

	"github.com/chetinchog/feedbackratingms/tools/env"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/streadway/amqp"
)

type productParams struct {
	ArticleId string `json:"productId" binding:"required"`
}

/**
 * @api {topic} rates/high-rate Notificación de Valoración Alta
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Si una reseña supera la regla de una buena valoración, se notifica.
 *
 * @apiSuccessExample {json} Message
 *	{
 *	   "type": "high-rate",
 *	   "queue": "rates"
 *	   "message": {
 *	        "articleId" : "{article's id}",
 *	        "userId" : "{user's id}",
 *	        "rate": "{article rate's value}",
 *	    }
 *	}
 */
func HighRate(rateMessage string) error {
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

	msg := message{}
	msg.Type = "high-rate"
	msg.Message = rateMessage

	message, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = chn.ExchangeDeclare(
		"rates", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	queue, err := chn.QueueDeclare(
		"rates", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"rates",    // exchange
		false,
		nil)

	if err != nil {
		return err
	}

	err = chn.Publish(
		"rates", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ: Buena Calificación")

	return nil
}

/**
 * @api {topic} rates/low-rate Alerta de Valoración Baja
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Si una reseña supera la regla de una mala valoración, se genera una alerta.
 *
 * @apiSuccessExample {json} Message
 *	{
 *	   "type": "low-rate",
 *	   "queue": "rates"
 *	   "message": {
 *	        "articleId" : "{article's id}",
 *	        "userId" : "{user's id}",
 *	        "rate": "{article rate's value}",
 *	    }
 *	}
 */
func LowRate(rateMessage string) error {
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

	msg := message{}
	msg.Type = "low-rate"
	msg.Message = rateMessage

	message, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = chn.ExchangeDeclare(
		"rates", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	queue, err := chn.QueueDeclare(
		"rates", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"rates",    // exchange
		false,
		nil)

	if err != nil {
		return err
	}

	err = chn.Publish(
		"rates", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ: Mala Calificación")

	return nil
}

/**
 * @api {topic} rates/article-change-rate Notificación de cambio de Valoración de Artículo
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Se notifica cada vez que cambia el promedio de la valoración de un artículo.
 *
 * @apiSuccessExample {json} Message
 *	{
 *	   "type": "article-change-rate",
 *	   "queue": "rates"
 *	   "message": {
 *	        "articleId": "{article's id}",
 *	        "newRate": "{article rate's value}",
 *	        "feedAmount": "{amount of califications}"
 *	    }
 *	}
 */
func RateChange(rateMessage string) error {
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

	msg := message{}
	msg.Type = "article-change-rate"
	msg.Message = rateMessage

	message, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = chn.ExchangeDeclare(
		"rates", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	queue, err := chn.QueueDeclare(
		"rates", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"rates",    // exchange
		false,
		nil)

	if err != nil {
		return err
	}

	err = chn.Publish(
		"rates", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ: Cambio de Rating")

	return nil
}

/**
* @api {direct} cart/article-exist Product Validation
* @apiGroup RabbitMQ POST
*
* @apiDescription Sending a validation request for a product.
*
* @apiSuccessExample {json} Message
*     {
*			"type": "article-exist",
*			"queue": "catalog",
*			"exchange": "",
*			"message" : {
*				"referenceId": "{referenceId}",
*             	"articleId": "{articleId}",
*			}
*		}
 */
func ProductValidation(articleId string, feedbackID objectid.ObjectID) error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer chn.Close()

	msg := response{}
	msg.Message.ArticleId = articleId
	msg.Exchange = "rating-article"
	msg.Queue = "rating-article"
	msg.Type = "article-exist"
	feed, err := json.Marshal(feedbackID)

	if err != nil {
		fmt.Println(err)
		return err
	}

	msg.Message.FeedbackID = string(feed)

	resp, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = chn.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
		false,     // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = chn.Publish(
		"catalog", // exchange
		"catalog", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(resp),
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("RabbitMQ: Envio de validación")

	return nil
}
