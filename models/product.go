package models

import "gopkg.in/mgo.v2/bson"

// Represents a Consumer, we use the bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Product struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Manufacturer string        `bson:"manufacturer" json:"manufacturer"`
	Price        float32       `bson:"price" json:"price"`
}

type Customer struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	CNPJ    string        `bson:"cnpj" json:"cnpj"`
	Name    string        `bson:"name" json:"name"`
	Address string        `bson:"address" json:"address"`
	Email   string        `bson:"email" json:"email"`
	Phone   string        `bson:"phone" json:"phone"`
}

type CustomerOrder struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	CustomerId bson.ObjectId `bson:"customer_id" json:"customer_id"`
	Date       string        `bson:"date" json:"date"`
}
