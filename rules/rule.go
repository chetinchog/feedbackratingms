package rules

import (
	"time"

	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	validator "gopkg.in/go-playground/validator.v9"
)

// Rule data structure
type Rule struct {
	ID        objectid.ObjectID `bson:"_id"`
	ArticleId string            `bson:"articleId" validate:"required"`
	LowRate   int               `bson:"lowRate"`
	HighRate  int               `bson:"highRate"`
	Created   time.Time         `bson:"created" validate:"required"`
	Modified  time.Time         `bson:"modified" validate:"required"`
	Enabled   bool              `bson:"enabled" validate:"required"`
}

func NewRule() *Rule {
	return &Rule{
		ID:       objectid.New(),
		LowRate:  0,
		HighRate: 0,
		Created:  time.Now(),
		Modified: time.Now(),
		Enabled:  true,
	}
}

var ErrData = errors.NewValidationField("rule", "invalid")

func (e *Rule) ValidateSchema() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}
