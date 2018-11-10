package rates

import (
	"context"
	"fmt"
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
	fmt.Println(rate.ID)

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

	_, err = d.collection.UpdateOne(context.Background(),
		bson.NewDocument(doc.LookupElement("_id")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				doc.LookupElement("articleId"),
				doc.LookupElement("rate"),
				doc.LookupElement("ra1"),
				doc.LookupElement("ra2"),
				doc.LookupElement("ra3"),
				doc.LookupElement("ra4"),
				doc.LookupElement("ra5"),
				doc.LookupElement("feedAmount"),
				doc.LookupElement("badRate"),
				doc.LookupElement("goodRate"),
				doc.LookupElement("history"),
				doc.LookupElement("created"),
				doc.LookupElement("modified"),
				doc.LookupElement("enabled"),
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
