package rules

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
	Insert(rule *Rule) (*Rule, error)
	Update(rule *Rule) (*Rule, error)
	FindAll() ([]*Rule, error)
	FindByID(ruleID string) (*Rule, error)
	FindByLogin(login string) (*Rule, error)
}

// New dao es interno a este modulo, nadie fuera del modulo tiene acceso
func newDao() (Dao, error) {
	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("rules")

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

func (d daoStruct) Insert(rule *Rule) (*Rule, error) {
	if err := rule.ValidateSchema(); err != nil {
		return nil, err
	}

	if _, err := d.collection.InsertOne(context.Background(), rule); err != nil {
		return nil, err
	}

	return rule, nil
}

func (d daoStruct) Update(rule *Rule) (*Rule, error) {
	if err := rule.ValidateSchema(); err != nil {
		return nil, err
	}

	rule.Updated = time.Now()

	doc, err := bson.NewDocumentEncoder().EncodeDocument(rule)
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

	return rule, nil
}

// FindAll devuelve todos los usuarios
func (d daoStruct) FindAll() ([]*Rule, error) {
	filter := bson.NewDocument()
	cur, err := d.collection.Find(context.Background(), filter, nil)
	defer cur.Close(context.Background())

	if err != nil {
		return nil, err
	}

	rules := []*Rule{}
	for cur.Next(context.Background()) {
		rule := &Rule{}
		if err := cur.Decode(rule); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	return rules, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindByID(ruleID string) (*Rule, error) {
	_id, err := objectid.FromHex(ruleID)
	if err != nil {
		return nil, errors.ErrID
	}

	rule := &Rule{}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", _id))
	if err = d.collection.FindOne(context.Background(), filter).Decode(rule); err != nil {
		return nil, err
	}

	return rule, nil
}
