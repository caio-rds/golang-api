package user

type FindUserByIdRequest struct {
	ID int64 `json:"id" uri:"id" binding:"required"`
}
