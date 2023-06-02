package user

type Request struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3,max=40"`
	Password string `json:"password" binding:"required"`
	Age      int8   `json:"age" binding:"required"`
}
