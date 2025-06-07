package entity

import (
	"errors"
	"time"

	"training-frontend/package/util"
)

// AuditTrails DataStructure
type AuditTrails struct {
	ID        int32
	IPAddress string
	Client    string
	UserID    int32
	Action    string
	Method    string
	User      *User
	Url       string
	Data      string
	CreatedAt time.Time
	DeletedAt time.Time
	DeletedBy int32
}

// NewAuditTrails Instantiate the new object when inserting new record without ID
func NewAuditTrails(userID int32, ipAddress, client, action, url string,
	data string) (*AuditTrails, error) {

	auditTrails := &AuditTrails{
		IPAddress: ipAddress,
		Client:    client,
		UserID:    userID,
		Action:    action,
		Url:       url,
		Data:      data,
	}
	err := auditTrails.ValidateNewAuditTrails()
	if util.IsError(err) {
		return &AuditTrails{}, err
	}
	return auditTrails, err

}

// NewAuditTrailsWithID NewAuditTrails Instantiate the new object when inserting new record without ID
func NewAuditTrailsWithID(id, userID int32, ipAdress, client, action, url string,
	data string) (*AuditTrails, error) {

	auditTrails := &AuditTrails{
		ID:        id,
		IPAddress: ipAdress,
		Client:    client,
		UserID:    userID,
		Action:    action,
		Url:       url,
		Data:      data,
	}
	err := auditTrails.ValidateNewAuditTrails()
	if util.IsError(err) {
		return &AuditTrails{}, err
	}
	return auditTrails, err

}

// ValidateNewAuditTrails Validates when inserting new record without ID
func (r *AuditTrails) ValidateNewAuditTrails() error {
	if r.IPAddress == "" {
		return errors.New("error validating audit-trails entity,ip address field required")
	}
	if r.Action == "" {
		return errors.New("error validating audit-trails entity,action field required")
	}
	if r.Url == "" {
		return errors.New("error validating audit-trails entity,url field required")
	}
	return nil
}

// ValidateNewAuditTrailsWithID Validates when inserting new record with ID
func (r *AuditTrails) ValidateNewAuditTrailsWithID() error {
	if r.ID <= 0 {
		return errors.New("error validating audit-trails entity,id field required")
	}
	if r.IPAddress == "" {
		return errors.New("error validating audit-trails entity,ip address field required")
	}
	if r.Action == "" {
		return errors.New("error validating audit-trails entity,action field required")
	}
	if r.Url == "" {
		return errors.New("error validating audit-trails entity,url field required")
	}
	return nil
}

// ValidateUpdateAuditTrails Validates when updating without ID
func (r *AuditTrails) ValidateUpdateAuditTrails() error {
	if r.ID <= 0 {
		return errors.New("error validating audit-trails entity,id field required")
	}
	if r.IPAddress == "" {
		return errors.New("error validating audit-trails entity,ip address field required")
	}
	if r.Action == "" {
		return errors.New("error validating audit-trails entity,action field required")
	}
	if r.Url == "" {
		return errors.New("error validating audit-trails entity,url field required")
	}
	return nil
}
