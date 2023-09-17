package v1

func (p *productHandlers) MapRoutes() {
	p.group.GET("", p.GetAllProduct)
	p.group.POST("", p.CreateProduct)
	p.group.GET("/:id", p.GetProductByID)
	p.group.PUT("/:id", p.UpdateProduct)
	p.group.DELETE("/:id", p.DeleteProduct)
	p.group.DELETE("/:id/hard-delete", p.HardDeleteProduct)
}
