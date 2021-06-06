package v1

import (
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/api/helper"
	v1Seller "coding-challenge-go/pkg/api/seller/v1"
	"coding-challenge-go/pkg/api/utils"
	"coding-challenge-go/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	LIST_PAGE_SIZE = 10
)

func NewController(repository Repository, sellerRepository v1Seller.Repository, notifiersFactory *helper.NotifiersFactory) *controller {
	return &controller{
		gdgLogger:        logger.WithPrefix("v1-product-controller"),
		repository:       repository,
		sellerRepository: sellerRepository,
		notifiersFactory: notifiersFactory,
	}
}

type controller struct {
	gdgLogger        logger.Logger
	repository       Repository
	sellerRepository v1Seller.Repository
	notifiersFactory *helper.NotifiersFactory
}

func (pc *controller) List(c *gin.Context) {
	request := &listRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		pc.gdgLogger.Errorln(err)
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	products, err := pc.repository.list((request.Page-1)*LIST_PAGE_SIZE, LIST_PAGE_SIZE)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQueryProductList, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQueryProductList})
		return
	}

	productsJson, err := json.Marshal(products)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalProducts, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalProducts})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, productsJson)
}

func (pc *controller) Get(c *gin.Context) {
	request := &getRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	product, err := pc.repository.findByUUID(request.UUID)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQueryProductByUUID, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQueryProductByUUID})
		return
	}

	if product == nil {
		pc.gdgLogger.Errorf(dictionary.FailToQueryProductByUUID)
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.ProductIsNotFound})
		return
	}

	productJson, err := json.Marshal(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalProduct})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, productJson)
}

func (pc *controller) Post(c *gin.Context) {
	request := &postRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	seller, err := pc.sellerRepository.FindByUUID(request.Seller)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQuerySellerByUUID, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQuerySellerByUUID})
		return
	}

	if seller == nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.SellerIsNotFound})
		return
	}

	product := &product{
		UUID:       uuid.New().String(),
		Name:       request.Name,
		Brand:      request.Brand,
		Stock:      request.Stock,
		SellerUUID: seller.UUID,
	}

	err = pc.repository.insert(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToInsertProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToInsertProduct})
		return
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalProduct})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, jsonData)
}

func (pc *controller) Put(c *gin.Context) {
	queryRequest := &putRequest{}

	if err := c.ShouldBindQuery(queryRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	product, err := pc.repository.findByUUID(queryRequest.UUID)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQueryProductByUUID, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQueryProductByUUID})
		return
	}

	if product == nil {
		pc.gdgLogger.Errorf(dictionary.ProductIsNotFound)
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.ProductIsNotFound})
		return
	}

	request := &putRequestBody{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	oldStock := product.Stock

	product.Name = request.Name
	product.Brand = request.Brand
	product.Stock = request.Stock

	err = pc.repository.update(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToInsertProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToInsertProduct})
		return
	}

	if oldStock != product.Stock {
		seller, err := pc.sellerRepository.FindByUUID(product.SellerUUID)

		if err != nil {
			pc.gdgLogger.Errorln(dictionary.FailToQuerySellerByUUID, err)
			c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQuerySellerByUUID})
			return
		}

		if seller == nil {
			pc.gdgLogger.Error(dictionary.SellerIsNotFound)
			c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.SellerIsNotFound})
			return
		}

		notificationInfo := generateNotificationInfo(seller, product, oldStock, product.Stock)
		go pc.notifiersFactory.SendNotification(notificationInfo)
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalProduct})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, jsonData)
}

func (pc *controller) Delete(c *gin.Context) {
	request := &deleteRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: err.Error()})
		return
	}

	product, err := pc.repository.findByUUID(request.UUID)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQueryProductByUUID, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQueryProductByUUID})
		return
	}

	if product == nil {
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.ProductIsNotFound})
		return
	}

	err = pc.repository.delete(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToDeleteProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToDeleteProduct})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func generateNotificationInfo(seller *v1Seller.Seller, prd *product, OldStock, NewStock int) helper.NotificationsInfo {
	return helper.NotificationsInfo{
		SellerUUID:  seller.UUID,
		SellerPhone: seller.Phone,
		SellerEmail: seller.Email,
		SellerName:  seller.Name,
		ProductName: prd.Name,
		OldStock:    OldStock,
		NewStock:    NewStock,
	}
}
