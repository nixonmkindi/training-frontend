package models

type ID struct {
	ID int32 `json:"id" form:"id" validate:"required,numeric"`
}

type StringID struct {
	ID string `json:"id" form:"id" validate:"required"`
}
