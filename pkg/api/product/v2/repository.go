package v2

import (
	"database/sql"
)

const (
	LIST_QUERY = "SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p " +
		"INNER JOIN seller s ON(s.id_seller = p.fk_seller) LIMIT ? OFFSET ?"
	FIND_BY_UUID_QUERY = "SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p " +
		"INNER JOIN seller s ON(s.id_seller = p.fk_seller) WHERE p.uuid = ?"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) list(offset int, limit int) ([]*product, error) {
	rows, err := r.db.Query(LIST_QUERY, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*product

	for rows.Next() {
		product := &product{}

		err = rows.Scan(&product.ProductID, &product.Name, &product.Brand, &product.Stock, &product.SellerUUID, &product.UUID)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *Repository) findByUUID(uuid string) (*product, error) {
	rows, err := r.db.Query(FIND_BY_UUID_QUERY, uuid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	product := &product{}

	err = rows.Scan(&product.ProductID, &product.Name, &product.Brand, &product.Stock, &product.SellerUUID, &product.UUID)

	if err != nil {
		return nil, err
	}

	return product, nil
}
