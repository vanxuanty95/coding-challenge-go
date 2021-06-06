package v1

type (
	listRequest struct {
		Page int `form:"page,default=1"`
	}
	getRequest struct {
		UUID string `form:"id" binding:"required"`
	}
	postRequest struct {
		Name   string `form:"name"`
		Brand  string `form:"brand"`
		Stock  int    `form:"stock"`
		Seller string `form:"seller"`
	}
	putRequest struct {
		UUID string `form:"id" binding:"required"`
	}
	putRequestBody struct {
		Name  string `form:"name"`
		Brand string `form:"brand"`
		Stock int    `form:"stock"`
	}
	deleteRequest struct {
		UUID string `form:"id" binding:"required"`
	}
)
