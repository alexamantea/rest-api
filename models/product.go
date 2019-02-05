package models

import "gopkg.in/mgo.v2/bson"

// Represents a Product
type Product struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Manufacturer string        `bson:"manufacturer" json:"manufacturer"`
	Price        float32       `bson:"price" json:"price"`
}
