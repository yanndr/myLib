package services

import "fmt"

// DuplicateErr represents an error while attempting to insert an item that already exists.
type DuplicateErr struct {
	entityType string
	entityID   any
}

// NewDuplicateErr returns a new DuplicateErr error.
func NewDuplicateErr(entityType string, entityID any) DuplicateErr {
	return DuplicateErr{
		entityID:   entityID,
		entityType: entityType,
	}
}

func (e DuplicateErr) Error() string {
	return fmt.Sprintf("error: entity %v of type %s already exists", e.entityID, e.entityType)
}

// ValidationErr represents an error that is raised when a structure is invalid.
type ValidationErr struct {
	entityType string
	reason     string
}

// NewValidationErr returns a new ValidationErr error
func NewValidationErr(entityType string, err string) ValidationErr {
	return ValidationErr{
		entityType: entityType,
		reason:     err,
	}
}

func (e ValidationErr) Error() string {
	return fmt.Sprintf("validation error: entity %s is invalid: %s", e.entityType, e.reason)
}
