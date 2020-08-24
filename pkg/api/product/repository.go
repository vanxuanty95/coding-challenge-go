package product

func NewRepository() *repository {
	return &repository{}
}

type repository struct {

}

func (r *repository) delete(product *product) {
	// FIXME XXX: implement me
}

func (r *repository) insert(product *product) {
	// FIXME XXX: implement me
}

func (r *repository) update(product *product) {
	// FIXME XXX: implement me
}

func (r *repository) list(offset int, limit int) []*product {
	// FIXME XXX: implement me
	return []*product{
		{
			ProductID: 0,
			UUID:      "baobab",
			Name:      "",
			Brand:     "",
			Stock:     0,
			SellerUUID: "dddd",
		},
	}
}

func (r *repository) findByUUID(uuid string) *product {
	// FIXME XXX: implement me
	return &product{
		ProductID: 0,
		UUID:      "baobab",
		Name:      "",
		Brand:     "",
		Stock:     0,
		SellerUUID:    "dddd",
	}
}