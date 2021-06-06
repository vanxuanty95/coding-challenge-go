package v2

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/api/utils"
	"coding-challenge-go/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	LIST_PAGE_SIZE = 10
	SELF_LINK_KEY  = "self"
)

func NewController(cfg *config.Config, repository Repository) *controller {
	return &controller{
		gdgLogger:  logger.WithPrefix("v2-product-controller"),
		cfg:        cfg,
		repository: repository,
	}
}

type controller struct {
	gdgLogger  logger.Logger
	cfg        *config.Config
	repository Repository
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

	for _, currProduct := range products {
		currProduct.Seller = pc.getSeller(currProduct.SellerUUID)
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
		pc.gdgLogger.Errorln(dictionary.FailToQueryProductByUUID)
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.ProductIsNotFound})
		return
	}

	product.Seller = pc.getSeller(product.SellerUUID)

	productJson, err := json.Marshal(product)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalProduct, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalProduct})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, productJson)
}

func (pc *controller) getSeller(sellerUUID string) productSeller {
	links := make(map[string]link)
	hrefSelf := fmt.Sprintf("http://%v:%v/sellers/%v", pc.cfg.RestfulAPI.Host, pc.cfg.RestfulAPI.Port, sellerUUID)
	links[SELF_LINK_KEY] = link{Href: hrefSelf}
	sellerObj := productSeller{
		UUID:  sellerUUID,
		Links: links,
	}
	return sellerObj
}
