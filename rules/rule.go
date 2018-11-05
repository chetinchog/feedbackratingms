package rules

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	validator "gopkg.in/go-playground/validator.v9"
)

// Rule data structure
type Rule struct {
	ID        objectid.ObjectID `bson:"_id"`
	ArticleId string            `bson:"articleId" validate:"required"`
	LowRate   int               `bson:"lowRate"`
	HighRate  int               `bson:"highRate"`
	Created   time.Time         `bson:"created"`
	Modified  time.Time         `bson:"modified"`
	Enabled   bool              `bson:"enabled"`
}

func NewRule() *Rule {
	return &Rule{
		ID:       objectid.New(),
		Enabled:  true,
		Created:  time.Now(),
		Modified: time.Now(),
	}
}

func (e *Rule) ValidateSchema() error {
	validate := validator.New()
	return validate.Struct(e)
}
