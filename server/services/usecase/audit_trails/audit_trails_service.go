package audit_trails

import (
	"errors"
	"training-frontend/package/log"
	"training-frontend/server/services/entity"
	"training-frontend/server/services/repository"
)

// Service Initialize repository
type Service struct {
	repo Repository
}

// NewService Instantiate new service
func NewService() *Service {
	repo := repository.NewAuditTrails()
	return &Service{
		repo: repo,
	}
}

// CreateAuditTrails Calls create new record repository without ID
func (s *Service) CreateAuditTrails(userID int32, ipAddress, client, action, url, data string) (int32, error) {
	auditTrail, err := entity.NewAuditTrails(userID, ipAddress, client, action, url, data)
	if err != nil {
		return auditTrail.ID, err
	}
	auditTrailID, err := s.repo.Create(auditTrail)
	if err != nil {
		log.Errorf("error creating audit trail: %v", err)
		return auditTrail.ID, errors.New("cannot create record")
	}
	return auditTrailID, err
}

// CreateAllAuditTrailsByID Creates many records with ID at once
func (s *Service) CreateAllAuditTrailsByID(e []*entity.AuditTrails) (int32, error) {

	for _, n := range e {
		auditTrail, err := entity.NewAuditTrailsWithID(n.ID, n.UserID, n.IPAddress, n.Client, n.Action, n.Url, n.Data)
		if err != nil {
			return auditTrail.ID, err
		}
		status, err := s.CheckAuditTrailsExistence(n.ID)
		if err != nil {
			return n.ID, err
		}
		if status {
			continue
		}
		auditTrailID, err := s.repo.CreateByID(n)
		if err != nil {
			return auditTrailID, errors.New("cannot create record")
		}
	}
	return 0, nil
}

// CheckAuditTrailsExistence Calls checks if record exists in DB repository
func (s *Service) CheckAuditTrailsExistence(id int32) (bool, error) {
	exist, err := s.repo.CheckAuditTrails(id)
	if err != nil {
		return exist, errors.New("no record found")
	}
	return exist, err
}

// ListAuditTrail Calls list records repository
func (s *Service) ListAuditTrail() ([]*entity.AuditTrails, error) {

	auditTrails, err := s.repo.List()
	if err == nil {
		return auditTrails, err
	}
	if err.Error() == "no rows in result set" {
		return auditTrails, err
	}
	if err != nil {
		return auditTrails, err
	}
	return auditTrails, err
}

// GetAuditTrail Calls list records repository
func (s *Service) GetAuditTrail(id int32) ([]*entity.AuditTrails, error) {

	auditTrails, err := s.repo.Get(id)
	if err == nil {
		return auditTrails, err
	}
	if err.Error() == "no rows in result set" {
		return auditTrails, err
	}
	if err != nil {
		return auditTrails, errors.New("no record found")
	}
	return auditTrails, err
}
