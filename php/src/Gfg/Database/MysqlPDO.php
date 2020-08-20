<?php
namespace Gfg\Database;

class MysqlPDO
{
    /**
     * @var \PDO
     */
    private $connection;

    public function __construct(
        string $username,
        string $password,
        string $host,
        string $port,
        string $database
    ) {
        $dsn = sprintf('mysql:host=%s;port:%s;dbname=%s', $host, $port, $database);
        $this->connection = new \PDO($dsn, $username, $password);

        $this->connection->exec('use ' . $database);
    }

    public function getConnection(): \PDO
    {
        return $this->connection;
    }
}
