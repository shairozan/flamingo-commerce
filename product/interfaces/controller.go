package interfaces

import (
	"flamingo/core/product/domain"
	"flamingo/framework/web"
	"flamingo/framework/web/responder"
	"net/url"
)

type (
	// ViewController demonstrates a product view controller
	ViewController struct {
		*responder.ErrorAware    `inject:""`
		*responder.RenderAware   `inject:""`
		*responder.RedirectAware `inject:""`
		domain.ProductService    `inject:""`
	}

	// ViewData is used for product rendering
	ViewData struct {
		Product *domain.Product
	}
)

// Get Response for Product matching sku param
func (vc *ViewController) Get(c web.Context) web.Response {
	product, err := vc.ProductService.Get(c, c.Param1("uid"))

	// catch error
	if err != nil {
		return vc.Error(c, err)
	}

	// normalize URL
	if url.QueryEscape(product.InternalName) != c.Param1("name") {
		return vc.Redirect("product.view", "uid", c.Param1("uid"), "name", url.QueryEscape(product.InternalName))
	}

	// render page
	return vc.Render(c, "pages/product/configurable", ViewData{Product: product})
}
