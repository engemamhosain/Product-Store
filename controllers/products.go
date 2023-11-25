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

// ProductsController operations for Products
type ProductsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Products
// @Param	body		body 	dtos.ProductsPostDto	true		"body for Products content"
// @Success 201 {object} models.Products
// @Failure 403 body is empty
// @router / [post]
func (c *ProductsController) Post() {

	var productsPostDto dtos.ProductsPostDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &productsPostDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	
	product := models.Products{
		Name:        productsPostDto.Name,
		Description:  productsPostDto.Description,
		Specification: productsPostDto.Specification,
		BrandId:      productsPostDto.BrandId,
		CategoryId:   productsPostDto.CategoryId,
		SupplierId:   productsPostDto.SupplierId,
		UnitPrice:    productsPostDto.UnitPrice,
		DiscountPrice: productsPostDto.DiscountPrice,
		Tag:          productsPostDto.Tag,
		StatusId:     1,
		CreatedAt:time.Now(),
	}

	err = product.Save()
	if err != nil {
		logs.Error("Error saving product:", err)
		c.Data["json"] = map[string]string{"message": "Error saving product"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product added successfully"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Products by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Products
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductsController) GetOne() {


	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	product, err := models.GetProductByID(id)
	if err != nil {
		logs.Error("Error getting product:", err)
		c.Data["json"] = map[string]string{"message": "Error getting product"}
	} else {
		c.Data["json"] = product
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Products
// @Param	product_name	query	string	false	" QuantumStream Smart TV"
// @Param	brand_name	query	string	false	" QuantumStream Smart TV"
// @Param	category_name	query	string	false	" QuantumStream Smart TV"
// @Param	product_id	query	string	false	""
// @Param	limit	query	int	true	""
// @Param	offset	query	int	true	""
// @Success 200 {object} models.Products
// @Failure 403
// @router / [get]
func (c *ProductsController) GetAll() {
	filters := make(map[string]interface{})

	if product_name := c.GetString("product_name"); product_name != "" {
		filters["name"] = product_name
	}
	if brand_name := c.GetString("brand_name"); brand_name != "" {
		filters["brand_info.name"] = brand_name
	}

	if category_name := c.GetString("category_name"); category_name != "" {
		filters["category_info.name"] = category_name
	}

	if id, err := c.GetInt("product_id", 0); err == nil {
		if(id!=0){
			filters["id"] = id
		}
		
	 }
	
	 limit, _ := c.GetInt("limit", 10)
	 offset, _ := c.GetInt("offset", 0)
 
	products, err := models.GetProducts(filters,limit,offset)
	if err != nil {
		logs.Error("Error getting products:", err)
		c.Data["json"] = map[string]string{"message": "Error getting products"}
	} else {
		c.Data["json"] = products
	}
	c.ServeJSON()
}



// GetAll ...
// @Title GetAll
// @Description get Products
// @Param	id	query	int	false	"1"
// @Param	name	query	int	false	" QuantumStream Smart TV"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Products
// @Failure 403
// @router /filter [get]
func (c *ProductsController) FilterProducts() {
	filters := make(map[string]interface{})

	// Extract filters from query parameters
	if id, err := c.GetInt("id", 0); err == nil {
	   filters["id"] = id
	}
 
	if brandID, err := c.GetInt("brand_id", 0); err == nil {
	   filters["brand_id"] = brandID
	}
 
	// Add more filters as needed...
 
	// Extract limit and offset parameters for pagination
	limit, _ := c.GetInt("limit", 10)
	offset, _ := c.GetInt("offset", 0)
 
	// Call the method to retrieve paginated and filtered products
	products, err := models.GetPaginatedFilteredProducts(filters, limit, offset)
	if err != nil {
	   c.Data["json"] = map[string]string{"message": "Error getting products"}
	} else {
	   c.Data["json"] = products
	}
	c.ServeJSON()
}



// Put ...
// @Title Put
// @Description update the Products
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	dtos.ProductsPutDto	true		"body for Products content"
// @Success 200 {object} models.Products
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductsController) Put() {


	var productsPutDto dtos.ProductsPutDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &productsPutDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}


	product := models.Products{
		ID:productsPutDto.ID,
		Name:        productsPutDto.Name,
		Description:  productsPutDto.Description,
		Specification: productsPutDto.Specification,
		BrandId:      productsPutDto.BrandId,
		CategoryId:   productsPutDto.CategoryId,
		SupplierId:   productsPutDto.SupplierId,
		UnitPrice:    productsPutDto.UnitPrice,
		DiscountPrice: productsPutDto.DiscountPrice,
		Tag:          productsPutDto.Tag,
		StatusId:     productsPutDto.StatusId,
	}


	err = models.UpdateProduct(&product)
	if err != nil {
		logs.Error("Error updating product:", err)
		c.Data["json"] = map[string]string{"message": "Error updating product"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product updated successfully"}
	}
	c.ServeJSON()

}

// Delete ...
// @Title Delete
// @Description delete the Products
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}

	err = models.DeleteProduct(id)
	if err != nil {
		logs.Error("Error deleting product:", err)
		c.Data["json"] = map[string]string{"message": "Error deleting product"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product deleted successfully"}
	}
	c.ServeJSON()
}
