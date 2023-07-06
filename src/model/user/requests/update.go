package requests

type EditUser struct {
	Username  string  `json:"username,omitempty" bson:"username,omitempty" binding:"required,min=3,max=30"`
	Email     *string `json:"email,omitempty" bson:"email,omitempty"`
	Name      *string `json:"name,omitempty" bson:"name,omitempty"`
	Password  *string `json:"password,omitempty" bson:"password,omitempty"`
	BirthDate *string `json:"birth_date,omitempty" bson:"birthdate,omitempty"`
}
