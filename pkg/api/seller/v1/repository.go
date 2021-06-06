package v1

import "database/sql"

const (
	FIND_BY_UUID_QUERY = "SELECT id_seller, name, email, phone, uuid FROM seller WHERE uuid = ?"
	LIST_QUERY         = "SELECT id_seller, name, email, phone, uuid FROM seller"
)

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

type RepositoryImpl struct {
	db *sql.DB
}

//go:generate mockgen -package v1 -destination repository_mock.go coding-challenge-go/pkg/api/seller/v1 Repository
type Repository interface {
	FindByUUID(uuid string) (*Seller, error)
	list() ([]*Seller, error)
}

func (r *RepositoryImpl) FindByUUID(uuid string) (*Seller, error) {
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

func (r *RepositoryImpl) list() ([]*Seller, error) {
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
