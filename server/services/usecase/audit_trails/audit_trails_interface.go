package audit_trails

import "training-frontend/server/services/entity"

// Reader interface
type Reader interface {
	CheckAuditTrails(id int32) (bool, error)
	List() ([]*entity.AuditTrails, error)
	Get(id int32) ([]*entity.AuditTrails, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.AuditTrails) (int32, error)
	CreateByID(e *entity.AuditTrails) (int32, error)
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
}
