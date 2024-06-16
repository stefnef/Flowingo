package pkg

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var generator = NewGeneratorImpl()

func TestGeneratorImpl_ImplementsInterface(t *testing.T) {
	assert.Implements(t, (*Generator)(nil), generator)
}

func TestGeneratorImpl_GenerateUUID(t *testing.T) {
	var uuid = generator.GenerateUUID()
	var anotherUuid = generator.GenerateUUID()

	assert.NotEmpty(t, uuid)
	assert.True(t, isUUID(uuid))
	assert.NotEqual(t, anotherUuid, uuid)
}

func TestGeneratorImpl_GenerateNumber(t *testing.T) {
	var number = generator.GenerateNumber()
	var anotherNumber = generator.GenerateNumber()

	assert.NotEqual(t, 0, number)
	assert.NotEqual(t, anotherNumber, number)
}

func isUUID(uuid string) bool {
	re := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return re.MatchString(uuid)
}
