package rates

import (
	"context"
	"log"
	"time"

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
	FindAll() ([]*Rate, error)
	FindByID(rateID string) (*Rate, error)
	FindByLogin(login string) (*Rate, error)
}

// New dao es interno a este modulo, nadie fuera del modulo tiene acceso
func newDao() (Dao, error) {
	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("rates")

	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.String("login", ""),
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

// MockedDao sirve para poder mockear el db.Collection y testear el modulo
func MockedDao(coll db.Collection) Dao {
	return daoStruct{
		collection: coll,
	}
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

	rate.Updated = time.Now()

	doc, err := bson.NewDocumentEncoder().EncodeDocument(rate)
	if err != nil {
		return nil, err
	}

	_, err = d.collection.UpdateOne(context.Background(),
		bson.NewDocument(doc.LookupElement("_id")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				doc.LookupElement("password"),
				doc.LookupElement("name"),
				doc.LookupElement("enabled"),
				doc.LookupElement("updated"),
				doc.LookupElement("permissions"),
			),
		))

	if err != nil {
		return nil, err
	}

	return rate, nil
}

// FindAll devuelve todos los usuarios
func (d daoStruct) FindAll() ([]*Rate, error) {
	filter := bson.NewDocument()
	cur, err := d.collection.Find(context.Background(), filter, nil)
	defer cur.Close(context.Background())

	if err != nil {
		return nil, err
	}

	rates := []*Rate{}
	for cur.Next(context.Background()) {
		rate := &Rate{}
		if err := cur.Decode(rate); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindByID(rateID string) (*Rate, error) {
	_id, err := objectid.FromHex(rateID)
	if err != nil {
		return nil, errors.ErrID
	}

	rate := &Rate{}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", _id))
	if err = d.collection.FindOne(context.Background(), filter).Decode(rate); err != nil {
		return nil, err
	}

	return rate, nil
}

// FindByLogin lee un usuario desde la db
func (d daoStruct) FindByLogin(login string) (*Rate, error) {
	rate := &Rate{}
	filter := bson.NewDocument(bson.EC.String("login", login))
	err := d.collection.FindOne(context.Background(), filter).Decode(rate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrLogin
		}
		return nil, err
	}

	return rate, nil
}
