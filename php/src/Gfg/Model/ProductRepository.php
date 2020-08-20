<?php
namespace Gfg\Model;

use Gfg\Database\MysqlPDO;

class ProductRepository
{
    /** @var MysqlPDO */
    private $mysqlPDO;

    public function __construct(MysqlPDO $mysqlPDO)
    {
        $this->mysqlPDO = $mysqlPDO;
    }

    /**
     * @return Product[]
     */
    public function getProducts(int $offset = 0, int $limit = 20): array
    {
        $query =
            'SELECT 
                product.*,
                seller.id_seller,
                seller.name as seller_name,
                seller.email as seller_email,
                seller.phone as seller_phone,
                seller.uuid as seller_uuid
            FROM product INNER JOIN seller ON fk_seller = id_seller
            ORDER BY product.id_product ASC LIMIT %d,%d';

        $query = sprintf($query, $offset, $limit);

        $statement = $this->mysqlPDO->getConnection()->prepare($query);
        $statement->execute();

        return $this->buildArray($statement->fetchAll());
    }

    function getByUuid(string $uuid): Product
    {
        $query =
            'SELECT 
                product.*,
                seller.id_seller,
                seller.name as seller_name,
                seller.email as seller_email,
                seller.phone as seller_phone,
                seller.uuid as seller_uuid
            FROM product INNER JOIN seller ON fk_seller = id_seller
            WHERE product.uuid = :uuid';

        $statement = $this->mysqlPDO->getConnection()->prepare($query);
        $statement->execute([':uuid' => $uuid]);

        $products = $this->buildArray($statement->fetchAll());

        if (empty($products)) {
            throw new \Exception(
                sprintf(
                    'Product with UUID %s not found',
                    $uuid
                )
            );
        }

        return $products[0];
    }

    public function add(Product $product)
    {
        $query =
            'INSERT INTO product (uuid, name, brand, stock, fk_seller)
              VALUES (:uuid, :name, :brand, :stock, :fk_seller)';

        $statement = $this->mysqlPDO->getConnection()->prepare($query);
        $success = $statement->execute(
            [
                ':uuid' => $product->getUuid(),
                ':name' => $product->getName(),
                ':brand' => $product->getBrand(),
                ':stock' => $product->getStock(),
                ':fk_seller' => $product->getSeller()->getSellerId(),
            ]
        );

        if (false === $success) {
            throw new \Exception(
                sprintf(
                    'Failed to add product %s to the repository',
                    $product->getUuid()
                )
            );
        }
    }

    public function update(Product $product)
    {
        $query =
            'UPDATE product SET 
                name = :name,
                brand = :brand,
                stock = :stock,
                fk_seller = :fk_seller
              WHERE id_product = :id_product';

        $statement = $this->mysqlPDO->getConnection()->prepare($query);

        $success = $statement->execute(
            [
                ':id_product' => $product->getProductId(),
                ':name' => $product->getName(),
                ':brand' => $product->getBrand(),
                ':stock' => $product->getStock(),
                ':fk_seller' => $product->getSeller()->getSellerId(),
            ]
        );



        if (false === $success) {
            throw new \Exception(
                sprintf(
                    'Failed to update product %s in the repository. Error info: \'%s\'',
                    $product->getUuid(),
                    $statement->errorInfo()
                )
            );
        }
    }

    public function deleteByUuid(string $uuid)
    {
        $query = 'DELETE FROM product WHERE uuid = :uuid';

        $statement = $this->mysqlPDO->getConnection()->prepare($query);
        $success = $statement->execute(
            [
                ':uuid' => $uuid,
            ]
        );

        if (false === $success) {
            throw new \Exception(
                sprintf(
                    'Failed to delete product %s from the repository',
                    $uuid
                )
            );
        }
    }

    /**
     * @param array $statementResults
     * @return Product[]
     */
    private function buildArray(array $statementResults): array
    {
        $objects = [];

        foreach ($statementResults as $result) {
            $objects[] = new Product(
                $result['uuid'],
                $result['name'],
                $result['brand'],
                $result['stock'],
                new Seller(
                    $result['id_seller'],
                    $result['seller_uuid'],
                    $result['seller_name'],
                    $result['seller_email'],
                    $result['seller_phone']
                ),
                $result['id_product']
            );
        }


        return $objects;
    }
}
