package nub

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Nub(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	context := sdkTypes.UnwrapSDKContext(goCtx)
	nubIDProperty := base.NewMetaProperty(base.NewID(properties.NubID), base.NewMetaFact(base.NewIDData(message.NubID)))

	immutableProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(nubIDProperty)))
	if Error != nil {
		return nil, Error
	}

	authenticationProperty := base.NewMetaProperty(base.NewID(properties.Authentication), base.NewMetaFact(base.NewListData(base.NewAccAddressData(message.From.AsSDKTypesAccAddress()))))

	mutableProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(authenticationProperty)))
	if Error != nil {
		return nil, Error
	}

	classificationID, Error := define.GetClassificationIDFromResponse(msgServer.transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base.NewID(properties.NubID), base.NewFact(base.NewIDData(base.NewID(""))))), base.NewProperties(base.NewProperty(base.NewID(properties.Authentication), base.NewFact(base.NewAccAddressData(nil).ZeroValue()))))))
	if classificationID == nil && Error != nil {
		return nil, Error
	}

	identityID := key.NewIdentityID(classificationID, immutableProperties)

	identities := msgServer.transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))
	if identities.Get(key.FromID(identityID)) != nil {
		return nil, errors.EntityAlreadyExists
	}

	identities.Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
