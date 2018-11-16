package controllers

import (
	"fmt"
	"strings"

	"github.com/chetinchog/feedbackratingms/security"
	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/gin-gonic/gin"
)

// get token from Authorization header
func getTokenHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "Bearer ") != 0 {
		return "", errors.Unauthorized
	}
	return tokenString[7:], nil
}

func validateAuthentication(c *gin.Context) error {
	tokenString, err := getTokenHeader(c)
	if err != nil {
		return errors.Unauthorized
	}

	fmt.Println(tokenString)
	if _, err = security.Validate(tokenString); err != nil {
		return errors.Unauthorized
	}

	return nil
}
