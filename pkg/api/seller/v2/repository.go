package v2

import "database/sql"

const (
	TOP_SELLER_QUERY = "SELECT id_seller, name, email, phone, uuid FROM seller " +
		"INNER JOIN (SELECT fk_seller, COUNT(*) FROM product GROUP BY fk_seller ORDER BY COUNT(*) DESC LIMIT ?) " +
		"AS products ON seller.id_seller = products.fk_seller;"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) getTopSellers(top int) ([]*Seller, error) {
	rows, err := r.db.Query(TOP_SELLER_QUERY, top)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var topSellers []*Seller
	for rows.Next() {
		seller := &Seller{}
		err = rows.Scan(&seller.SellerID, &seller.Name, &seller.Email, &seller.Phone, &seller.UUID)
		if err != nil {
			return nil, err
		}

		topSellers = append(topSellers, seller)
	}

	return topSellers, nil
}
