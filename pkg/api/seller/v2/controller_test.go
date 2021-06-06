package v2

import (
	"coding-challenge-go/pkg/logger"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_controller_TopSeller(t *testing.T) {
	gdgLogger := logger.WithPrefix("Test_controller_TopSeller")

	type fields struct {
		gdgLogger logger.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		args     args
		mockRepo func() Repository
		fields   fields
	}{
		{
			name: "top 1",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/sellers/top1", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().getTopSellers(10).Return([]*Seller{
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120002",
						Name:  "Christene Maggio",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
				}, nil)
				return r
			},
			fields: fields{
				gdgLogger: gdgLogger,
			},
		},
		{
			name: "top 10",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/sellers/top10", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().getTopSellers(10).Return([]*Seller{
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120002",
						Name:  "Christene Maggio1",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120003",
						Name:  "Christene Maggio2",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120004",
						Name:  "Christene Maggio3",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120005",
						Name:  "Christene Maggio4",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120006",
						Name:  "Christene Maggio5",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
					{
						UUID:  "8cabd5d0-c37e-11eb-b86b-0242ac120007",
						Name:  "Christene Maggio6",
						Email: "christene.maggio@seller.com",
						Phone: "202-555-0143",
					},
				}, nil)
				return r
			},
			fields: fields{
				gdgLogger: gdgLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &controller{
				gdgLogger:  tt.fields.gdgLogger,
				repository: tt.mockRepo(),
			}
			pc.TopSeller(tt.args.c)
		})
	}
}
