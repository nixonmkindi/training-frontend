package models

//RegisterRequest defines logs struct
type RegisterRequest struct {
	Password        string `json:"password" form:"password" validate:"required,min=8"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	UserCategoryID  int32  `json:"user_category_id" form:"user_category_id" validate:"required,numeric"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=1,eqfield=Password"`
}
