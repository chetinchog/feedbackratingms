package rates

import (
	"time"

	"github.com/chetinchog/feedbackratingms/tools/errors"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	validator "gopkg.in/go-playground/validator.v9"
)

type History struct {
	Rate    int       `bson:"rate" validate:"required"`
	UserId  string    `bson:"userId" validate:"required"`
	Created time.Time `bson:"created" validate:"required"`
}

// Rate data structure
type Rate struct {
	ID        objectid.ObjectID `bson:"_id"`
	ArticleId string            `bson:"articleId" validate:"required"`
	Ra1       int               `bson:"ra1" validate:"required"`
	Ra2       int               `bson:"ra2" validate:"required"`
	Ra3       int               `bson:"ra3" validate:"required"`
	Ra4       int               `bson:"ra4" validate:"required"`
	Ra5       int               `bson:"ra5" validate:"required"`
	BadRate   bool              `bson:"badRate" validate:"required"`
	GoodRate  bool              `bson:"goodRate" validate:"required"`
	History   []*History        `bson:"history"`
	Created   time.Time         `bson:"created" validate:"required"`
	Modified  time.Time         `bson:"modified" validate:"required"`
	Enabled   bool              `bson:"enabled" validate:"required"`
}

func NewRate() *Rate {
	return &Rate{
		ID:       objectid.New(),
		Ra1:      0,
		Ra2:      0,
		Ra3:      0,
		Ra4:      0,
		Ra5:      0,
		BadRate:  false,
		GoodRate: false,
		History:  []*History{},
		Created:  time.Now(),
		Modified: time.Now(),
		Enabled:  true,
	}
}

func NewHistory() *History {
	return &History{
		Created: time.Now(),
		Rate:    0,
		UserId:  "",
	}
}

var ErrData = errors.NewValidationField("rule", "invalid")

func (e *Rate) ValidateSchema() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}
