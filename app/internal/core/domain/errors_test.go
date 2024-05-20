package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_not_found_should_have_id_and_name(t *testing.T) {
	err := NewNotFoundError("the-resource", "some-id")

	assert.Implements(t, (*error)(nil), err)
	assert.ErrorContains(t, err, "could not find resource 'the-resource' with id 'some-id'")
}
