package users

type User struct {
	Id        int `json:"id"`
	FirstName int `json:"first_name"`
	LastName  int `json:"last_name"`
	Email     int `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
