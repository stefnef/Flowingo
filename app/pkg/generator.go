package pkg

import (
	"github.com/google/uuid"
	"math/rand"
)

type Generator interface {
	GenerateUUID() string
	GenerateNumber() int
}

type GeneratorImpl struct {
	uuid uuid.UUID
	rand *rand.Rand
}

func (g GeneratorImpl) GenerateUUID() string {
	return g.uuid.String()
}

func (g GeneratorImpl) GenerateNumber() int {
	return g.rand.Int()
}

func NewGeneratorImpl() GeneratorImpl {
	return GeneratorImpl{
		uuid: uuid.New(),
		rand: rand.New(rand.NewSource(int64(rand.Uint64()))),
	}
}
