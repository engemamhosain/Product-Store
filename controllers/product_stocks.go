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
// Product_stocksController operations for Product_stocks
type Product_stocksController struct {
	beego.Controller
}

// URLMapping ...
func (c *Product_stocksController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Product_stocks
// @Param	body		body 	dtos.ProductStocksPostDto	true		"body for Product_stocks content"
// @Success 201 {object} models.ProductStocks
// @Failure 403 body is empty
// @router / [post]
func (c *Product_stocksController) Post() {

	var productStocksPostDto dtos.ProductStocksPostDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &productStocksPostDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}

	productStocks:= models.ProductStocks{
		ProductId:productStocksPostDto.ProductId,
		StockQuantity:productStocksPostDto.StockQuantity,
		UpdatedAt:time.Now(),
	}

	err = productStocks.Save()
	if err != nil {
		logs.Error("Error saving product stock:", err)
		c.Data["json"] = map[string]string{"message": "Error saving product stock"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product stock added successfully"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Product_stocks by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProductStocks
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Product_stocksController) GetOne() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	productStock, err := models.GetProductStockByID(id)
	if err != nil {
		logs.Error("Error getting product stock:", err)
		c.Data["json"] = map[string]string{"message": "Error getting product stock"}
	} else {
		c.Data["json"] = productStock
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Product_stocks
// @Success 200 {object} models.ProductStocks
// @Failure 403
// @router / [get]
func (c *Product_stocksController) GetAll() {
	productStocks, err := models.GetProductStocks()
	if err != nil {
		logs.Error("Error getting product stocks:", err)
		c.Data["json"] = map[string]string{"message": "Error getting product stocks"}
	} else {
		c.Data["json"] = productStocks
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Product_stocks
// @Param	body		body 	dtos.ProductStocksPutDto	true		"body for Product_stocks content"
// @Success 200 {object} models.ProductStocks
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Product_stocksController) Put() {
	
	var productStocksPutDto dtos.ProductStocksPutDto
	err := json.Unmarshal( c.Ctx.Input.RequestBody, &productStocksPutDto)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "payload not valid"}
		c.ServeJSON()
		return
	}


	productStocks:=models.ProductStocks{
		ID:productStocksPutDto.ID,
		ProductId:productStocksPutDto.ProductId,
		StockQuantity:productStocksPutDto.StockQuantity,
	}

	err = models.UpdateProductStock(&productStocks)
	if err != nil {
		logs.Error("Error updating product stock:", err)
		c.Data["json"] = map[string]string{"message": "Error updating product stock"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product stock updated successfully"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Product_stocks
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Product_stocksController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("Error:", err)
		c.Data["json"] = map[string]string{"error": "param not valid"}
		return
	}
	err = models.DeleteProductStock(id)
	if err != nil {
		logs.Error("Error deleting product stock:", err)
		c.Data["json"] = map[string]string{"message": "Error deleting product stock"}
	} else {
		c.Data["json"] = map[string]string{"message": "Product stock deleted successfully"}
	}
	c.ServeJSON()
}
