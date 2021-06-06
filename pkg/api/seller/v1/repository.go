package v1

import "database/sql"

const (
	FIND_BY_UUID_QUERY = "SELECT id_seller, name, email, phone, uuid FROM seller WHERE uuid = ?"
	LIST_QUERY         = "SELECT id_seller, name, email, phone, uuid FROM seller"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) FindByUUID(uuid string) (*Seller, error) {
	rows, err := r.db.Query(FIND_BY_UUID_QUERY, uuid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	seller := &Seller{}

	err = rows.Scan(&seller.SellerID, &seller.Name, &seller.Email, &seller.Phone, &seller.UUID)

	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *Repository) list() ([]*Seller, error) {
	rows, err := r.db.Query(LIST_QUERY)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sellers []*Seller

	for rows.Next() {
		seller := &Seller{}

		err := rows.Scan(&seller.SellerID, &seller.Name, &seller.Email, &seller.Phone, &seller.UUID)
		if err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}
