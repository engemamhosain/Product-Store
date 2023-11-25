package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Brands struct {
	_ID        bson.ObjectId `bson:"_id,omitempty"`
	ID       int `bson:"id"`
	Name      string        `bson:"name"`
	StatusId  int        `bson:"status_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

// Save inserts a new brand into the MongoDB collection
func (brand *Brands) Save() error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("brands")
	brand._ID = bson.NewObjectId()
	brand.ID = getNextSequence(session, "brand_id")
	return c.Insert(brand)
}

// GetBrands retrieves all brands from the MongoDB collection
func GetBrands() ([]Brands, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("brands")

	var brands []Brands
	err := c.Find(bson.M{"status_id": 1}).All(&brands)
	if err != nil {
		return nil, err
	}

	return brands, nil
}

// GetBrandByID retrieves a brand by ID from the MongoDB collection
func GetBrandByID(id int) (*Brands, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("brands")

	var brand Brands
	err := c.Find(bson.M{"id": id}).One(&brand)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

// UpdateBrand updates a brand in the MongoDB collection
func UpdateBrand(brand *Brands) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("brands")

	return c.Update(bson.M{"id": brand.ID}, bson.M{"$set":bson.M{
		"name" : brand.Name,
		"status_id" :brand.StatusId,
	} })
}

// DeleteBrand deletes a brand from the MongoDB collection
func DeleteBrand(id int) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("brands")

	return c.RemoveId(id)
}

func GetPaginatedFilteredBrands(filters map[string]interface{}, limit, offset int) ([]Brands, error) {
	session := GetSession()
	defer session.Close()
 
	c := session.DB("product_store").C("brands") // Replace "your_database_name" with your actual database name
 
	var brands []Brands
	query := bson.M{}
 
	// Add filters to the query
	for key, value := range filters {
	   query[key] = value
	}
 
	err := c.Find(query).Skip(offset).Limit(limit).All(&brands)
	if err != nil {
	   return nil, err
	}
 
	return brands, nil
 }
