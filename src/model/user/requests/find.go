package requests

type FindByUsernameRequest struct {
	Username string `json:"username" bson:"username" uri:"username" binding:"required"`
}

type FindUserByEmailRequest struct {
	Email string `json:"email" uri:"email" binding:"required,email"`
}
