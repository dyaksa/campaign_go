package user

type Formatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	ImgeURL    string `json:"image_url"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) Formatter {
	jsonFormatter := Formatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		ImgeURL:    user.AvatarFileName,
		Email:      user.Email,
		Token:      token,
	}
	return jsonFormatter
}
