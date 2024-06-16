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
	rand *rand.Rand
}

func (g GeneratorImpl) GenerateUUID() string {
	return uuid.NewString()
}

func (g GeneratorImpl) GenerateNumber() int {
	return g.rand.Int()
}

func NewGeneratorImpl() GeneratorImpl {
	return GeneratorImpl{
		rand: rand.New(rand.NewSource(int64(rand.Uint64()))),
	}
}
