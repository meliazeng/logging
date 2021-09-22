package connector

import (
	"cap-logging-service/internal/types"
	"strings"
)

func ShouldForward(envelope types.ExtendedEnvelope) bool {
	if len(strings.TrimSpace(envelope.TextPayload)) == 0 {
		return false
	}
	return true
}
