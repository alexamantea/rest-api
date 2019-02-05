package main

import (
	"encoding/json"
	"fmt"
	. "github.com/alexamantea/go/rest-api/config"
	. "github.com/alexamantea/go/rest-api/dao"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

var config = Config{}
var dao = APIDao{}

// AllCustomersEndpoint Get all customers
func AllCustomersEndpoint(w http.ResponseWriter, r *http.Request) {
	customers, err := dao.FindAllCustomers()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, customers)
}

// FindCustomerEndpoint Find one customer
func FindCustomerEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//CreateCustomerEndPoint Creates a customer
func CreateCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//UpdateCustomerEndPoint Updates a customer
func UpdateCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteCustomerEndPoint Remove a customer
func DeleteCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//AllProductsEndpoint Lists all products
func AllProductsEndpoint(w http.ResponseWriter, r *http.Request) {
	products, err := dao.FindAllProducts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, products)
}

//FindProductEndpoint Find a product by ID
func FindProductEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//CreateProductEndPoint Create a new product
func CreateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//UpdateProductEndPoint Update a product
func UpdateProductEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//DeleteProductEndPoint Removes a product
func DeleteProductEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//AllCustomerOrdersEndpoint Lists all orders
func AllCustomerOrdersEndpoint(w http.ResponseWriter, r *http.Request) {
	orders, err := dao.FindAllCustomerOrders()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, orders)
}

//FindOrdersEndpoint - Find order by customer ID
func FindOrdersEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customers", AllCustomersEndpoint).Methods("GET")
	r.HandleFunc("/customers", CreateCustomerEndPoint).Methods("POST")
	r.HandleFunc("/customers", UpdateCustomerEndPoint).Methods("PUT")
	r.HandleFunc("/customers", DeleteCustomerEndPoint).Methods("DELETE")
	r.HandleFunc("/customers/{id}", FindCustomerEndpoint).Methods("GET")

	r.HandleFunc("/products", AllProductsEndpoint).Methods("GET")
	r.HandleFunc("/products", CreateProductEndPoint).Methods("POST")
	r.HandleFunc("/products", UpdateProductEndPoint).Methods("PUT")
	r.HandleFunc("/products", DeleteProductEndPoint).Methods("DELETE")
	r.HandleFunc("/products/{id}", FindProductEndpoint).Methods("GET")

	r.HandleFunc("/orders", AllCustomerOrdersEndpoint).Methods("GET")
	r.HandleFunc("/orders/{id}", FindOrdersEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
