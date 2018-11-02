package controllers

import (
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /v1/ Check
 * @apiName FeedbackRating
 * @apiGroup Sistema
 *
 * @apiDescription Verifica estado del Sistema
 *
 * @apiSuccessExample {json} Response
 *		HTTP/1.1 200 OK
 *		{
 *			"msg": "Running",
 *		}
 */
func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "✔",
	})
}
