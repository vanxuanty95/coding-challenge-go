package seller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewController(repository *Repository) *controller {
	return &controller{
		repository: repository,
	}
}

type controller struct {
	repository *Repository
}

func (pc *controller) List(c *gin.Context) {
	sellers := pc.repository.list()
	sellersJson, err := json.Marshal(sellers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", sellersJson)
}
