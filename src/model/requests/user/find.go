package user

type FindUserByIdRequest struct {
	ID int64 `json:"id" uri:"id" binding:"required"`
}

type FindUserByEmailRequest struct {
	Email string `json:"email" uri:"email" binding:"required,email"`
}
