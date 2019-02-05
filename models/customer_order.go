package models

import "gopkg.in/mgo.v2/bson"

// Represents a CustomerOrder
type CustomerOrder struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	CustomerId bson.ObjectId `bson:"customer_id" json:"customer_id"`
	Date       string        `bson:"date" json:"date"`
}
