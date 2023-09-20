package user

type RegisterUserInput struct {
	Name string 	  `json:"name" binding:"required"`
	Email string 	  `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
