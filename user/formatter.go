package user

type UserFormatter struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	Occupation string `json:"occupation"`
	Email string `json:"email"`
	Token string `json:"token"`

}

func FormatterUser(user User, token string) UserFormatter{
	formatter := UserFormatter{
		ID: user.ID,
		Nama: user.Nama,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: token,
	}

	return formatter
}