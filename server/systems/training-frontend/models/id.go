package models

type GetID struct {
	ID int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
}

type StringID struct {
	ID string `json:"id,omitempty" form:"id" validate:"required"`
}

type CreateIDs struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric"`
	CreatedBy int32 `json:"created_by,omitempty" form:"created_by" validate:"numeric,required"`
}
type UpdateIDs struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
	UpdatedBy int32 `json:"updated_by,omitempty" form:"updated_by" validate:"numeric,required"`
}

type VerifyIDs struct {
	ID         int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
	VerifiedBy int32 `json:"verified_by,omitempty" form:"verified_by" validate:"numeric,required"`
}

type DeleteIDs struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
	DeletedBy int32 `json:"deleted_by,omitempty" form:"deleted_by" validate:"numeric,required"`
}

type AttachmentDetail struct {
	Path           string `json:"path"`
	AttachmentName string `json:"attachment_name"`
}
type UpdateOnlyIDs struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric"`
	UpdatedBy int32 `json:"updated_by,omitempty" form:"updated_by" validate:"numeric,required"`
}

type CheckExistance struct {
	Exists bool  `json:"exists,omitempty" form:"exists"`
	ID     int32 `json:"id,omitempty" form:"id" validate:"numeric"`
}
type CloseWindow struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"numeric,required"`
	ClosedBy  int32 `json:"closed_by,omitempty" form:"closed_by" validate:"numeric,required"`
	DeletedBy int32 `json:"deleted_by,omitempty" form:"deleted_by" validate:"numeric,required"`
}

type CustomIDModel struct {
	ID              int32  `json:"id,omitempty" form:"id" `
	ClassID         int32  `json:"class_id,omitempty" form:"class_id" `
	ExamComponentID int32  `json:"exam_component_id,omitempty" form:"exam_component_id"`
	ModuleID        int32  `json:"module_id,omitempty" form:"module_id"`
	Role            string `json:"role[]" form:"role[]" `
}

type CustomIDModelForAttachment struct {
	ID              int32  `json:"id,string,omitempty" form:"id" `
	ModuleID        int32  `json:"module_id,string,omitempty" form:"module_id" `
	ClassID         int32  `json:"class_id,string,omitempty" form:"class_id" `
	ExamType        bool   `json:"exam_type,string,omitempty" form:"exam_type"`
	ClassName       string `json:"class_name" form:"class_name" `
	ExamComponentID int32  `json:"exam_component_id,string,omitempty" form:"exam_component_id"`
	FeeStructure    string `json:"fee_structure" form:"fee_structure"`
	NTALevel        string `json:"nta_level" form:"nta_level"`
	StudyYear       string `json:"study_year" form:"study_year"`
	Curriculum      string `json:"curriculum" form:"curriculum"`
	Program         string `json:"program" form:"program"`
}
