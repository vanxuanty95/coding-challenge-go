package v2

import (
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/api/utils"
	"coding-challenge-go/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	TOP_SELLER_PARAMETER = "number"
	TOP_SELLER_DEFAULT   = 10
)

func NewController(repository *Repository) *controller {
	return &controller{
		gdgLogger:  logger.WithPrefix("v2-seller-controller"),
		repository: repository,
	}
}

type controller struct {
	gdgLogger  logger.Logger
	repository *Repository
}

func (pc *controller) TopSeller(c *gin.Context) {
	number := TOP_SELLER_DEFAULT
	numberAsString := c.Param(TOP_SELLER_PARAMETER)
	if numberAsString != "" {
		numberParse, err := strconv.Atoi(numberAsString)
		if err != nil {
			pc.gdgLogger.Errorln(dictionary.TopSellerMustANumber, err)
			c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.TopSellerMustANumber})
			return
		}
		number = numberParse
	}

	if number <= 0 {
		pc.gdgLogger.Error(dictionary.TopSellerMustGreaterThan0)
		c.JSON(http.StatusBadRequest, errorResponse{ErrorMessage: dictionary.TopSellerMustGreaterThan0})
		return
	}

	topSellers, err := pc.repository.getTopSellers(number)
	if err != nil {
		pc.gdgLogger.Errorf(dictionary.FailToQueryTopNSellers, number, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: fmt.Sprintf(dictionary.FailToQueryTopNSellers, number)})
		return
	}

	sellersJson, err := json.Marshal(topSellers)

	if err != nil {
		pc.gdgLogger.Errorf(dictionary.FailToMarshalTopNSellers, number, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: fmt.Sprintf(dictionary.FailToMarshalTopNSellers, number)})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, sellersJson)
}
