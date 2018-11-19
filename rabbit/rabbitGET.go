package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chetinchog/feedbackratingms/controllers"
	"github.com/chetinchog/feedbackratingms/rates"
	"github.com/chetinchog/feedbackratingms/security"
	"github.com/chetinchog/feedbackratingms/tools/env"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/chetinchog/feedbackratingms/tools/fn"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.NewCustom(400, "Channel not initialized")

type message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type msj struct {
	FeedbackID string `json:"referenceId"`
	ArticleId  string `json:"articleId"`
}

type response struct {
	Type     string `json:"type"`
	Exchange string `json:"exchange"`
	Queue    string `json:"queue"`
	Message  msj    `json:"message"`
}

type messageValidation struct {
	ArticleId string `json:"articleId"`
	Valid     bool   `json:"valid"`
}

type msgRateChange struct {
	ArticleId  string  `json:"articleId"`
	NewRate    float64 `json:"newRate"`
	FeedAmount int     `json:"feedAmount"`
}

type msgClasification struct {
	ArticleId string  `json:"articleId"`
	Rate      float64 `json:"rate"`
}

// Init se queda escuchando broadcasts de logout
func Init() {
	go func() {
		for {
			err := listenNewFeedback()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RabbitMQ: Listen New Feedback -> Reconnect in 5...")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := listenLogout()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RabbitMQ: Listen Logout -> Reconnect in 5...")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := listenProductValidation()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RabbitMQ: Listen Product Validation -> Reconnect in 5...")
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
		true,       // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name,       // queue name
		"",               // routing key
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
					NewFeedback(newMessage.Message)
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

	err = chn.ExchangeDeclare(
		"rating-article", // name
		"direct",         // type
		false,            // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)

	queue, err := chn.QueueDeclare(
		"rating-article", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name,       // queue name
		"",               // routing key
		"rating-article", // exchange
		false,
		nil)

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

	fmt.Println("RabbitMQ: Listening Product Validation")

	for d := range msg {
		log.Printf("Received a message Validation Product: %s", d.Body)
		newMessage := &message{}
		err = json.Unmarshal(d.Body, newMessage)
		ValidateProduct(newMessage.Message)
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
		false,  // exclusive
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

func NewFeedback(feed string) {

	newFeed := &controllers.Feedback{}

	err := json.Unmarshal([]byte(feed), newFeed)
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return
	}

	dao, err := rates.GetDao()
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return
	}

	articleId := newFeed.ArticleId

	rate, err := dao.FindByArticleID(articleId)
	if err != nil {
	}

	userIdNF := newFeed.UserID
	rateNF := newFeed.Rate
	if rate == nil {
		rate = rates.NewRate()
		rate.ArticleId = articleId

		newHistory := rates.NewHistory()
		newHistory.Rate = rateNF
		newHistory.UserId = userIdNF
		rate.History = append(rate.History, newHistory)

		rate = fn.AddRate(rate, rateNF)
		rate = fn.Classify(rate)
		rate.Enabled = false

		if err := ProductValidation(newFeed.ArticleId, newFeed.ID); err != nil {
			return
		}

		newRule, err := dao.Insert(rate)
		if err != nil || newRule == nil {
			fmt.Println(" ---------------------------- ")
			fmt.Println(err)
			fmt.Println(" ---------------------------- ")
			return
		}
	} else {

		newHistory := rates.NewHistory()
		newHistory.Rate = rateNF
		newHistory.UserId = userIdNF
		rate.History = append(rate.History, newHistory)

		rate = fn.AddRate(rate, rateNF)
		rate = fn.Classify(rate)

		newRule, err := dao.Update(rate)
		if err != nil || newRule == nil {
			fmt.Println(" ---------------------------- ")
			fmt.Println(err)
			fmt.Println(" ---------------------------- ")
			return
		}
	}
}

func ValidateProduct(validation string) {
	newMessage := &messageValidation{}
	err := json.Unmarshal([]byte(validation), newMessage)
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return
	}

	dao, err := rates.GetDao()
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return
	}

	articleId := newMessage.ArticleId

	rate, err := dao.FindByArticleID(articleId)
	if err != nil {
	}

	if rate != nil {
		if newMessage.Valid == true {
			err = dao.EnableByArticleId(articleId)

			rating, amount := fn.CalculateRates(rate)

			msgRate := &msgRateChange{
				ArticleId:  articleId,
				NewRate:    rating,
				FeedAmount: amount,
			}
			rateMsg, err := json.Marshal(msgRate)
			if err != nil {
				return
			}
			RateChange(string(rateMsg[:]))

			msgRateClass := &msgClasification{
				ArticleId: articleId,
				Rate:      rating,
			}
			rateClassMsg, err := json.Marshal(msgRateClass)
			if err != nil {
				return
			}
			if rate.GoodRate == true {
				HighRate(string(rateClassMsg[:]))
			} else if rate.BadRate == true {
				LowRate(string(rateClassMsg[:]))
			}

		} else {
			err = dao.DisableByArticleId(articleId)
		}
	}
}
