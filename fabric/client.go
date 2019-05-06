package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

type sdkUtil struct {
	Sdk         *fabsdk.FabricSDK
	Client      *channel.Client
	*SdkConfig
}

var sdk *sdkUtil

func ObtainSdkUtil() *sdkUtil{
	return sdk
}

func NewSdkUtil(sdkConfig *SdkConfig) *sdkUtil{
	sdk = &sdkUtil {
		SdkConfig   : sdkConfig,
	}
	return sdk
}

func (sc *sdkUtil) Start() {
	configProvider := config.FromFile(sc.ConfigFile)
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("create sdk fail: %s\n", err.Error())
	}

	mspCli, err := msp.New(sdk.Context(), msp.WithOrg(sc.Org))
	if err != nil {
		log.Fatalf("create msp client fail: %s\n", err.Error())
	}

	_, err = mspCli.GetSigningIdentity(sc.User)
	if err != nil {
		log.Fatalf("sign identity fail: %s\n", err.Error())
	} else {
		log.Println("msp identity is successful")
	}

	channelProvider := sdk.ChannelContext(sc.Channel,
		fabsdk.WithUser(sc.User),
		fabsdk.WithOrg(sc.Org))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		log.Fatalf("channel client create fail: %s\n", err.Error())
	}
	sc.Sdk = sdk
	sc.Client = channelClient
}

func (sc *sdkUtil) Obtain(fcn string, args [][]byte) interface{}{
	request := channel.Request{
		ChaincodeID: sc.ChainCodeId,
		Fcn:         fcn,
		Args:        args,
	}
	response, err := sc.Client.Query(request)
	if err != nil {
		log.Fatalf("obtain fail: %s", err.Error())
		return nil
	}
	log.Printf("obtain response is %s\n", response.Payload)
	return string(response.Payload)
}

func (sc *sdkUtil) Invoke(fcn string, args [][]byte) interface{}{
	request := channel.Request{
		ChaincodeID: sc.ChainCodeId,
		Fcn:         fcn,
		Args:        args,
	}
	response, err := sc.Client.Execute(request)
	if err != nil {
		log.Fatalf("Invoke fail: %s", err.Error())
		return nil
	}
	log.Printf("Invoke response is %s\n", response.Payload)
	return string(response.Payload)
}
