package issue

import (
	"context"
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

func (msgServer msgServer) Issue(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)

	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	identityID := key.NewIdentityID(message.ClassificationID, immutableProperties)

	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(identityID))
	if identities.Get(key.FromID(identityID)) != nil {
		return nil, errors.EntityAlreadyExists
	}

	authenticationProperty := base.NewMetaProperty(base.NewID(properties.Authentication), base.NewMetaFact(base.NewListData(base.NewAccAddressData(message.To.AsSDKTypesAccAddress()))))
	updatedMutableMetaProperties := append(message.MutableMetaProperties.GetList(), authenticationProperty)

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(updatedMutableMetaProperties...)))

	if Error != nil {
		return nil, Error
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := msgServer.transactionKeeper.conformAuxiliary.GetKeeper().Help(ctx, conform.NewAuxiliaryRequest(message.ClassificationID, immutableProperties, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identities.Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
