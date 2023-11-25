package controllers

import (
	"strconv"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"time"
	"product_store/models"
	"product_store/dtos"
	beego "github.com/beego/beego/v2/server/web"
)

// CategoriesController operations for Categories
type CategoriesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CategoriesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Categories
// @Param	body		body 	dtos.CategoriesPostDto	true		"body for Categories content"
// @Success 201 {object} models.Categories
// @Failure 403 body is empty
// @router / [post]
func (c *CategoriesController) Post() {

	var categoriesPostDto dtos.CategoriesPostDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &categoriesPostDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	category := models.Categories{
		Name:      categoriesPostDto.Name,
		Sequence:  categoriesPostDto.Sequence,
		ParentId:  categoriesPostDto.ParentId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}

	err = category.Save()
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Category created successfully"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Categories by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Categories
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoriesController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}

	category, err := models.GetCategoryByID(id)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = category
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Categories
// @Success 200 {object} models.Categories
// @Failure 403
// @router / [get]
func (c *CategoriesController) GetAll() {
	categories, err := models.GetCategories()
	if err != nil {

		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = categories
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Categories
// @Param	body		body 	models.CategoriesPutDto	true		"body for Categories content"
// @Success 200 {object} models.Categories
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CategoriesController) Put() {

	var categoriesPutDto dtos.CategoriesPutDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &categoriesPutDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	

	category := models.Categories{
		ID:        categoriesPutDto.ID,
		Name:      categoriesPutDto.Name,
		Sequence:  categoriesPutDto.Sequence,
		ParentId:  categoriesPutDto.ParentId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}

	err = models.UpdateCategory(&category)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Category updated successfully"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Categories
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoriesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	err = models.DeleteCategory(id)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Category deleted successfully"}
	}
	c.ServeJSON()
}
