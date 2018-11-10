package rules

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
	Insert(rule *Rule) (*Rule, error)
	Update(rule *Rule) (*Rule, error)
	FindByID(ruleID string) (*Rule, error)
	FindByArticleID(articleId string) (*Rule, error)
}

// New dao es interno a este modulo, nadie fuera del modulo tiene acceso
func GetDao() (Dao, error) {
	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("rules")

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

func (d daoStruct) Insert(rule *Rule) (*Rule, error) {
	fmt.Println(rule.ID)

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

	rule.Modified = time.Now()

	doc, err := db.EncodeDocument(rule)
	if err != nil {
		return nil, err
	}

	_, err = d.collection.UpdateOne(context.Background(),
		bson.NewDocument(doc.LookupElement("_id")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				doc.LookupElement("articleId"),
				doc.LookupElement("lowRate"),
				doc.LookupElement("highRate"),
				doc.LookupElement("created"),
				doc.LookupElement("modified"),
				doc.LookupElement("enabled"),
			),
		))

	if err != nil {
		return nil, err
	}

	return rule, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindByID(ruleID string) (*Rule, error) {
	_id, err := objectid.FromHex(ruleID)
	if err != nil {
		return nil, err
	}

	rule := &Rule{}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", _id))
	if err := d.collection.FindOne(context.Background(), filter).Decode(rule); err != nil {
		return nil, err
	}

	return rule, nil
}

// FindByID lee un usuario desde la db
func (d daoStruct) FindByArticleID(articleId string) (*Rule, error) {
	rule := &Rule{}
	filter := bson.NewDocument(bson.EC.String("articleId", articleId))
	if err := d.collection.FindOne(context.Background(), filter).Decode(rule); err != nil {
		return nil, err
	}

	return rule, nil
}
