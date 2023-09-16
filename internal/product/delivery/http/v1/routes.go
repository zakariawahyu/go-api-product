package v1

func (p *productHandlers) MapRoutes() {
	p.group.POST("", p.Create)
	p.group.GET("/:id", p.GetByID)
}
