package api

import (
	config2 "coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api/helper"
	v1Product "coding-challenge-go/pkg/api/product/v1"
	v2Product "coding-challenge-go/pkg/api/product/v2"
	v1Seller "coding-challenge-go/pkg/api/seller/v1"
	v2Seller "coding-challenge-go/pkg/api/seller/v2"
	"database/sql"
	"github.com/gin-gonic/gin"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(cfg *config2.Config, db *sql.DB) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	CreateV1API(cfg, r, db)
	CreateV2API(cfg, r, db)
	return r, nil
}

func CreateV1API(cfg *config2.Config, r *gin.Engine, db *sql.DB) {
	apiV1 := r.Group("api/v1")

	productRepository := v1Product.NewRepository(db)
	sellerRepository := v1Seller.NewRepository(db)
	notifiersFactory := helper.NewNotifiersFactory(cfg)

	CreateProductV1API(apiV1, productRepository, sellerRepository, notifiersFactory)
	CreateSellerV1API(apiV1, sellerRepository)
}

func CreateV2API(cfg *config2.Config, r *gin.Engine, db *sql.DB) {
	apiV2 := r.Group("api/v2")
	productRepository := v2Product.NewRepository(db)
	sellerRepository := v2Seller.NewRepository(db)

	CreateProductV2API(cfg, apiV2, productRepository, sellerRepository)
	CreateSellerV2API(apiV2, sellerRepository)
}

func CreateProductV1API(routerGroup *gin.RouterGroup, productRepository *v1Product.Repository, sellerRepository *v1Seller.Repository, notifiersFactory *helper.NotifiersFactory) {
	productController := v1Product.NewController(productRepository, sellerRepository, notifiersFactory)
	routerGroup.GET("products", productController.List)
	routerGroup.GET("product", productController.Get)
	routerGroup.POST("product", productController.Post)
	routerGroup.PUT("product", productController.Put)
	routerGroup.DELETE("product", productController.Delete)
}

func CreateSellerV1API(routerGroup *gin.RouterGroup, sellerRepository *v1Seller.Repository) {
	sellerController := v1Seller.NewController(sellerRepository)
	routerGroup.GET("sellers", sellerController.List)
}

func CreateProductV2API(cfg *config2.Config, routerGroup *gin.RouterGroup, productRepository *v2Product.Repository, sellerRepository *v2Seller.Repository) {
	productController := v2Product.NewController(cfg, productRepository, sellerRepository)
	routerGroup.GET("products", productController.List)
	routerGroup.GET("product", productController.Get)
}

func CreateSellerV2API(routerGroup *gin.RouterGroup, sellerRepository *v2Seller.Repository) {
	sellerController := v2Seller.NewController(sellerRepository)
	routerGroup.GET("sellers/top:number", sellerController.TopSeller)
}
