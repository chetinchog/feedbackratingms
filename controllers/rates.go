package controllers

import (
	"github.com/gin-gonic/gin"
)

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
	c.JSON(200, gin.H{
		"msg": "GetRate",
	})
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
 *			"id": "{article's id}",
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
	c.JSON(200, gin.H{
		"msg": "GetHistory",
	})
}
