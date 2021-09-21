package provision

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Provision(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	identityID := message.IdentityID
	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&identityID))

	identity := identities.Get(key.FromID(&identityID))
	if identity == nil {
		return nil, errors.EntityNotFound
	}

	if !identity.(mappables.InterIdentity).IsProvisioned(message.From.AsSDKTypesAccAddress()) {
		return nil, errors.NotAuthorized
	}

	if !identity.(mappables.InterIdentity).IsProvisioned(message.To.AsSDKTypesAccAddress()) {
		return nil, errors.EntityAlreadyExists
	}

	identityAuthenticationProperty := identity.(mappables.InterIdentity).GetAuthentication()
	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(identityAuthenticationProperty)))
	if Error != nil {
		return nil, Error
	}
	listData, Error := metaProperties.GetList()[0].GetMetaFact().GetData().AsListData()
	if Error != nil {
		return nil, Error
	}
	listData.Add(base.NewAccAddressData(message.To.AsSDKTypesAccAddress()))
	authenticationProperty := base.NewMetaProperty(base.NewID(properties.Authentication), base.NewMetaFact(listData))
	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(authenticationProperty)))
	if Error != nil {
		return nil, Error
	}
	modifiedMutableProperties := identity.(mappables.InterIdentity).GetMutableProperties().Mutate(mutableMetaProperties.GetList()...)
	identities.Mutate(mappable.NewIdentity(&identityID, identity.(mappables.InterIdentity).GetImmutableProperties(), modifiedMutableProperties))
	//identities.Mutate(identity.(mappables.InterIdentity).ProvisionAddress(message.To.AsSDKTypesAccAddress()))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
