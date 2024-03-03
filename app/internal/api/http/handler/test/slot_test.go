package handler_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type slotParam struct {
	name  string
	value interface{}
}

type Slot struct {
	t         *testing.T
	functions map[string]*[]slotParam
}

func (s *Slot) appendParameter(function string, parameterName string, parameterValue interface{}) {
	var params = s.functions[function]
	*params = append(*params, slotParam{
		name:  parameterName,
		value: parameterValue,
	})
}

func (s *Slot) getValues(function string, parameterName string) []interface{} {
	var result []interface{}
	for _, value := range *s.functions[function] {
		if value.name == parameterName {
			result = append(result, value.value)
		}
	}
	return result
}

func (s *Slot) wasCalledWith(function string, paramName string, paramValue interface{}) bool {
	values := s.getValues(function, paramName)
	for _, value := range values {
		if value == paramValue {
			return true
		}
	}
	s.t.Logf("function %s was not called with parameter '%s' and value '%s'", function, paramName, paramValue)
	if len(values) > 0 {
		s.t.Logf("possible calls were found with value: '%s'", values)
	}
	return false
}

func (s *Slot) verify(function string, parameterName string, value string) {
	calledWith := s.wasCalledWith(function, parameterName, value)
	assert.True(s.t, calledWith)
}

func (s *Slot) verifyFunctionNotCalled(function string) {
	calls, exists := s.functions[function]
	assert.True(s.t, exists)
	assert.Empty(s.t, calls)
}
