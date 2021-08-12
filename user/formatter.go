package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Occuption string `json:"occuption"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Name:      user.Name,
		Occuption: user.Occupation,
		Email:     user.Email,
		Token:     token,
	}

	return formatter
}
