package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"training-frontend/server/services/database"
	"training-frontend/server/services/entity"
	"time"

	"training-frontend/package/util"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

// AuditTrailsRepoConn Initializes connection to DB
type AuditTrailsRepoConn struct {
	conn *pgxpool.Pool
}

// NewAuditTrails Connects to DB
func NewAuditTrails() *AuditTrailsRepoConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &AuditTrailsRepoConn{
		conn: conn,
	}
}

// Create Inserts new record to DB without ID
func (con *AuditTrailsRepoConn) Create(e *entity.AuditTrails) (int32, error) {
	var id int32
	query := `INSERT INTO audit_trails (ip_address,client,user_id,action,url,data,created_at) 
				VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.IPAddress, e.Client, e.UserID, e.Action, e.Url, e.Data, time.Now()).Scan(&id)
	return id, err
}

// CreateByID  Inserts new record to DB with ID
func (con *AuditTrailsRepoConn) CreateByID(e *entity.AuditTrails) (int32, error) {
	var id int32
	query := `INSERT INTO audit_trails (id,ip_address,client,user_id,action,url,data,created_at) 
 	             VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.ID, e.IPAddress, e.Client, e.UserID, e.Action, e.Url, e.Data, time.Now()).Scan(&id)

	return id, err
}

// CheckAuditTrailsID Checks if record exists in DB
func (con *AuditTrailsRepoConn) CheckAuditTrails(id int32) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM audit_trails WHERE id  = $1)"
	err := con.conn.QueryRow(context.Background(), query, id).Scan(&exists)
	return exists, err
}

// GetLastCreated Gets the CreatedAt date for last inserted record
func (con *AuditTrailsRepoConn) GetLastCreated() (time.Time, error) {
	var date pgtype.Timestamp
	var query = `SELECT created_at FROM audit_trails ORDER BY created_at DESC LIMIT 1`
	err := con.conn.QueryRow(context.Background(), query).Scan(&date)
	if util.IsError(err) {
		return time.Time{}, err
	}
	return date.Time, err
}

// GetLastUpdated Gets the UpdatedAt date for last updated record
func (con *AuditTrailsRepoConn) GetLastUpdated() (time.Time, error) {
	var date pgtype.Timestamp
	var query = `SELECT updated_at FROM audit_trails ORDER BY updated_at DESC LIMIT 1`
	err := con.conn.QueryRow(context.Background(), query).Scan(&date)
	if util.IsError(err) {
		return time.Time{}, err
	}
	return date.Time, err
}

func (con *AuditTrailsRepoConn) GetLastDeleted() (time.Time, error) {
	var date pgtype.Timestamp
	var query = `SELECT deleted_at FROM audit_trails ORDER BY deleted_at DESC LIMIT 1`
	err := con.conn.QueryRow(context.Background(), query).Scan(&date)
	if util.IsError(err) {
		return time.Time{}, err
	}
	return date.Time, err
}

// List Lists all records
func (con *AuditTrailsRepoConn) List() ([]*entity.AuditTrails, error) {

	var id pgtype.Int4
	var ipAddress, client, action, method, url, data pgtype.GenericText
	var userID pgtype.Int4
	var userEmail, userName string
	var createdAt, deletedAt pgtype.Timestamp

	var auditTrails []*entity.AuditTrails
	var query = `SELECT
					audit_trails.id,
					audit_trails.ip_address,
					audit_trails.client,
					audit_trails.action,
					audit_trails.method,
					audit_trails.url,
					audit_trails.data,
					usr.id,
					usr.name,
					usr.email,
					audit_trails.created_at,
					audit_trails.deleted_at
				FROM audit_trails
				INNER JOIN "user" usr ON usr.id = audit_trails.user_id`
	rows, err := con.conn.Query(context.Background(), query)
	if util.IsError(err) {
		return []*entity.AuditTrails{}, errors.New("error selecting region")
	}
	for rows.Next() {
		if err := rows.Scan(&id, &ipAddress, &client, &action, &method, &url, &data, &userID, &userName, &userEmail, &createdAt, &deletedAt); util.IsError(err) {
			return []*entity.AuditTrails{}, err
		}
		usr := &entity.User{
			ID:    userID.Int,
			Name:  userName,
			Email: userEmail,
		}
		auditTrail := &entity.AuditTrails{
			ID:        id.Int,
			IPAddress: ipAddress.String,
			Client:    client.String,
			Action:    action.String,
			User:      usr,
			Method:    method.String,
			Url:       url.String,
			Data:      data.String,
			CreatedAt: createdAt.Time,
			DeletedAt: deletedAt.Time,
		}
		auditTrails = append(auditTrails, auditTrail)
	}
	return auditTrails, err
}

// Get Gets single record by ID field
func (con *AuditTrailsRepoConn) Get(id int32) ([]*entity.AuditTrails, error) {
	var ipAddress, client, action, method, url, data pgtype.GenericText
	var userID pgtype.Int4
	var userEmail, userName string
	var createdAt, deletedAt pgtype.Timestamp
	var auditTrails []*entity.AuditTrails
	query := `SELECT
					audit_trails.id,
					audit_trails.ip_address,
					audit_trails.client,
					audit_trails.action,
					audit_trails.method,
					audit_trails.url,
					audit_trails.data,
					"user".id,
					"user".name,
					"user".email,
					audit_trails.created_at,
					audit_trails.deleted_at
	            FROM audit_trails
			 INNER JOIN "user" ON "user".id = audit_trails.user_id WHERE audit_trails.deleted_at is null`
	err := con.conn.QueryRow(context.Background(), query).Scan(&id, &ipAddress, &client, &action, &method, &url, &data, &userID, &userName, &userEmail, &createdAt, &deletedAt)

	if util.IsError(err) {
		return []*entity.AuditTrails{}, err
	}
	usr := &entity.User{
		ID:    userID.Int,
		Name:  userName,
		Email: userEmail,
	}
	auditTrail := &entity.AuditTrails{
		ID:        id,
		IPAddress: ipAddress.String,
		Client:    client.String,
		Action:    action.String,
		Method:    method.String,
		Url:       url.String,
		User:      usr,
		Data:      data.String,
		CreatedAt: createdAt.Time,
		DeletedAt: deletedAt.Time,
	}
	auditTrails = append(auditTrails, auditTrail)
	return auditTrails, err

}
