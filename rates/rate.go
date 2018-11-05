package rates

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	validator "gopkg.in/go-playground/validator.v9"
)

// Rate data structure
type Rate struct {
	ID         objectid.ObjectID `bson:"_id"`
	ArticleId  string            `bson:"articleId" validate:"required"`
	Rate       float32           `bson:"rate"`
	Ra1        int               `bson:"ra1"`
	Ra2        int               `bson:"ra2"`
	Ra3        int               `bson:"ra3"`
	Ra4        int               `bson:"ra4"`
	Ra5        int               `bson:"ra5"`
	FeedAmount int               `bson:"feedAmount"`
	BadRate    int               `bson:"badRate"`
	GoodRate   int               `bson:"goodRate"`
	Created    time.Time         `bson:"created"`
	Modified   time.Time         `bson:"modified"`
	Enabled    bool              `bson:"enabled"`
}

func NewRate() *Rate {
	return &Rate{
		ID:      objectid.New(),
		Enabled: true,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

func (e *Rate) ValidateSchema() error {
	validate := validator.New()
	return validate.Struct(e)
}
