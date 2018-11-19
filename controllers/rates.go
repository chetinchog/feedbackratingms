package controllers

import (
	"fmt"
	"time"

	"github.com/chetinchog/feedbackratingms/rates"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/chetinchog/feedbackratingms/tools/fn"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type getRateResponse struct {
	ArticleId  string    `json:"articleId" validate:"required"`
	Rate       float64   `json:"rate"`
	Ra1        int       `json:"ra1"`
	Ra2        int       `json:"ra2"`
	Ra3        int       `json:"ra3"`
	Ra4        int       `json:"ra4"`
	Ra5        int       `json:"ra5"`
	FeedAmount int       `json:"feedAmount"`
	BadRate    bool      `json:"badRate" validate:"required"`
	GoodRate   bool      `json:"goodRate" validate:"required"`
	Created    time.Time `json:"created" validate:"required"`
	Modified   time.Time `json:"modified" validate:"required"`
}

/**
 * @api {get} /v1/rates/:articleId/ Buscar Valoración de Artículo
 * @apiName FeedbackRating
 * @apiGroup Valoración
 *
 * @apiDescription ABM Reglas
 *
 * @apiSuccessExample {json} Response
 *		HTTP/1.1 200 OK
 *		{
 *			"articleId": "{article's id}",
 *			"rate": "{article rate's value}",
 *			"ra1": "{amount of rates with value 1}",
 *			"ra2": "{amount of rates with value 2}",
 *			"ra3": "{amount of rates with value 3}",
 *			"ra4": "{amount of rates with value 4}",
 *			"ra5": "{amount of rates with value 5}",
 *			"feedAmount": "{amount of feedbacks made}",
 *			"badRate": "{is this category (boolean)}",
 *			"goodRate": "{is this category (boolean)}",
 *			"created": "{creation date}",
 *			"modified": "{modification date}"
 *		}
 */
func GetRate(c *gin.Context) {
	articleId := c.Param("articleId")

	if articleId == "" {
		c.JSON(400, gin.H{
			"error": "ArticleId not sended",
		})
		return
	}

	dao, err := rates.GetDao()
	if err != nil {
		errors.Handle(c, err)
		return
	}

	rate, err := dao.FindByArticleID(articleId)
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		if rate == nil {
			c.JSON(400, gin.H{
				"error": "Article not rated",
			})
			return
		}
	}

	feedRate, feedAmount := fn.CalculateRates(rate)

	responseRate := getRateResponse{}
	responseRate.ArticleId = rate.ArticleId
	responseRate.Rate = feedRate
	responseRate.Ra1 = rate.Ra1
	responseRate.Ra2 = rate.Ra2
	responseRate.Ra3 = rate.Ra3
	responseRate.Ra4 = rate.Ra4
	responseRate.Ra5 = rate.Ra5
	responseRate.FeedAmount = feedAmount
	responseRate.BadRate = rate.BadRate
	responseRate.GoodRate = rate.GoodRate
	responseRate.Created = rate.Created
	responseRate.Modified = rate.Modified

	c.JSON(200, responseRate)
}

type getHistoryResponseArray struct {
	Rate    int       `json:"rate" validate:"required"`
	UserId  string    `json:"userId" validate:"required"`
	Created time.Time `json:"created" validate:"required"`
}

type getHistoryResponse struct {
	ArticleId string                    `json:"articleId" validate:"required"`
	History   []getHistoryResponseArray `json:"history"`
}

/**
 * @api {get} /v1/rates/:articlesd/history Buscar Historial de Artículo
 * @apiName FeedbackRating
 * @apiGroup Valoración
 *
 * @apiDescription ABM Reglas
 *
 * @apiSuccessExample {json} Response
 *		HTTP/1.1 200 OK
 *		{
 *			"articleId": "{article's id}",
 *			"history": [
 *			    {
 *			        "rate": "{rate's value}",
 *			        "userId": "{user's id}",
 *			        "created": "{creation date}"
 *			    }
 *			]
 *		}
 */
func GetHistory(c *gin.Context) {
	articleId := c.Param("articleId")

	if articleId == "" {
		c.JSON(400, gin.H{
			"error": "ArticleId not sended",
		})
		return
	}

	dao, err := rates.GetDao()
	if err != nil {
		errors.Handle(c, err)
		return
	}

	rate, err := dao.FindByArticleID(articleId)
	if err != nil {
		// errors.Handle(c, err)
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		if rate == nil {
			c.JSON(400, gin.H{
				"error": "Article not rated",
			})
			return
		}
	}

	responseHistory := getHistoryResponse{}
	responseHistory.ArticleId = rate.ArticleId

	historyArray := []getHistoryResponseArray{}
	for _, history := range rate.History {
		newHistory := getHistoryResponseArray{}
		newHistory.Rate = history.Rate
		newHistory.UserId = history.UserId
		newHistory.Created = history.Created
		historyArray = append(historyArray, newHistory)
	}

	responseHistory.History = historyArray

	c.JSON(200, responseHistory)
}

type Feedback struct {
	ID        objectid.ObjectID `json:"_id"`
	UserID    string            `json:"userId"`
	Text      string            `json:"text"`
	ArticleId string            `json:"productId"`
	Rate      int               `json:"rate" `
	Moderated bool              `json:"moderated"`
	Created   time.Time         `json:"created"`
	Updated   time.Time         `json:"updated"`
}
