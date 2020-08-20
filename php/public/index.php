<?php
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

require __DIR__ . '/../vendor/autoload.php';

$container = new \DI\Container();

AppFactory::setContainer($container);
$app = AppFactory::create();

$container->set('mysqlPDO', function () {
    return new \Gfg\Database\MysqlPDO('user', 'password', 'db', '3306', 'product');
});

$container->set('productRepository', function () use ($container) {
    return new \Gfg\Model\ProductRepository($container->get('mysqlPDO'));
});

$container->set('sellerRepository', function () use ($container) {
    return new \Gfg\Model\SellerRepository($container->get('mysqlPDO'));
});

$container->set('emailProvider', function () use ($container) {
    return new \Gfg\Mail\EmailProvider(__DIR__. '/../logs/emails.log');
});

$container->set('apiProductController', function () use ($container) {
    return new \Gfg\Api\Controller\ProductController(
        $container->get('productRepository'),
        $container->get('sellerRepository'),
        $container->get('emailProvider')
    );
});

$container->set('apiSellerController', function () use ($container) {
    return new \Gfg\Api\Controller\SellerController(
        $container->get('sellerRepository')
    );
});

$app->get('/api/v1/products', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\ProductController $controller */
    $controller = $container->get('apiProductController');
    return $controller->getList($request, $response);
});

$app->get('/api/v1/product', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\ProductController $controller */
    $controller = $container->get('apiProductController');
    return $controller->get($request, $response);
});

$app->post('/api/v1/product', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\ProductController $controller */
    $controller = $container->get('apiProductController');
    return $controller->post($request, $response);
});

$app->delete('/api/v1/product', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\ProductController $controller */
    $controller = $container->get('apiProductController');
    return $controller->delete($request, $response);
});

$app->put('/api/v1/product', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\ProductController $controller */
    $controller = $container->get('apiProductController');
    return $controller->put($request, $response);
});

$app->get('/api/v1/sellers', function (Request $request, Response $response, $args) use ($container) {
    /** @var \Gfg\Api\Controller\SellerController $controller */
    $controller = $container->get('apiSellerController');
    return $controller->getList($request, $response);
});

// Run app
$app->run();
