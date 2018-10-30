package controllers

import (
	"github.com/gin-gonic/gin"
)

// Check  Verifica estado
/**
 * @api {get} /v1/ Check
 * @apiName FeedbackRating
 * @apiGroup Server Status
 *
 * @apiDescription Verify MicroService Status
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
