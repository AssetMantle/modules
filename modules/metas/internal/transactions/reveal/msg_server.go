package reveal

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Reveal(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	metaID := key.GenerateMetaID(message.MetaFact.GetData())
	metas := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(metaID))

	meta := metas.Get(key.FromID(metaID))
	if meta != nil {
		return nil, errors.EntityAlreadyExists
	}

	if message.MetaFact.GetHashID().Compare(base.NewID("")) != 0 {
		metas.Add(mappable.NewMeta(message.MetaFact.GetData()))
	}
	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
