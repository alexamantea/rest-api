package dao

import (
	. "github.com/alexamantea/go/rest-api/models"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type APIDao struct {
	Server   string
	Database string
}

var db *mgo.Database

//Collections
const (
	CustomerCollection      = "customer"
	ProductCollection       = "product"
	CustomerOrderCollection = "customer_orders"
)

// Establish a connection to database
func (m *APIDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAllCustomers - Find list of customers
func (m *APIDao) FindAllCustomers() ([]Customer, error) {
	var customers []Customer
	err := db.C(CustomerCollection).Find(bson.M{}).All(&customers)
	return customers, err
}

// FindAllProducts - Find list of products
func (m *APIDao) FindAllProducts() ([]Product, error) {
	var products []Product
	err := db.C(ProductCollection).Find(bson.M{}).All(&products)
	return products, err
}

// FindAllCustomerOrders - Find list of orders
func (m *APIDao) FindAllCustomerOrders() ([]CustomerOrder, error) {
	var orders []CustomerOrder
	err := db.C(CustomerOrderCollection).Find(bson.M{}).All(&orders)
	return orders, err
}
