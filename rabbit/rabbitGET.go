package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chetinchog/feedbackratingms/security"
	"github.com/chetinchog/feedbackratingms/tools/env"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.NewCustom(400, "Channel not initialized")

type message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type feedbackParams struct {
	ProductID string `json:"productID"`
	CartID    string `json:"cardID"`
}

// Init se queda escuchando broadcasts de logout
func Init() {
	go func() {
		for {
			listenNewFeedback()
			fmt.Println("RabbitMQ: Listen New Feedback -> Reconnect in 5...")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			listenLogout()
			fmt.Println("RabbitMQ: Listen Logout -> Reconnect in 5...")
			time.Sleep(5 * time.Second)
		}
	}()
}

/**
 * @api {direct} feedback/new-feedback Buscar Reseña
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Escucha los mensajes de creación de Feedback para obtener las valoraciones
 *
 * @apiSuccessExample {json} Message
 * 		{
 *    		"type": "new-feedback",
 *    		"message": {
 *        		"id" : "{feedback's id}"
*      		 	"idUser" : "{user's id}",
 *        		"idProduct" : "{product's id}",
 *        		"rate" : "{feedback's rate}",
 *        		"created" : "{creation date}",
 *        		"modified" : "{modification date}"
 *    		}
 *		}
*/
func listenNewFeedback() error {

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

	queue, err := chn.QueueDeclare(
		"feedback", // name
		false,      // durable
		false,      // delete when usused
		true,       // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name,       // queue name
		"feedback",       // routing key
		"feedback_topic", // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	msgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ: Listening New Feedback")

	go func() {
		for d := range msgs {
			newMessage := &message{}
			err = json.Unmarshal(d.Body, newMessage)
			if err == nil {
				if newMessage.Type == "new-feedback" {
					log.Output(1, "RabbitMQ: New Feedback Message")
					fmt.Println(newMessage)
				}
			}
		}
	}()

	fmt.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

/**
 * @api {direct} feedback/article-validation Buscar Validación de Artículo
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Listen validation product messages from catalog.
 *
 * @apiSuccessExample {json} Message
 * 		{
 *      	"type": "article-exist",
 *			"message" :
 *				{
 *					"articleId": "{articleId}",
 *					"valid": true|false
 *				}
 *      }
 */
func listenProductValidation() error {
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

	msg, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		return err
	}

	for d := range msg {
		log.Printf("Received a message: %s", d.Body)
	}

	return nil
}

/**
 * @api {fanout} auth/logout Logout de Usuarios
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Escucha de mensajes de logout desde auth.
 *
 * @apiSuccessExample {json} Message
 *     {
 *        "type": "logout",
 *        "message": "{tokenId}"
 *     }
 */
func listenLogout() error {
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
		"auth",   // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	queue, err := chn.QueueDeclare(
		"auth", // name
		false,  // durable
		false,  // delete when unused
		true,   // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	fmt.Println("RabbitMQ: Listening Logout")

	go func() {
		for d := range mgs {
			log.Output(1, "Mensage recibido")
			newMessage := &message{}
			err = json.Unmarshal(d.Body, newMessage)
			if err == nil {
				if newMessage.Type == "logout" {
					security.Invalidate(newMessage.Message)
				}
			}
		}
	}()

	fmt.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}
