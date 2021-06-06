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

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

type RepositoryImpl struct {
	db *sql.DB
}

//go:generate mockgen -package v2 -destination repository_mock.go coding-challenge-go/pkg/api/product/v2 Repository
type Repository interface {
	list(offset int, limit int) ([]*product, error)
	findByUUID(uuid string) (*product, error)
}

func (r *RepositoryImpl) list(offset int, limit int) ([]*product, error) {
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

func (r *RepositoryImpl) findByUUID(uuid string) (*product, error) {
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
