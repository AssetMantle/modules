package issue

import (
	"context"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Issue(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	fmt.Println("1111111111111")
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		fmt.Println("222222222", auxiliaryResponse.GetError())
		return nil, auxiliaryResponse.GetError()
	}

	fmt.Println("333333")
	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	fmt.Println("Printing immutable meta properties in Issue server", immutableMetaProperties)
	if Error != nil {
		return nil, Error
	}

	fmt.Println("555555")
	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	fmt.Println("Printing immutableProperties in Issue server", immutableProperties)
	identityID := key.NewIdentityID(&message.ClassificationID, immutableProperties)

	fmt.Println("Printing identityID in Issue server", identityID)
	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(identityID))

	if identities.Get(key.FromID(identityID)) != nil {
		return nil, errors.EntityAlreadyExists
	}

	fmt.Println("999999")
	authenticationProperty := base.NewMetaProperty(base.NewID(properties.Authentication), base.NewMetaFact(base.NewListData(base.NewAccAddressData(message.To.AsSDKTypesAccAddress()))))
	fmt.Println("Printing authenticationProperty in Issue server", authenticationProperty)
	updatedMutableMetaProperties := append(message.MutableMetaProperties.GetList(), authenticationProperty)

	fmt.Println("10101010")
	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(updatedMutableMetaProperties...)))

	fmt.Println("Printing mutableMetaProperties in Issue server", mutableMetaProperties)
	if Error != nil {
		return nil, Error
	}

	fmt.Println("+++++++")
	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	fmt.Println("Printing mutableProperties in Issue server", mutableProperties)
	if auxiliaryResponse := msgServer.transactionKeeper.conformAuxiliary.GetKeeper().Help(ctx, conform.NewAuxiliaryRequest(&message.ClassificationID, immutableProperties, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		fmt.Println("######")
		return nil, auxiliaryResponse.GetError()
	}

	fmt.Println("@@@@@@")
	identities.Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	fmt.Println("Printing identities in Issue server", identities)
	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
