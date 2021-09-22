package connector

import (
	"cap-logging-service/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFilterEmpty(t *testing.T) {
	envelope := types.ExtendedEnvelope{TextPayload: "\n"}
	fwd := ShouldForward(envelope)
	assert.False(t, fwd, "Should not forward whitespace")
}

func TestShouldFilterNFT(t *testing.T) {
	envelope := types.ExtendedEnvelope{TextPayload: "dummy\n"}
	envelope.Labels = types.PodLabels{NameSpaceId: "test-nft"}
	fwd := ShouldForward(envelope)
	assert.False(t, fwd, "Should not forward NFT lgos")

}
