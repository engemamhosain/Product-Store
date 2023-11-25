package controllers

import (
	"strconv"
	"time"
	"encoding/json"
    "product_store/models"
	"github.com/beego/beego/v2/core/logs"
	"product_store/dtos"
	beego "github.com/beego/beego/v2/server/web"
)

// BrandsController operations for Brands
type BrandsController struct {
	beego.Controller
}

// URLMapping ...
func (c *BrandsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Brands
// @Param	body		body 	dtos.BrandsPostDto	true		"body for Brands content"
// @Success 201 {object} models.Brands
// @Failure 403 body is empty
// @router / [post]
func (c *BrandsController) Post() {

	var brandsPostDto dtos.BrandsPostDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &brandsPostDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}


	brand := models.Brands{
		Name: brandsPostDto.Name,       
		StatusId:      1,
		CreatedAt:      time.Now(),
	}
	err = brand.Save()
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Brand created successfully"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Brands by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Brands
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BrandsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	brand, err := models.GetBrandByID(id)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = brand
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Brands
// @Success 200 {object} models.Brands
// @Failure 403
// @router / [get]
func (c *BrandsController) GetAll() {
	brands, err := models.GetBrands()
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "error"}
	} else {
		c.Data["json"] = brands
	}
	c.ServeJSON()
}


// Put ...
// @Title Put
// @Description update the Brands
// @Param	body		body 	dtos.BrandsPutDto	true		"body for Brands content"
// @Success 200 {object} models.Brands
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BrandsController) Put() {

	var brandsPutDto dtos.BrandsPutDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &brandsPutDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	brand := models.Brands{
		ID:brandsPutDto.ID,
		Name: brandsPutDto.Name,       
		StatusId:      brandsPutDto.StatusId,
	}
	err = models.UpdateBrand(&brand)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "Brand updated error"}
	} else {
		c.Data["json"] = map[string]string{"message": "Brand updated successfully"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Brands
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BrandsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	err = models.DeleteBrand(id)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Brand deleted successfully"}
	}
	c.ServeJSON()
}
