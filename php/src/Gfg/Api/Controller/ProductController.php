<?php
namespace Gfg\Api\Controller;

use Gfg\Mail\EmailProvider;
use Gfg\Model\Product;
use Gfg\Model\ProductRepository;
use Gfg\Model\SellerRepository;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Ramsey\Uuid\Uuid;

class ProductController
{
    const PAGE_SIZE = 10;
    const PRODUCT_ID_PARAM = 'id';

    /** @var ProductRepository */
    private $productRepository;

    /** @var SellerRepository */
    private $sellerRepository;

    /** @var EmailProvider */
    private $emailProvider;

    public function __construct(
        ProductRepository $productRepository,
        SellerRepository $sellerRepository,
        EmailProvider $emailProvider
    ) {
        $this->productRepository = $productRepository;
        $this->sellerRepository = $sellerRepository;
        $this->emailProvider = $emailProvider;
    }

    public function getList(Request $request, Response $response): Response
    {
        $serializable = [];

        $page = $request->getQueryParams()['page'] ?? 1;

        foreach ($this->productRepository->getProducts(
            ($page - 1) * static::PAGE_SIZE,
            static::PAGE_SIZE
        ) as $product) {
            $serializable[] = $this->toJsonSerializable($product);
        }

        $response = $response->withHeader('Content-Type', 'application/json');
        $response->getBody()->write(json_encode($serializable));

        return $response;
    }

    public function get(Request $request, Response $response): Response
    {
        $uuid = $request->getQueryParams()[self::PRODUCT_ID_PARAM] ?? null;
        $product = $this->productRepository->getByUuid($uuid);

        $response = $response->withHeader('Content-Type', 'application/json');
        $response->getBody()->write(json_encode($this->toJsonSerializable($product)));

        return $response;
    }

    public function post(Request $request, Response $response): Response
    {
        $requestJson = json_decode($request->getBody()->getContents(), true);

        $seller = $this->sellerRepository->getByUuid($requestJson['seller']);

        $uuid = Uuid::uuid1(1)->toString();

        $product = new Product(
            $uuid,
            $requestJson['name'],
            $requestJson['brand'],
            $requestJson['stock'],
            $seller
        );

        $this->productRepository->add($product);

        $savedProduct = $this->productRepository->getByUuid($uuid);

        $serializable = $this->toJsonSerializable($savedProduct);

        $response = $response->withHeader('Content-Type', 'application/json');
        $response->getBody()->write(json_encode($serializable));

        return $response;
    }

    public function put(Request $request, Response $response): Response
    {
        $uuid = $request->getQueryParams()[self::PRODUCT_ID_PARAM] ?? null;
        $original = $this->productRepository->getByUuid($uuid);

        $requestJson = json_decode($request->getBody()->getContents(), true);

        $modified = new Product(
            $original->getUuid(),
            $requestJson['name'],
            $requestJson['brand'],
            $requestJson['stock'],
            $original->getSeller(),
            $original->getProductId()
        );

        $this->productRepository->update($modified);

        $updated = $this->productRepository->getByUuid($uuid);

        if ($original->getStock() !== $updated->getStock()) {
            $this->emailProvider->sendStockChangedEmail(
                $updated->getName(),
                $original->getStock(),
                $updated->getStock(),
                $updated->getSeller()->getEmail()
            );
        }

        $response = $response->withHeader('Content-Type', 'application/json');
        $response->getBody()->write(json_encode($this->toJsonSerializable($updated)));

        return $response;
    }

    public function delete(Request $request, Response $response): Response
    {
        $uuid = $request->getQueryParams()[self::PRODUCT_ID_PARAM] ?? null;
        $this->productRepository->deleteByUuid($uuid);
        return $response;
    }

    private function toJsonSerializable(Product $product): array
    {
        return [
            'uuid' => $product->getUuid(),
            'name' => $product->getName(),
            'brand' => $product->getBrand(),
            'stock' => $product->getStock(),
            'sellerUuid' => $product->getSeller()->getUuid()
        ];
    }
}
