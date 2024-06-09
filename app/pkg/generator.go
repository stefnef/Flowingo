package pkg

import "github.com/google/uuid"

type Generator interface {
	GenerateUUID() string
}

type GeneratorImpl struct {
	uuid uuid.UUID
}

func (g GeneratorImpl) GenerateUUID() string {
	return g.uuid.String()
}

func NewGeneratorImpl() GeneratorImpl {
	return GeneratorImpl{
		uuid: uuid.New(),
	}
}
