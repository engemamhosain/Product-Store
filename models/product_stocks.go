package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ProductStocks struct {
	_ID            bson.ObjectId `bson:"_id,omitempty"`
	ID            int `bson:"id"`
	ProductId     int           `bson:"product_id"`
	StockQuantity int           `bson:"stock_quantity"`
	UpdatedAt     time.Time     `bson:"updated_at"`
}

// Save inserts a new product stock into the MongoDB collection
func (productStock *ProductStocks) Save() error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("product_stocks")
	productStock._ID = bson.NewObjectId()
	productStock.ID = getNextSequence(session, "productStock_id")
	return c.Insert(productStock)
}

// GetProductStocks retrieves all product stocks from the MongoDB collection
func GetProductStocks() ([]ProductStocks, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("product_stocks")

	var productStocks []ProductStocks
	err := c.Find(nil).All(&productStocks)
	if err != nil {
		return nil, err
	}

	return productStocks, nil
}

// GetProductStockByID retrieves a product stock by ID from the MongoDB collection
func GetProductStockByID(id int) (*ProductStocks, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("product_stocks")

	var productStock ProductStocks
	err := c.Find(bson.M{"id": id}).One(&productStock)
	if err != nil {
		return nil, err
	}

	return &productStock, nil
}

// UpdateProductStock updates a product stock in the MongoDB collection
func UpdateProductStock(productStock *ProductStocks) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("product_stocks")

	return c.Update(bson.M{"id": productStock.ID}, bson.M{"$set":bson.M{
		"product_id":productStock.ProductId,
		"stock_quantity	":productStock.StockQuantity,
	}})
}

// DeleteProductStock deletes a product stock from the MongoDB collection
func DeleteProductStock(id int) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("product_stocks")

	return c.RemoveId(id)
}

