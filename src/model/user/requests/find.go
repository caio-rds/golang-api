package requests

type FindByUsername struct {
	Username string `json:"username" bson:"username" uri:"username" binding:"required"`
}

type FindUserByEmail struct {
	Email string `json:"email" uri:"email" binding:"required,email"`
}
