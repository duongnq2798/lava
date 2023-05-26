package chainlib

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	dyncodec "github.com/lavanet/lava/protocol/chainlib/grpcproxy/dyncodec"
	"github.com/lavanet/lava/protocol/parser"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/lavanet/lava/protocol/chainlib/grpcproxy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcInterfaceMessages"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcclient"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/protocol/metrics"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	reflectionpbo "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type GrpcChainParser struct {
	BaseChainParser

	registry *dyncodec.Registry
	codec    *dyncodec.Codec
}

// NewGrpcChainParser creates a new instance of GrpcChainParser
func NewGrpcChainParser() (chainParser *GrpcChainParser, err error) {
	return &GrpcChainParser{}, nil
}

func (apip *GrpcChainParser) setupForConsumer(relayer grpcproxy.ProxyCallBack) {
	apip.registry = dyncodec.NewRegistry(dyncodec.NewRelayerRemote(relayer))
	apip.codec = dyncodec.NewCodec(apip.registry)
}

func (apip *GrpcChainParser) setupForProvider(grpcEndpoint string) error {
	remote, err := dyncodec.NewGRPCReflectionProtoFileRegistry(grpcEndpoint)
	if err != nil {
		return err
	}
	apip.registry = dyncodec.NewRegistry(remote)
	apip.codec = dyncodec.NewCodec(apip.registry)
	return nil
}

func (apip *GrpcChainParser) CraftMessage(parsing *spectypes.Parsing, connectionType string, craftData *CraftData) (ChainMessageForSend, error) {
	if craftData != nil {
		return apip.ParseMsg(craftData.Path, craftData.Data, craftData.ConnectionType, nil)
	}

	grpcMessage := &rpcInterfaceMessages.GrpcMessage{
		Msg:  nil,
		Path: parsing.ApiName,
	}
	apiCont, err := apip.getSupportedApi(parsing.ApiName, connectionType)
	if err != nil {
		return nil, err
	}
	apiCollection, err := apip.getApiCollection(craftData.ConnectionType, apiCont.collectionKey.InternalPath, apiCont.collectionKey.Addon)
	if err != nil {
		return nil, err
	}
	return apip.newChainMessage(apiCont.api, spectypes.NOT_APPLICABLE, grpcMessage, apiCollection), nil
}

// ParseMsg parses message data into chain message object
func (apip *GrpcChainParser) ParseMsg(url string, data []byte, connectionType string, metadata []pairingtypes.Metadata) (ChainMessage, error) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return nil, errors.New("GrpcChainParser not defined")
	}

	// Check API is supported and save it in nodeMsg.
	apiCont, err := apip.getSupportedApi(url, connectionType)
	if err != nil {
		return nil, utils.LavaFormatError("failed to getSupportedApi gRPC", err)
	}

	// Construct grpcMessage
	grpcMessage := rpcInterfaceMessages.GrpcMessage{
		Msg:      data,
		Path:     url,
		Codec:    apip.codec,
		Registry: apip.registry,
		Header:   metadata,
	}

	// // Fetch requested block, it is used for data reliability
	// // Extract default block parser
	blockParser := apiCont.api.BlockParsing
	requestedBlock, err := parser.ParseBlockFromParams(grpcMessage, blockParser)
	if err != nil {
		return nil, utils.LavaFormatError("ParseBlockFromParams failed parsing block", err, utils.Attribute{Key: "chain", Value: apip.spec.Name}, utils.Attribute{Key: "blockParsing", Value: apiCont.api.BlockParsing})
	}

	apiCollection, err := apip.getApiCollection(connectionType, apiCont.collectionKey.InternalPath, apiCont.collectionKey.Addon)
	if err != nil {
		return nil, utils.LavaFormatError("failed to getApiCollection gRPC", err)
	}

	nodeMsg := apip.newChainMessage(apiCont.api, requestedBlock, &grpcMessage, apiCollection)
	return nodeMsg, nil
}

func (*GrpcChainParser) newChainMessage(api *spectypes.Api, requestedBlock int64, grpcMessage *rpcInterfaceMessages.GrpcMessage, apiCollection *spectypes.ApiCollection) *parsedMessage {
	nodeMsg := &parsedMessage{
		api:            api,
		msg:            grpcMessage, // setting the grpc message as a pointer so we can set descriptors for parsing
		requestedBlock: requestedBlock,
		apiCollection:  apiCollection,
	}
	return nodeMsg
}

// SetSpec sets the spec for the GrpcChainParser
func (apip *GrpcChainParser) SetSpec(spec spectypes.Spec) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return
	}

	// Add a read-write lock to ensure thread safety
	apip.rwLock.Lock()
	defer apip.rwLock.Unlock()

	// extract server and tagged apis from spec
	serverApis, taggedApis, apiCollections := getServiceApis(spec, spectypes.APIInterfaceGrpc)
	apip.BaseChainParser.Construct(spec, taggedApis, serverApis, apiCollections)
}

// DataReliabilityParams returns data reliability params from spec (spec.enabled and spec.dataReliabilityThreshold)
func (apip *GrpcChainParser) DataReliabilityParams() (enabled bool, dataReliabilityThreshold uint32) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return false, 0
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Return enabled and data reliability threshold from spec
	return apip.spec.Enabled, apip.spec.GetReliabilityThreshold()
}

// ChainBlockStats returns block stats from spec
// (spec.AllowedBlockLagForQosSync, spec.AverageBlockTime, spec.BlockDistanceForFinalizedData)
func (apip *GrpcChainParser) ChainBlockStats() (allowedBlockLagForQosSync int64, averageBlockTime time.Duration, blockDistanceForFinalizedData uint32, blocksInFinalizationProof uint32) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return 0, 0, 0, 0
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Convert average block time from int64 -> time.Duration
	averageBlockTime = time.Duration(apip.spec.AverageBlockTime) * time.Millisecond

	// Return allowedBlockLagForQosSync, averageBlockTime, blockDistanceForFinalizedData from spec
	return apip.spec.AllowedBlockLagForQosSync, averageBlockTime, apip.spec.BlockDistanceForFinalizedData, apip.spec.BlocksInFinalizationProof
}

type GrpcChainListener struct {
	endpoint    *lavasession.RPCEndpoint
	relaySender RelaySender
	logger      *metrics.RPCConsumerLogs
	chainParser *GrpcChainParser
}

func NewGrpcChainListener(
	ctx context.Context,
	listenEndpoint *lavasession.RPCEndpoint,
	relaySender RelaySender,
	rpcConsumerLogs *metrics.RPCConsumerLogs,
	chainParser ChainParser) (chainListener *GrpcChainListener) {
	// Create a new instance of GrpcChainListener
	chainListener = &GrpcChainListener{
		listenEndpoint,
		relaySender,
		rpcConsumerLogs,
		chainParser.(*GrpcChainParser),
	}
	return chainListener
}

// Serve http server for GrpcChainListener
func (apil *GrpcChainListener) Serve(ctx context.Context) {
	// Guard that the GrpcChainListener instance exists
	if apil == nil {
		return
	}

	utils.LavaFormatInfo("gRPC PortalStart")

	lis := GetListenerWithRetryGrpc("tcp", apil.endpoint.NetworkAddress)
	apiInterface := apil.endpoint.ApiInterface
	sendRelayCallback := func(ctx context.Context, method string, reqBody []byte) ([]byte, metadata.MD, error) {
		ctx = utils.WithUniqueIdentifier(ctx, utils.GenerateUniqueIdentifier())
		msgSeed := apil.logger.GetMessageSeed()
		metadataValues, _ := metadata.FromIncomingContext(ctx)
		grpcHeaders := convertToMetadataMapOfSlices(metadataValues)
		utils.LavaFormatInfo("GRPC Got Relay ", utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "method", Value: method})
		var relayReply *pairingtypes.RelayReply
		metricsData := metrics.NewRelayAnalytics("NoDappID", apil.endpoint.ChainID, apiInterface)
		relayReply, _, err := apil.relaySender.SendRelay(ctx, method, string(reqBody), "", "NoDappID", metricsData, grpcHeaders)
		go apil.logger.AddMetricForGrpc(metricsData, err, &metadataValues)

		if err != nil {
			errMasking := apil.logger.GetUniqueGuidResponseForError(err, msgSeed)
			apil.logger.LogRequestAndResponse("http in/out", true, method, string(reqBody), "", errMasking, msgSeed, err)
			return nil, nil, utils.LavaFormatError("Failed to SendRelay", fmt.Errorf(errMasking))
		}
		apil.logger.LogRequestAndResponse("http in/out", false, method, string(reqBody), "", "", msgSeed, nil)
		return relayReply.Data, convertRelayMetaDataToMDMetaData(relayReply.Metadata), nil
	}

	_, httpServer, err := grpcproxy.NewGRPCProxy(sendRelayCallback)
	if err != nil {
		utils.LavaFormatFatal("provider failure RegisterServer", err, utils.Attribute{Key: "listenAddr", Value: apil.endpoint.NetworkAddress})
	}

	// setup chain parser
	apil.chainParser.setupForConsumer(sendRelayCallback)

	utils.LavaFormatInfo("Server listening", utils.Attribute{Key: "Address", Value: lis.Addr()})

	if err := httpServer.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		utils.LavaFormatFatal("Portal failed to serve", err, utils.Attribute{Key: "Address", Value: lis.Addr()}, utils.Attribute{Key: "ChainID", Value: apil.endpoint.ChainID})
	}
}

type GrpcChainProxy struct {
	BaseChainProxy
	conn *chainproxy.GRPCConnector
}

func NewGrpcChainProxy(ctx context.Context, nConns uint, rpcProviderEndpoint *lavasession.RPCProviderEndpoint, averageBlockTime time.Duration, parser ChainParser) (ChainProxy, error) {
	if len(rpcProviderEndpoint.NodeUrls) == 0 {
		return nil, utils.LavaFormatError("rpcProviderEndpoint.NodeUrl list is empty missing node url", nil, utils.Attribute{Key: "chainID", Value: rpcProviderEndpoint.ChainID}, utils.Attribute{Key: "ApiInterface", Value: rpcProviderEndpoint.ApiInterface})
	}
	cp := &GrpcChainProxy{
		BaseChainProxy: BaseChainProxy{averageBlockTime: averageBlockTime},
	}
	nodeUrl := rpcProviderEndpoint.NodeUrls[0]
	nodeUrl.Url = strings.TrimSuffix(nodeUrl.Url, "/") // remove suffix if exists
	conn, err := chainproxy.NewGRPCConnector(ctx, nConns, nodeUrl)
	if err != nil {
		return nil, err
	}
	cp.conn = conn
	if cp.conn == nil {
		return nil, utils.LavaFormatError("g_conn == nil", nil)
	}

	err = parser.(*GrpcChainParser).setupForProvider(nodeUrl.Url)
	if err != nil {
		return nil, fmt.Errorf("grpc chain proxy: failed to setup parser: %w", err)
	}
	return cp, nil
}

func (cp *GrpcChainProxy) SendNodeMsg(ctx context.Context, ch chan interface{}, chainMessage ChainMessageForSend) (relayReply *pairingtypes.RelayReply, subscriptionID string, relayReplyServer *rpcclient.ClientSubscription, err error) {
	if ch != nil {
		return nil, "", nil, utils.LavaFormatError("Subscribe is not allowed on grpc", nil, utils.Attribute{Key: "GUID", Value: ctx})
	}
	conn, err := cp.conn.GetRpc(ctx, true)
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("grpc get connection failed ", err, utils.Attribute{Key: "GUID", Value: ctx})
	}
	defer cp.conn.ReturnRpc(conn)

	rpcInputMessage := chainMessage.GetRPCMessage()
	nodeMessage, ok := rpcInputMessage.(*rpcInterfaceMessages.GrpcMessage)
	if !ok {
		return nil, "", nil, utils.LavaFormatError("invalid message type in grpc failed to cast RPCInput from chainMessage", nil, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "rpcMessage", Value: rpcInputMessage})
	}
	if len(nodeMessage.Header) > 0 {
		metadataMap := make(map[string]string)
		for _, metaData := range nodeMessage.Header {
			metadataMap[metaData.Name] = metaData.Value
		}
		md := metadata.New(metadataMap)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	relayTimeout := common.LocalNodeTimePerCu(chainMessage.GetApi().ComputeUnits)
	// check if this API is hanging (waiting for block confirmation)
	if chainMessage.GetApi().Category.HangingApi {
		relayTimeout += cp.averageBlockTime
	}
	connectCtx, cancel := cp.NodeUrl.LowerContextTimeout(ctx, relayTimeout)
	defer cancel()

	// TODO: improve functionality, this is reading descriptors every send
	// improvement would be caching the descriptors, instead of fetching them.
	cl := grpcreflect.NewClient(ctx, reflectionpbo.NewServerReflectionClient(conn))
	descriptorSource := rpcInterfaceMessages.DescriptorSourceFromServer(cl)
	svc, methodName := rpcInterfaceMessages.ParseSymbol(nodeMessage.Path)
	var descriptor desc.Descriptor
	if descriptor, err = descriptorSource.FindSymbol(svc); err != nil {
		return nil, "", nil, utils.LavaFormatError("descriptorSource.FindSymbol", err, utils.Attribute{Key: "GUID", Value: ctx})
	}

	serviceDescriptor, ok := descriptor.(*desc.ServiceDescriptor)
	if !ok {
		return nil, "", nil, utils.LavaFormatError("serviceDescriptor, ok := descriptor.(*desc.ServiceDescriptor)", err, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "descriptor", Value: descriptor})
	}
	methodDescriptor := serviceDescriptor.FindMethodByName(methodName)
	if methodDescriptor == nil {
		return nil, "", nil, utils.LavaFormatError("serviceDescriptor.FindMethodByName returned nil", err, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "methodName", Value: methodName})
	}
	msgFactory := dynamic.NewMessageFactoryWithDefaults()

	var reader io.Reader
	msg := msgFactory.NewMessage(methodDescriptor.GetInputType())
	formatMessage := false
	if len(nodeMessage.Msg) > 0 {
		// guess if json or binary
		if nodeMessage.Msg[0] != '{' && nodeMessage.Msg[0] != '[' {
			msgLocal := msgFactory.NewMessage(methodDescriptor.GetInputType())
			err = proto.Unmarshal(nodeMessage.Msg, msgLocal)
			if err != nil {
				return nil, "", nil, err
			}
			jsonBytes, err := marshalJSON(msgLocal)
			if err != nil {
				return nil, "", nil, err
			}
			reader = bytes.NewReader(jsonBytes)
		} else {
			reader = bytes.NewReader(nodeMessage.Msg)
		}
		formatMessage = true
	}

	rp, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.FormatJSON, descriptorSource, reader, grpcurl.FormatOptions{
		EmitJSONDefaultFields: false,
		IncludeTextSeparator:  false,
		AllowUnknownFields:    true,
	})
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("Failed to create formatter", err, utils.Attribute{Key: "GUID", Value: ctx})
	}

	// used when parsing the grpc result
	nodeMessage.SetParsingData(methodDescriptor, formatter)

	if formatMessage {
		err = rp.Next(msg)
		if err != nil {
			return nil, "", nil, utils.LavaFormatError("rp.Next(msg) Failed", err, utils.Attribute{Key: "GUID", Value: ctx})
		}
	}

	var respHeaders metadata.MD
	response := msgFactory.NewMessage(methodDescriptor.GetOutputType())
	err = conn.Invoke(connectCtx, "/"+nodeMessage.Path, msg, response, grpc.Header(&respHeaders))

	if err != nil {
		if connectCtx.Err() == context.DeadlineExceeded {
			// Not an rpc error, return provider error without disclosing the endpoint address
			return nil, "", nil, utils.LavaFormatError("Provider Failed Sending Message", context.DeadlineExceeded)
		}
		return nil, "", nil, utils.LavaFormatError("Invoke Failed", err, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "Method", Value: nodeMessage.Path}, utils.Attribute{Key: "msg", Value: nodeMessage.Msg})
	}

	var respBytes []byte
	respBytes, err = proto.Marshal(response)
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("proto.Marshal(response) Failed", err, utils.Attribute{Key: "GUID", Value: ctx})
	}

	reply := &pairingtypes.RelayReply{
		Data:     respBytes,
		Metadata: convertToMetadataMapOfSlices(respHeaders),
	}
	return reply, "", nil, nil
}

func marshalJSON(msg proto.Message) ([]byte, error) {
	if dyn, ok := msg.(*dynamic.Message); ok {
		return dyn.MarshalJSON()
	}
	buf := new(bytes.Buffer)
	err := (&jsonpb.Marshaler{}).Marshal(buf, msg)
	return buf.Bytes(), err
}
