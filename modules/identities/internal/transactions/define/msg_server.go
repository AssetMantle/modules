package define

import (
	"context"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Define(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identity := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&message.FromID)).Get(key.FromID(&message.FromID)).(mappables.InterIdentity)
	if identity == nil {
		return nil, errors.EntityNotFound
	}

	if !identity.IsProvisioned(message.From.AsSDKTypesAccAddress()) {
		return nil, errors.NotAuthorized
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	classificationID, Error := define.GetClassificationIDFromResponse(msgServer.transactionKeeper.defineAuxiliary.GetKeeper().Help(ctx, define.NewAuxiliaryRequest(immutableProperties, mutableProperties)))
	fmt.Println(classificationID, "Printing Classification ID")
	if Error != nil {
		return nil, Error
	}

	if auxiliaryResponse := msgServer.transactionKeeper.superAuxiliary.GetKeeper().Help(ctx, super.NewAuxiliaryRequest(classificationID, &message.FromID, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
