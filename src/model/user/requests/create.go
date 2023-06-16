package requests

type Request struct {
	Username  string `json:"username" binding:"required,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Name      string `json:"name" binding:"required,min=3,max=40"`
	Password  string `json:"password" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}
