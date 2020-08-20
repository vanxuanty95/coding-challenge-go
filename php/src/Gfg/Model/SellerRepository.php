<?php
namespace Gfg\Model;

use Gfg\Database\MysqlPDO;

class SellerRepository
{
    /** @var MysqlPDO */
    private $mysqlPDO;

    public function __construct(MysqlPDO $mysqlPDO)
    {

        $this->mysqlPDO = $mysqlPDO;
    }

    public function getSellers(int $offset = 0, int $limit = 20): array
    {
        $query = 'SELECT * FROM seller ORDER BY id_seller ASC LIMIT %d,%d';

        $query = sprintf($query, $offset, $limit);

        $statement = $this->mysqlPDO->getConnection()->prepare($query);
        $statement->execute();

        return $this->buildArray($statement->fetchAll());

    }

    function getByUuid(string $uuid): Seller
    {
        $queryPattern = 'SELECT * FROM seller WHERE uuid = :uuid';

        $statement = $this->mysqlPDO->getConnection()->prepare($queryPattern);
        $statement->execute([':uuid' => $uuid]);

        $record = $statement->fetch();

        if (false === $record) {
            throw new \Exception(
                sprintf(
                    'Seller with UUID %s not found',
                    $uuid
                )
            );
        }

        return new Seller(
            $record['id_seller'],
            $record['uuid'],
            $record['name'],
            $record['email'],
            $record['phone']
        );
    }

    private function buildArray(array $statementResults)
    {
        $objects = [];

        foreach ($statementResults as $result) {
            $objects[] = new Seller(
                $result['id_seller'],
                $result['uuid'],
                $result['name'],
                $result['email'],
                $result['phone']
            );
        }

        return $objects;
    }
}
