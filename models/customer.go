package models

import "gopkg.in/mgo.v2/bson"

// Represents a Customer
type Customer struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	CNPJ    string        `bson:"cnpj" json:"cnpj"`
	Name    string        `bson:"name" json:"name"`
	Address string        `bson:"address" json:"address"`
	Email   string        `bson:"email" json:"email"`
	Phone   string        `bson:"phone" json:"phone"`
}