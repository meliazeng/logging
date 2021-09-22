package connector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertUAAEventFromValid(t *testing.T) {
	exp := "Test Message"
	resp := "XXX[XXX-9XXXX5-XXXX] Test Message"
	str, err := PassthroughConverter{}.ConvertUAAEvent("{\"}")
	if err != nil {
		t.Errorf("Failed to parse: %s", err)
	}
	assert.Equal(t, resp, str.Message, "Log output not equal to expected")
}

func TestConvertUAAEventFromInvalid(t *testing.T) {
	str, err := PassthroughConverter{}.ConvertUAAEvent("FOOBAR!")

	assert.Error(t, err, "Should have an error duer to bad parse")
	assert.Empty(t, str, "Should be empty after error")
}
