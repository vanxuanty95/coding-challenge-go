<?php
namespace Gfg\Api\Controller;

use Gfg\Model\Seller;
use Gfg\Model\SellerRepository;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class SellerController
{
    const PAGE_SIZE = 10;

    /** @var SellerRepository */
    private $sellerRepository;

    public function __construct(SellerRepository $sellerRepository)
    {
        $this->sellerRepository = $sellerRepository;
    }

    public function getList(Request $request, Response $response): Response
    {
        $serializable = [];

        $page = $request->getQueryParams()['page'] ?? 1;

        foreach ($this->sellerRepository->getSellers(
            ($page - 1) * static::PAGE_SIZE,
            static::PAGE_SIZE
        ) as $product) {
            $serializable[] = $this->toJsonSerializable($product);
        }

        $response = $response->withHeader('Content-Type', 'application/json');
        $response->getBody()->write(json_encode($serializable));

        return $response;
    }

    private function toJsonSerializable(Seller $seller): array
    {
        return [
            'id' => $seller->getUuid(),
            'name' => $seller->getName(),
            'email' => $seller->getEmail(),
            'phone' => $seller->getPhone(),
        ];
    }
}
