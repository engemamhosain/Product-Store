package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Categories struct {
	_ID        bson.ObjectId `bson:"_id,omitempty"`
	ID        int `bson:"id"`
	Name      string        `bson:"name"`
	Sequence  int           `bson:"sequence"`
	ParentId  int `bson:"parent_id"`
	StatusId  int           `bson:"status_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

// Save inserts a new category into the MongoDB collection
func (category *Categories) Save() error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("categories")
	category._ID = bson.NewObjectId()
	category.ID = getNextSequence(session, "category_id")
	return c.Insert(category)
}

// GetCategories retrieves all categories from the MongoDB collection
func GetCategories() ([]Categories, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("categories")

	var categories []Categories
	err := c.Find(bson.M{"status_id": 1}).All(&categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByID retrieves a category by ID from the MongoDB collection
func GetCategoryByID(id int) (*Categories, error) {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("categories")

	var category Categories
	err := c.Find(bson.M{"id": id}).One(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// UpdateCategory updates a category in the MongoDB collection
func UpdateCategory(category *Categories) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("categories")

	return c.Update(bson.M{"id": category.ID}, bson.M{"$set":bson.M{
		"name":category.Name,
		"sequence":category.Sequence,
		"parent_id":category.ParentId,
		"statusId":category.StatusId,

	}})
}

// DeleteCategory deletes a category from the MongoDB collection
func DeleteCategory(id int) error {
	session := GetSession()
	defer session.Close()

	c := session.DB("product_store").C("categories")

	return c.RemoveId(id)
}

