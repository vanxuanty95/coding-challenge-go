package seller

func NewRepository() *Repository {
	return &Repository{}
}

type Repository struct {

}

func (r *Repository) FindByUUID(uuid string) *Seller {
	// FIXME XXX: implement me
	return &Seller{
		SellerID: 999,
		UUID:     "asdsad",
		Name:     "fdsf",
		Email:    "ewfw",
		Phone:    "fwefds",
	}
}

func (r *Repository) list() []*Seller {
	// FIXME XXX: implement me
	return []*Seller{
		{
			SellerID: 999,
			UUID:     "asdsad",
			Name:     "fdsf",
			Email:    "ewfw",
			Phone:    "fwefds",
		},
	}
}

