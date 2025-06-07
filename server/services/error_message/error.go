package error_message

import "errors"

var (
	//ErrNotFound not found
	ErrNotFound      = errors.New("records not found")
	ErrWrongPassword = errors.New("incorrect password entered")
	//error_message.ErrNoResultSet.Error() not found
	ErrNoResultSet = errors.New("no rows in result set")
	//ErrNotFound not found
	ErrCannotBeCreated = errors.New("record cannot be created")
	//ErrCannotBeDeleted cannot be deleted
	ErrCannotBeDeleted = errors.New("record cannot be deleted")
	//ErrNotEnoughBooks cannot borrow
	ErrCannotBeUpdated = errors.New("record cannot be updated")
)
