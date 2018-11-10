package controllers

import (
	"fmt"
	"time"

	"github.com/chetinchog/feedbackratingms/rules"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/gin-gonic/gin"
)

type setRulesRequest struct {
	LowRate  int `json:"lowRate"`
	HighRate int `json:"highRate"`
}

type setRulesResponse struct {
	ArticleId string    `json:"articleId"`
	LowRate   int       `json:"lowRate"`
	HighRate  int       `json:"highRate"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

func responseSetRules(c *gin.Context, rule *rules.Rule) {
	responseRule := setRulesResponse{}
	responseRule.ArticleId = rule.ArticleId
	responseRule.LowRate = rule.LowRate
	responseRule.HighRate = rule.HighRate
	responseRule.Created = rule.Created
	responseRule.Modified = rule.Modified
	c.JSON(200, responseRule)
}

/**
 * @api {post} /v1/rates/:articleId/rules Asignar Parámetro a Artículo
 * @apiName FeedbackRating
 * @apiGroup Reglas de Valoración
 *
 * @apiDescription ABM Reglas
 *
 * @apiExample {json} Body
 *    {
 *      "lowRate": "{bad rate's value}",
 *		"highRate": "{good rate's value}"
 *    }
 *
 * @apiSuccessExample {json} Response
 *		HTTP/1.1 200 OK
 *		{
 *   		"articleId": "{article's id}",
 *  		"lowRate": "{bad rate's value}",
 * 			"highRate": "{good rate's value}",
 *			"created": "{creation date}",
 *   		"modified": "{modification date}"
 *		}
 */
func SetRules(c *gin.Context) {
	fmt.Println("SetRules")
	body := setRulesRequest{}

	// if err := validateAuthentication(c); err != nil {
	// 	errors.Handle(c, err)
	// 	return
	// }

	if err := c.ShouldBindJSON(&body); err != nil {
		errors.Handle(c, err)
		return
	}

	articleId := c.Param("articleId")

	dao, err := rules.GetDao()
	if err != nil {
		errors.Handle(c, err)
		return
	}

	rule, err := dao.FindByArticleID(articleId)
	if err != nil {
		// errors.Handle(c, err)
	}

	if rule == nil {
		rule = rules.NewRule()
		rule.ArticleId = articleId
		rule.LowRate = body.LowRate
		rule.HighRate = body.HighRate
		newRule, err := dao.Insert(rule)
		if err != nil {
			errors.Handle(c, err)
			return
		}
		responseSetRules(c, newRule)
	} else {
		rule.LowRate = body.LowRate
		rule.HighRate = body.HighRate
		newRule, err := dao.Update(rule)
		if err != nil {
			errors.Handle(c, err)
			return
		}
		responseSetRules(c, newRule)
	}
}

type getRulesResponse struct {
	ArticleId string    `json:"articleId"`
	LowRate   int       `json:"lowRate"`
	HighRate  int       `json:"highRate"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

/**
 * @api {get} /v1/rates/:articleId/rules Buscar Parámetro a Artículo
 * @apiName FeedbackRating
 * @apiGroup Reglas de Valoración
 *
 * @apiDescription Get Reglas
 *
 * @apiSuccessExample {json} Response
 *		HTTP/1.1 200 OK
 *		{
 *   		"articleId": "{article's id}",
 *  		"lowRate": "{bad rate's value}",
 * 			"highRate": "{good rate's value}",
 *			"created": "{creation date}",
 *   		"modified": "{modification date}"
 *		}
 */
func GetRules(c *gin.Context) {
	articleId := c.Param("articleId")

	dao, err := rules.GetDao()
	if err != nil {
		errors.Handle(c, err)
		return
	}

	rule, err := dao.FindByArticleID(articleId)
	if err != nil {
		// errors.Handle(c, err)
	}

	responseRule := setRulesResponse{}
	responseRule.ArticleId = rule.ArticleId
	responseRule.LowRate = rule.LowRate
	responseRule.HighRate = rule.HighRate
	responseRule.Created = rule.Created
	responseRule.Modified = rule.Modified

	c.JSON(200, responseRule)
}
