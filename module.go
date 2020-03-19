package moduleDemo

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func Userdata() User {
	users := User{
		ID:    1,
		Name:  "mangesh",
		Email: "mangesh@gmail.com",
		Phone: "9970814741"}
	return users

}
