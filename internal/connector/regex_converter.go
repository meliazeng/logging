package connector

import (
	//"cap-logging-service/internal/connector/csel"
	"cap-logging-service/internal/types"
	"encoding/json"
	"strings"

	syslog "github.com/RackSec/srslog"
)

type RegexConverter struct {
}

/*
var (
	keyValuePattern = regexp.MustCompile(`(\w+):([^,\[\]\(\)]+)`)
)
*/
func (r RegexConverter) ConvertUAAEvent(data string) (LogMessage, error) {

	var envelope types.ExtendedEnvelope
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	err := d.Decode(&envelope)

	if err != nil {
		return LogMessage{}, err
	}

	if ShouldForward(envelope) {

		//rawVals := keyValuePattern.FindAllStringSubmatch(envelope.TextPayload, -1)
		//vmap := createMap(rawVals)
		priority := syslog.LOG_LOCAL0 //csel.GetPriority(vmap)

		logMessage := FormatForBalabit(priority, envelope)
		return logMessage, nil
	}
	return LogMessage{}, nil
}
