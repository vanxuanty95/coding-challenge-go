package v1

import (
	"coding-challenge-go/pkg/logger"
	"github.com/golang/mock/gomock"
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
			name: "exists",
			args: args{
				c: func() *gin.Context {
					c, _ := gin.CreateTestContext(httptest.NewRecorder())
					c.Request, _ = http.NewRequest("GET", "/v1/sellers", nil)
					return c
				}(),
			},
			mockRepo: func() Repository {
				r := NewMockRepository(gomock.NewController(t))
				r.EXPECT().list().Return([]*Seller{}, nil)
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
			pc.List(tt.args.c)
		})
	}
}
