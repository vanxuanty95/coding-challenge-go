package v1

import (
	"database/sql"
)

const (
	DELETE_QUERY = "DELETE FROM product WHERE uuid = ?"
	INSERT_QUERY = "INSERT INTO product (id_product, name, brand, stock, fk_seller, uuid) VALUES(?,?,?,?,(SELECT id_seller FROM seller WHERE uuid = ?),?)"
	UPDATE_QUERY = "UPDATE product SET name = ?, brand = ?, stock = ? WHERE uuid = ?"
	LIST_QUERY   = "SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p " +
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

//go:generate mockgen -package v1 -destination repository_mock.go coding-challenge-go/pkg/api/product/v1 Repository
type Repository interface {
	delete(product *product) error
	insert(product *product) error
	update(product *product) error
	list(offset int, limit int) ([]*product, error)
	findByUUID(uuid string) (*product, error)
}

func (r *RepositoryImpl) delete(product *product) error {
	rows, err := r.db.Query(DELETE_QUERY, product.UUID)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *RepositoryImpl) insert(product *product) error {
	rows, err := r.db.Query(
		INSERT_QUERY,
		product.ProductID, product.Name, product.Brand, product.Stock, product.SellerUUID, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *RepositoryImpl) update(product *product) error {
	rows, err := r.db.Query(
		UPDATE_QUERY,
		product.Name, product.Brand, product.Stock, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *RepositoryImpl) list(offset int, limit int) ([]*product, error) {
	rows, err := r.db.Query(
		LIST_QUERY,
		limit, offset,
	)

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
	rows, err := r.db.Query(
		FIND_BY_UUID_QUERY,
		uuid,
	)

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
