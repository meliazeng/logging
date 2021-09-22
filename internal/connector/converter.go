package connector

import (
	"cap-logging-service/internal/types"
	"fmt"
	"strings"

	syslog "github.com/RackSec/srslog"
)

type GenericEvent struct {
	Category  string
	Action    string
	Arguments string
	Context   map[string]string
}

type Converter interface {
	ConvertUAAEvent(data string) (LogMessage, error)
}
type PubSubEventConverter func(string) (LogMessage, error)

func FormatForBalabit(priority syslog.Priority, envelope types.ExtendedEnvelope) LogMessage {
	namespace := envelope.Resource.Labels.NamespaceId
	podId := envelope.Resource.Labels.PodId
	if len(namespace) < 1 {
		namespace = envelope.Resource.Labels.NamespaceName
	}

	if len(podId) < 1 {
		podId = envelope.Resource.Labels.PodName
		if len(podId) < 1 {
			podId = "unknown"
		}
	}

	envelope.TextPayload = strings.ReplaceAll(envelope.TextPayload, "\n", "\\n")
	return LogMessage{
		Message:  fmt.Sprintf("%s[%s] %s", namespace, podId, envelope.TextPayload),
		Priority: priority,
	}
}
