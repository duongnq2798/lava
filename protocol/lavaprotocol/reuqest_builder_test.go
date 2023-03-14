package lavaprotocol

import (
	"context"
	"testing"

	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/relayer/sigs"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	"github.com/stretchr/testify/require"
)

func TestSignAndExtract(t *testing.T) {
	ctx := context.Background()
	sk, address := sigs.GenerateFloatingKey()
	chainId := "LAV1"
	epoch := int64(100)
	singleConsumerSession := &lavasession.SingleConsumerSession{
		CuSum:                       20,
		LatestRelayCu:               10, // set by GetSession cuNeededForSession
		QoSInfo:                     lavasession.QoSReport{LastQoSReport: &pairingtypes.QualityOfServiceReport{}},
		SessionId:                   123,
		Client:                      nil,
		RelayNum:                    1,
		LatestBlock:                 epoch,
		Endpoint:                    nil,
		BlockListed:                 false, // if session lost sync we blacklist it.
		ConsecutiveNumberOfFailures: 0,     // number of times this session has failed
	}
	commonData := NewRelayRequestCommonData(chainId, "GET", "stub_url", []byte("stub_data"), 10, "tendermintrpc")
	relay, err := ConstructRelayRequest(ctx, sk, chainId, commonData, "lava@stubProviderAddress", singleConsumerSession, epoch, []byte("stubbytes"))
	require.Nil(t, err)

	// check signature
	extractedConsumerAddress, err := sigs.ExtractSignerAddress(relay)
	require.Nil(t, err)
	require.Equal(t, extractedConsumerAddress, address)
}