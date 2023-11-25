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
// SuppliersController operations for Suppliers
type SuppliersController struct {
	beego.Controller
}

// URLMapping ...
func (c *SuppliersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Suppliers
// @Param	body		body 	dtos.SuppliersPostDto	true		"body for Suppliers content"
// @Success 201 {object} models.Suppliers
// @Failure 403 body is empty
// @router / [post]
func (c *SuppliersController) Post() {


	var suppliersPostDto dtos.SuppliersPostDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &suppliersPostDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	
	supplier := models.Suppliers{

		Name:             suppliersPostDto.Name,
		Email:            suppliersPostDto.Email,
		Phone:            suppliersPostDto.Phone,
		StatusId:          1,
		IsVerifiedSupplier: 1,
		CreatedAt:         time.Now(),
	}

	err = supplier.Save()
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Supplier created successfully"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Suppliers by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Suppliers
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SuppliersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	supplier, err := models.GetSupplierByID(id)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = supplier
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Suppliers
// @Success 200 {object} models.Suppliers
// @Failure 403
// @router / [get]
func (c *SuppliersController) GetAll() {

	suppliers, err := models.GetSuppliers()
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = suppliers
	}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Suppliers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	dtos.SuppliersPutDto	true		"body for Suppliers content"
// @Success 200 {object} models.Suppliers
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SuppliersController) Put() {


	var suppliersPutDto dtos.SuppliersPutDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &suppliersPutDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}
	
	supplier := models.Suppliers{
		ID:suppliersPutDto.ID,
		Name:             suppliersPutDto.Name,
		Email:            suppliersPutDto.Email,
		Phone:            suppliersPutDto.Phone,
		StatusId:          suppliersPutDto.StatusId,
		IsVerifiedSupplier: suppliersPutDto.IsVerifiedSupplier,
	}

	err = models.UpdateSupplier(&supplier)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Supplier updated successfully"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Suppliers
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SuppliersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}

	err = models.DeleteSupplier(id)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
	} else {
		c.Data["json"] = map[string]string{"message": "Supplier deleted successfully"}
	}
	c.ServeJSON()
}
