package connector

import (
	"cap-logging-service/internal/types"
	"testing"

	"github.com/RackSec/srslog"
	"github.com/stretchr/testify/assert"
)

func TestFormatForBalabit(t *testing.T) {
	envelope := types.ExtendedEnvelope{}
	envelope.Labels = types.PodLabels{NameSpaceId: "testns"}
	envelope.Resource.Labels = types.ResourceLabels{PodId: "testcont", NamespaceName: "testns", ClusterName: "testcont"}
	envelope.TextPayload = "Test!"
	resp := FormatForBalabit(srslog.LOG_INFO, envelope)
	assert.Equal(t, "testns[testcont] Test!", resp.Message, "Invalid balabit format in converted message")
}
