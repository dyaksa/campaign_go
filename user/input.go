package user

type RegisterUserInput struct {
	Email          string `json:"email" binding:"required,email"`
	Name           string `json:"name" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
	Password       string `json:"password" binding:"required"`
	AvatarFileName string `json:"avatar_file_name"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required"`
}
