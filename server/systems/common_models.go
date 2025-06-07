package systems

// PrimaryKey struct
type UUIDModel struct {
	ID string `validate:"required" json:"id" form:"id"`
}

type PrimaryKey struct {
	ID string `validate:"required" json:"id" form:"id"`
}
