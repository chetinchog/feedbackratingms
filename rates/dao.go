package rates

import (
	"context"
	"log"
	"time"

	"github.com/chetinchog/feedbackratingms/tools/db"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type daoStruct struct {
	collection db.Collection
}

// Dao es la interface que exponse los servicios de acceso a la DB
type Dao interface {
	Insert(rate *Rate) (*Rate, error)
	Update(rate *Rate) (*Rate, error)
	FindAll() ([]Rate, error)
	FindByID(rateID string) (*Rate, error)
	FindByArticleID(articleId string) (*Rate, error)
	Disable(_id objectid.ObjectID) error
	DisableByArticleId(articleId string) error
	Enable(_id objectid.ObjectID) error
	EnableByArticleId(articleId string) error
}

// New dao es interno a este modulo, nadie fuera del modulo tiene acceso
func GetDao() (Dao, error) {
	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("rates")

	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.String("articleId", ""),
			),
			Options: bson.NewDocument(
				bson.EC.Boolean("unique", true),
			),
		},
	)
	if err != nil {
		log.Output(1, err.Error())
	}

	coll := db.WrapCollection(collection)
	return daoStruct{
		collection: coll,
	}, nil
}

func (d daoStruct) Insert(rate *Rate) (*Rate, error) {
	if err := rate.ValidateSchema(); err != nil {
		return nil, err
	}

	if _, err := d.collection.InsertOne(context.Background(), rate); err != nil {
		return nil, err
	}

	return rate, nil
}

func (d daoStruct) Update(rate *Rate) (*Rate, error) {
	if err := rate.ValidateSchema(); err != nil {
		return nil, err
	}

	rate.Modified = time.Now()

	doc, err := db.EncodeDocument(rate)
	if err != nil {
		return nil, err
	}

	ra1 := doc.LookupElement("ra1")
	if ra1 == nil {
		ra1 = bson.EC.Int32("ra1", 0)
	}
	ra2 := doc.LookupElement("ra2")
	if ra2 == nil {
		ra2 = bson.EC.Int32("ra2", 0)
	}
	ra3 := doc.LookupElement("ra3")
	if ra3 == nil {
		ra3 = bson.EC.Int32("ra3", 0)
	}
	ra4 := doc.LookupElement("ra4")
	if ra4 == nil {
		ra4 = bson.EC.Int32("ra4", 0)
	}
	ra5 := doc.LookupElement("ra5")
	if ra5 == nil {
		ra5 = bson.EC.Int32("ra5", 0)
	}

	badRate := doc.LookupElement("badRate")
	if badRate == nil {
		badRate = bson.EC.Boolean("badRate", false)
	}
	goodRate := doc.LookupElement("goodRate")
	if goodRate == nil {
		goodRate = bson.EC.Boolean("goodRate", false)
	}

	_, err = d.collection.UpdateOne(context.Background(),
		bson.NewDocument(doc.LookupElement("_id")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				doc.LookupElement("articleId"),
				ra1,
				ra2,
				ra3,
				ra4,
				ra5,
				badRate,
				goodRate,
				doc.LookupElement("history"),
				doc.LookupElement("created"),
				doc.LookupElement("modified"),
				bson.EC.Boolean("enabled", true),
			),
		))

	if err != nil {
		return nil, err
	}

	return rate, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindAll() ([]Rate, error) {
	result, err := d.collection.Find(context.Background(), nil)
	defer result.Close(context.Background())
	if err != nil {
		return nil, err
	}

	rates := []Rate{}
	for result.Next(context.Background()) {
		rate := Rate{}
		result.Decode(rate)
		rates = append(rates, rate)
	}
	return rates, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindByID(rateID string) (*Rate, error) {
	_id, err := objectid.FromHex(rateID)
	if err != nil {
		return nil, err
	}

	rate := &Rate{}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", _id))
	if err := d.collection.FindOne(context.Background(), filter).Decode(rate); err != nil {
		return nil, err
	}

	return rate, nil
}

// FindByArticleID lee un usuario desde la db filtrando por artículo
func (d daoStruct) FindByArticleID(articleId string) (*Rate, error) {
	rate := &Rate{}
	filter := bson.NewDocument(bson.EC.String("articleId", articleId))
	if err := d.collection.FindOne(context.Background(), filter).Decode(rate); err != nil {
		return nil, err
	}

	return rate, nil
}

func (d daoStruct) Disable(_id objectid.ObjectID) error {
	_, err := d.collection.UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", _id)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.Boolean("enabled", false),
			),
		))
	if err != nil {
		return err
	}
	return nil
}

func (d daoStruct) DisableByArticleId(articleId string) error {
	_, err := d.collection.UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.String("articleId", articleId)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.Boolean("enabled", false),
			),
		))
	if err != nil {
		return err
	}
	return nil
}

func (d daoStruct) Enable(_id objectid.ObjectID) error {
	_, err := d.collection.UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", _id)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.Boolean("enabled", true),
			),
		))
	if err != nil {
		return err
	}
	return nil
}

func (d daoStruct) EnableByArticleId(articleId string) error {
	_, err := d.collection.UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.String("articleId", articleId)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.Boolean("enabled", true),
			),
		))
	if err != nil {
		return err
	}
	return nil
}
