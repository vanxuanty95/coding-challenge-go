<?php
namespace Gfg\Model;

class Seller
{
    /** @var int */
    private $sellerId;

    /** @var string */
    private $uuid;

    /** @var string */
    private $name;

    /** @var string */
    private $email;

    /** @var string */
    private $phone;

    public function __construct(
        int $sellerId,
        string $uuid,
        string $name,
        string $email,
        string $phone
    ) {
        $this->sellerId = $sellerId;
        $this->uuid = $uuid;
        $this->name = $name;
        $this->email = $email;
        $this->phone = $phone;
    }

    public function getSellerId(): int
    {
        return $this->sellerId;
    }

    public function getUuid(): string
    {
        return $this->uuid;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getEmail(): string
    {
        return $this->email;
    }

    public function getPhone(): string
    {
        return $this->phone;
    }
}
