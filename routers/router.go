// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"product_store/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/category",
                        beego.NSInclude(
                                &controllers.CategoriesController{},
                        ),
                ),
 		beego.NSNamespace("/brand",
                        beego.NSInclude(
                                &controllers.BrandsController{},
                        ),
                ),
 		beego.NSNamespace("/supplier",
                        beego.NSInclude(
                                &controllers.SuppliersController{},
                        ),
                ),
		beego.NSNamespace("/product_stocks",
                        beego.NSInclude(
                                &controllers.Product_stocksController{},
                        ),
                ),

                beego.NSNamespace("/product",
                        beego.NSInclude(
                                &controllers.ProductsController{},
                        ),
                ),
	)
	beego.AddNamespace(ns)
}
