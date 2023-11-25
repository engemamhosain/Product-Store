package models

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Products struct {
	_ID           bson.ObjectId `bson:"_id,omitempty"`
	ID           int `bson:"id"`
	Name         string        `bson:"name"`
	Description  string        `bson:"description"`
	Specification string        `bson:"specification"`
	BrandId       int            `bson:"brand_id"`
	CategoryId    int            `bson:"category_id"`
	SupplierId    int            `bson:"supplier_id"`
	UnitPrice    int           `bson:"unit_price"`
	DiscountPrice int           `bson:"discount_price"`
	Tag          string        `bson:"tag"`
	StatusId     int           `bson:"status_id"`
	CreatedAt    time.Time     `bson:"created_at"`
}

// Save inserts a new product into the MongoDB collection
func (product *Products) Save() error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("products")
	product._ID = bson.NewObjectId()
	product.ID = getNextSequence(session, "product_id")
	return c.Insert(product)
}

// GetProducts retrieves all products from the MongoDB collection
func GetProducts(filters map[string]interface{}, limit, offset int) ([]bson.M, error) {
	session := GetSession()
	defer session.Close()

	// c := session.DB("product_store").C("products")

	productCollection := session.DB("product_store").C("products")
	// brandCollection := session.DB("product_store").C("brands")
	// categoryCollection := session.DB("product_store").C("categories")
	
	condition:=bson.M{
		"status_id":1,
	}

	  for key, value := range filters {
		condition[key] = value
	 }
	pipe := productCollection.Pipe([]bson.M{
		{
				"$lookup": bson.M{
						"from":         "brands",
						"localField":   "brand_id",
						"foreignField": "id",
						"as":           "brand_info",
				},
		},
		{
				"$lookup": bson.M{
						"from":         "categories",
						"localField":   "brand_info.id",
						"foreignField": "parent_id",
						"as":           "category_info",
				},
		},
		{
			"$sort": bson.M{
			  "products.unit_price": 1, // 1 for ascending, -1 for descending
			},
		  },
		{

				"$match":condition,
		},
		{
			"$skip":offset,
			
		},
	   {
		 "$limit":limit, 
	   },
		{
				"$project": bson.M{
				 "description": 1,
				  "discount_price": 1,
				  "id": 1,
				  "name": 1,
				  "specification": 1,
				  "status_id": 1,
				  "supplier_id": 1,
				  "tag": 1,
				  "unit_price": 1,
				  "brand_info":bson.M {
						 "id": 1,
						"name": 1,
						"categories": "$category_info",
				  },

				},
		},

})

	var results []bson.M
	if err := pipe.All(&results); err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	return results, nil
	// var products []Products
	// err := c.Find(bson.M{"status_id": 1}).All(&products)
	// if err != nil {
	// 	return nil, err
	// }

	// return products, nil
}

// GetProductByID retrieves a product by ID from the MongoDB collection
func GetProductByID(id int) (*Products, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("products")

	var product Products
	err := c.Find(bson.M{"id": id}).One(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProduct updates a product in the MongoDB collection
func UpdateProduct(product *Products) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("products")


	
	return c.Update(bson.M{"id": product.ID}, bson.M{"$set":bson.M{
		"name":             product.Name,
		"description":             product.Description,
		"specification":             product.Specification,
		"brandId":             product.BrandId,
		"category_id":             product.CategoryId,
		"supplier_id":             product.SupplierId,
		"unit_price":             product.UnitPrice,
		"discount_price":             product.DiscountPrice,
		"tag":             product.Tag,
		"status_id":             product.StatusId,

	
	}})
	
}

// DeleteProduct deletes a product from the MongoDB collection
func DeleteProduct(id int) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("products")

	return c.RemoveId(id)
}


func GetPaginatedFilteredProducts(filters map[string]interface{}, limit, offset int) ([]Products, error) {
	session := GetSession()
	defer session.Close()
 
	c := session.DB("product_store").C("products") // Replace "your_database_name" with your actual database name
 
	var products []Products
	query := bson.M{}
 
	// Add filters to the query
	for key, value := range filters {
	   query[key] = value
	}
 
	err := c.Find(query).Skip(offset).Limit(limit).All(&products)
	if err != nil {
	   return nil, err
	}
 
	return products, nil
 }
