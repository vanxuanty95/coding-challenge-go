<?php
namespace Gfg\Model;

class Product
{
    /** @var int|null */
    private $productId;

    /** @var string */
    private $uuid;

    /** @var string */
    private $name;

    /** @var string */
    private $brand;

    /** @var int */
    private $stock;

    /** @var Seller */
    private $seller;

    public function __construct(
        string $uuid,
        string $name,
        string $brand,
        int $stock,
        Seller $seller,
        int $productId = null
    ) {
        $this->uuid = $uuid;
        $this->name = $name;
        $this->brand = $brand;
        $this->stock = $stock;
        $this->seller = $seller;
        $this->productId = $productId;
    }

    public function getProductId(): ?int
    {
        return $this->productId;
    }

    public function getUuid(): string
    {
        return $this->uuid;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getBrand(): string
    {
        return $this->brand;
    }

    public function getStock(): int
    {
        return $this->stock;
    }

    public function getSeller(): Seller
    {
        return $this->seller;
    }
}
