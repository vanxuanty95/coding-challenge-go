package product

import (
	"coding-challenge-go/pkg/api/seller"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/google/uuid"
)

const (
	LIST_PAGE_SIZE = 10
)

func NewController(repository *repository, sellerRepository *seller.Repository, sellerEmailProvider *seller.EmailProvider) *controller {
	return &controller{
		repository: repository,
		sellerRepository: sellerRepository,
		sellerEmailProvider: sellerEmailProvider,
	}
}

type controller struct {
	repository *repository
	sellerRepository *seller.Repository
	sellerEmailProvider *seller.EmailProvider
}

func (pc *controller) List(c *gin.Context) {
	request := &struct {
		Page int `form:"page,default=1"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products := pc.repository.list((request.Page - 1) * LIST_PAGE_SIZE, LIST_PAGE_SIZE)
	productsJson, err := json.Marshal(products)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", productsJson)
}

func (pc *controller) Get(c *gin.Context) {
	request := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := pc.repository.findByUUID(request.UUID)
	productJson, err := json.Marshal(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", productJson)
}

func (pc *controller) Post(c *gin.Context) {
	request := &struct {
		Name string `form:"name"`
		Brand string `form:"brand"`
		Stock int `form:"stock"`
		Seller string `form:"seller"`
	}{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seller := pc.sellerRepository.FindByUUID(request.Seller)

	if seller == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Seller is not found")})
		return
	}

	product := &product{
		UUID:      uuid.New().String(),
		Name:      request.Name,
		Brand:     request.Brand,
		Stock:     request.Stock,
		SellerUUID:    seller.UUID,
	}

	pc.repository.insert(product)
	jsonData, err := json.Marshal(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}

func (pc *controller) Put(c *gin.Context) {
	queryRequest := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(queryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := pc.repository.findByUUID(queryRequest.UUID)

	request := &struct {
		Name string `form:"name"`
		Brand string `form:"brand"`
		Stock int `form:"stock"`
	}{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldStock := product.Stock

	product.Name = request.Name
	product.Brand = request.Brand
	product.Stock = request.Stock

	pc.repository.update(product)

	if oldStock != product.Stock {
		seller := pc.sellerRepository.FindByUUID(product.SellerUUID)
		pc.sellerEmailProvider.StockChanged(oldStock, product.Stock, seller.Email)
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}

func (pc *controller) Delete(c *gin.Context) {
	request := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := pc.repository.findByUUID(request.UUID)

	if product == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Product is not found")})
		return
	}

	pc.repository.delete(product)
	c.JSON(http.StatusOK, gin.H{})
}
