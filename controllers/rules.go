package controllers

import (
	"github.com/gin-gonic/gin"
)

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
	c.JSON(200, gin.H{
		"msg": "SetRules",
	})
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
	c.JSON(200, gin.H{
		"msg": "GetRules",
	})
}
