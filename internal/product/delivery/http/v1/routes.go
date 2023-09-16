package v1

func (p *productHandlers) MapRoutes() {
	p.group.POST("", p.CreateProduct)
	p.group.GET("/:id", p.GetProductByID)
	p.group.PUT("/:id", p.UpdateProduct)
	p.group.DELETE("/:id", p.DeleteProduct)
}
