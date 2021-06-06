package v1

import (
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/api/utils"
	"coding-challenge-go/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewController(repository *Repository) *controller {
	return &controller{
		gdgLogger:  logger.WithPrefix("v1-seller-controller"),
		repository: repository,
	}
}

type controller struct {
	gdgLogger  logger.Logger
	repository *Repository
}

func (pc *controller) List(c *gin.Context) {
	sellers, err := pc.repository.list()

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToQuerySellerList, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToQuerySellerList})
		return
	}

	sellersJson, err := json.Marshal(sellers)

	if err != nil {
		pc.gdgLogger.Errorln(dictionary.FailToMarshalSellers, err)
		c.JSON(http.StatusInternalServerError, errorResponse{ErrorMessage: dictionary.FailToMarshalSellers})
		return
	}

	c.Data(http.StatusOK, utils.CONTENT_TYPE_DEFAULT, sellersJson)
}
