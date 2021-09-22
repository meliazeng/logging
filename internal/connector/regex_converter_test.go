package connector

import (
	"strings"
	"testing"

	syslog "github.com/RackSec/srslog"
	"github.com/stretchr/testify/assert"
)

func TestConvertRegexBadLog(t *testing.T) {
	//mess := "DLBPApphost:\\"
	str, _ := RegexConverter{}.ConvertUAAEvent("{\"")
	assert.Equal(t, str.Message, "", "String should be empty")
}

func TestConvertRegex(t *testing.T) {
	mess := "DLBPApphost:\\"
	str, _ := RegexConverter{}.ConvertUAAEvent("{\"")
	assert.Equal(t, "XXX[XXX-9242342345-34343] "+strings.ReplaceAll(mess, "\\", ""), str.Message, "Initial log line should be unchanged")
	assert.Equal(t, str.Message, "", "String should be empty")
}

func TestConvertRegexNoELCS(t *testing.T) {
	mess := "DLBPApphost:\\"
	str, _ := RegexConverter{}.ConvertUAAEvent("")
	assert.Equal(t, "XXX[XXX-934234342-342343] "+strings.ReplaceAll(mess, "\\", ""), str.Message, "Initial log line should be unchanged")
	assert.Equal(t, syslog.LOG_ERR, str.Priority, "Should be LOG_ERR")
}
