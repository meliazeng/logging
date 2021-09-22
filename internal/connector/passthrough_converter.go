package connector

import (
	"cap-logging-service/internal/types"
	"encoding/json"
	"strings"

	syslog "github.com/RackSec/srslog"
)

type PassthroughConverter struct {
}

func (p PassthroughConverter) ConvertUAAEvent(data string) (LogMessage, error) {

	var envelope types.ExtendedEnvelope
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	err := d.Decode(&envelope)

	if err != nil {
		return LogMessage{}, err
	}

	if ShouldForward(envelope) {
		logMessage := FormatForBalabit(syslog.LOG_INFO|syslog.LOG_LOCAL0, envelope)
		return logMessage, err
	}
	return LogMessage{}, nil

}
