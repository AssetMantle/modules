package reveal

import (
	"context"
	"fmt"
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
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(ctx, "Printing context")
	fmt.Println("")
	fmt.Println("")
	metaID := key.GenerateMetaID(msg.MetaFact.GetData())
	metas := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(metaID))

	meta := metas.Get(key.FromID(metaID))
	if meta != nil {
		return nil, errors.EntityAlreadyExists
	}
	fmt.Println(meta, "Printing Meta in ms_server ----------")
	if msg.MetaFact.GetHashID().Compare(base.NewID("")) != 0 {
		metas.Add(mappable.NewMeta(msg.MetaFact.GetData()))
	}
	fmt.Println(msg, "Printing ,msg in ms_server ++++++++++++")
	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
