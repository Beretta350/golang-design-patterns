package singleton_test

import (
	"testing"

	"github.com/Beretta350/golang-design-patterns/singleton"
	"github.com/stretchr/testify/assert"
)

func TestSingleInstance(t *testing.T) {
	firstInstance := singleton.SingleInstance()
	secondInstance := singleton.SingleInstance()
	thirdInstance := singleton.SingleInstance()

	assert.Equal(t, firstInstance, secondInstance)
	assert.Equal(t, secondInstance, thirdInstance)
}
