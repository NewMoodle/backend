package dto

type UserCreateDto struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"email,required"`
	Password  string `json:"password" binding:"required,min=8,max=64"`
	RoleName  string `json:"role" binding:"required"`
}
