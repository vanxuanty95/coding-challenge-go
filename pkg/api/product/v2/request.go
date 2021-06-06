package v2

type (
	listRequest struct {
		Page int `form:"page,default=1"`
	}
	getRequest struct {
		UUID string `form:"id" binding:"required"`
	}
)
