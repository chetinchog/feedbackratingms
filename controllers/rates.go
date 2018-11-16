package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/chetinchog/feedbackratingms/rates"
	"github.com/chetinchog/feedbackratingms/rules"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/gin-gonic/gin"
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

func calculateRates(feedRate *rates.Rate) (float64, int) {
	feedAmount := (feedRate.Ra1 +
		feedRate.Ra2 +
		feedRate.Ra3 +
		feedRate.Ra4 +
		feedRate.Ra5)

	rate := (float64(feedRate.Ra1)*1 +
		float64(feedRate.Ra2)*2 +
		float64(feedRate.Ra3)*3 +
		float64(feedRate.Ra4)*4 +
		float64(feedRate.Ra5)*5) / float64(feedAmount)
	return rate, feedAmount
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

	feedRate, feedAmount := calculateRates(rate)

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
	UserID    string `json:"userId" binding:"required"`
	Text      string `json:"text" binding:"required"`
	ProductID string `json:"productId" binding:"required"`
	Rate      int    `json:"rate" binding:"required"`
}

func addRate(rate *rates.Rate, rateNF int) *rates.Rate {
	switch rateNF {
	case 1:
		rate.Ra1++
		break
	case 2:
		rate.Ra2++
		break
	case 3:
		rate.Ra3++
		break
	case 4:
		rate.Ra4++
		break
	case 5:
		rate.Ra5++
		break
	default:
	}
	return rate
}

func classify(rate *rates.Rate) *rates.Rate {
	prom, amount := calculateRates(rate)
	if amount == 0 {
	}

	dao, err := rules.GetDao()
	if err != nil {
		fmt.Println(" ---------------------------- ")
		fmt.Println(err)
		fmt.Println(" ---------------------------- ")
		return rate
	}

	articleId := rate.ArticleId

	rule, err := dao.FindByArticleID(articleId)
	if err != nil {
	}

	if rule == nil {
		return rate
	}

	if prom <= float64(rule.LowRate) {
		rate.BadRate = true
	} else {
		rate.BadRate = false
	}
	if prom >= float64(rule.HighRate) {
		rate.GoodRate = true
	} else {
		rate.GoodRate = false
	}

	return rate
}

func NewFeedback(feed string) {

	newFeed := &Feedback{}

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

	articleId := newFeed.ProductID

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

		rate = addRate(rate, rateNF)
		rate = classify(rate)

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

		rate = addRate(rate, rateNF)
		rate = classify(rate)

		newRule, err := dao.Update(rate)
		if err != nil || newRule == nil {
			fmt.Println(" ---------------------------- ")
			fmt.Println(err)
			fmt.Println(" ---------------------------- ")
			return
		}
	}
}
