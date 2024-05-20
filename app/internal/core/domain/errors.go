package domain

import "fmt"

type NotFoundError struct {
	Resource string
	ID       string
}

func NewNotFoundError(resource string, id string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("could not find resource '%s' with id '%s'", e.Resource, e.ID)
}
