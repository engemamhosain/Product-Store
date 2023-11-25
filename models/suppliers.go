package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Suppliers struct {
	_ID                bson.ObjectId `bson:"_id,omitempty"`
	ID                int `bson:"id"`
	Name              string        `bson:"name"`
	Email             string        `bson:"email"`
	Phone             string        `bson:"phone"`
	StatusId          int        `bson:"status_id"`
	IsVerifiedSupplier int           `bson:"is_verified_supplier"`
	CreatedAt         time.Time     `bson:"created_at"`
}

// Save inserts a new supplier into the MongoDB collection
func (supplier *Suppliers) Save() error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("suppliers")
	supplier._ID = bson.NewObjectId()
	supplier.ID = getNextSequence(session, "supplier_id")
	return c.Insert(supplier)
}

// GetSuppliers retrieves all suppliers from the MongoDB collection
func GetSuppliers() ([]Suppliers, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("suppliers")

	var suppliers []Suppliers
	err := c.Find(bson.M{"status_id": 1}).All(&suppliers)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

// GetSupplierByID retrieves a supplier by ID from the MongoDB collection
func GetSupplierByID(id int) (*Suppliers, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("suppliers")

	var supplier Suppliers
	err := c.Find(bson.M{"id": id}).One(&supplier)
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

// UpdateSupplier updates a supplier in the MongoDB collection
func UpdateSupplier(supplier *Suppliers) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("suppliers")

	return c.Update(bson.M{"id": supplier.ID}, bson.M{"$set":bson.M{
		"name":             supplier.Name,
		"email":            supplier.Email,
		"phone":            supplier.Phone,
		"status_id":          supplier.StatusId,
		"is_verified_supplier": supplier.IsVerifiedSupplier,
	}})
}

// DeleteSupplier deletes a supplier from the MongoDB collection
func DeleteSupplier(id int) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("suppliers")

	return c.RemoveId(id)
}

