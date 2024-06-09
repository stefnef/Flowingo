package repository

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
)

type ResourceRepository interface {
	GetResources() []domain.Resource
	GetResourceById(id string) (*domain.Resource, error)
	ExistsResourceByName(name string) bool
	SaveResource(name string, magicNumber int) *domain.Resource
}
