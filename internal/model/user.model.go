package model

type User struct {
	ID           int      `json:"id"`
	Username     string   `json:"username"`
	Profile      *Profile `json:"profile"`
	PasswordHash string   `json:"password_hash"`
	Role         *Role    `json:"role"`
}

type Profile struct {
	ID        int    `json:"-"`
	UserID    int    `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
