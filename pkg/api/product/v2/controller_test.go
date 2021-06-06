package v2

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/logger"
	"github.com/golang/mock/gomock"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_controller_List(t *testing.T) {
	gdgLogger := logger.WithPrefix("Test_controller_List")

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
			name: "empty page",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/products", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().list(0, 10).Return([]*product{}, nil)
				return r
			},
			fields: fields{
				gdgLogger: gdgLogger,
			},
		},
		{
			name: "page = 0",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/products?page=0", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().list(-10, 10).Return(nil, nil)
				return r
			},
			fields: fields{
				gdgLogger: gdgLogger,
			},
		},
		{
			name: "page = 2",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/products?page=2", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().list(10, 10).Return(nil, nil)
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
				cfg:        generateCfg(),
				repository: tt.mockRepo(),
			}
			pc.List(tt.args.c)
		})
	}
}

func Test_controller_Get(t *testing.T) {
	gdgLogger := logger.WithPrefix("Test_controller_Get")

	type fields struct {
		gdgLogger  logger.Logger
		repository Repository
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
			name: "not exist",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/product?id=1", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().findByUUID("1").Return(nil, nil)
				return r
			},
			fields: fields{
				gdgLogger: gdgLogger,
			},
		},
		{
			name: "exist",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v2/product?id=8cac9bb8-c37e-11eb-b86b-0242ac120002", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().findByUUID("8cac9bb8-c37e-11eb-b86b-0242ac120002").Return(&product{
					ProductID:  0,
					UUID:       "8cac9bb8-c37e-11eb-b86b-0242ac120002",
					Name:       "test",
					Brand:      "test",
					Stock:      10,
					SellerUUID: "100",
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
				cfg:        generateCfg(),
				repository: tt.mockRepo(),
			}
			pc.Get(tt.args.c)
		})
	}
}

func generateCfg() *config.Config {
	//cfgPath := fmt.Sprintf("../../config/config.%v.yml", state) //TODO
	cfgPath := "../../../../cmd/api/config/config.local.yml"
	f, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}

	var cfg config.Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	cfg.State = "local"
	return &cfg
}
