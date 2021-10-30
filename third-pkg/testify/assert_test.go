package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertOk(t *testing.T) {
	assert.Equal(t, true, true , "TestAssertOk")
}

func TestAssertFail(t *testing.T) {
	assert.Equal(t, true, false , "TestAssertFail")
}