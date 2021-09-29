package mutate

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Mutate(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&message.IdentityID))

	identity := identities.Get(key.FromID(&message.IdentityID))
	if identity == nil {
		return nil, errors.EntityNotFound
	}

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := msgServer.transactionKeeper.conformAuxiliary.GetKeeper().Help(ctx, conform.NewAuxiliaryRequest(identity.(mappables.InterIdentity).GetClassificationID(), nil, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := msgServer.transactionKeeper.maintainAuxiliary.GetKeeper().Help(ctx, maintain.NewAuxiliaryRequest(identity.(mappables.InterIdentity).GetClassificationID(), &message.FromID, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identities.Mutate(mappable.NewIdentity(identity.(mappables.InterIdentity).GetID(), identity.(mappables.InterIdentity).GetImmutableProperties(), identity.(mappables.InterIdentity).GetImmutableProperties().Mutate(mutableProperties.GetList()...)))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
