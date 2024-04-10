package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	assert.NotNil(t, instance1)
	assert.Equal(t, 0, instance1.count)
	expectedCount := 1
	count := instance1.AddOne()
	assert.Equal(t, expectedCount, count)

	instance2 := GetInstance()
	assert.Equal(t, instance1, instance2)
	expectedCount = 2
	count = instance1.AddOne()
	assert.Equal(t, expectedCount, count)
}
