package make

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO make private
type Message struct {
	From                sdkTypes.AccAddress
	MaintainersID       types.ID
	ClassificationID    types.ID
	Properties          types.Properties
	Lock                types.Height
	Burn                types.Height
	TakerAddress        sdkTypes.AccAddress
	SenderAddress       sdkTypes.AccAddress
	FeeRecipientAddress sdkTypes.AccAddress
	MakerAssetAmount    sdkTypes.Dec
	MakerAssetData      types.ID
	MakerFee            sdkTypes.Dec
	MakerFeeAssetData   types.ID
	TakerAssetAmount    sdkTypes.Dec
	TakerAssetData      types.ID
	TakerFee            sdkTypes.Dec
	TakerFeeAssetData   types.ID
	ExpirationTime      types.Height
	Salt                types.Height
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return Transaction.GetModuleName() }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}

func (message Message) GenerateHash() types.ID {
	hasher := sha512.New()
	//bz := []byte(message.MakerAddress.String() + message.TakerAddress.String() + message.SenderAddress.String() +
	//	message.FeeRecipientAddress.String() + message.MakerAssetAmount.String() + message.MakerAssetData.String() +
	//	message.MakerFee.String() + message.MakerFeeAssetData.String() + message.TakerAssetAmount.String() +
	//	message.TakerAssetData.String() + message.TakerFee.String() + message.TakerFeeAssetData.String() +
	//	string(message.ExpirationTime.Get()) + string(message.Salt.Get()))
	bz := []byte(message.From.String() + message.TakerAddress.String() +
		message.MakerAssetAmount.String() + message.MakerAssetData.String() +
		message.TakerAssetAmount.String() + message.TakerAssetData.String() + string(message.Salt.Get()))
	hasher.Write(bz)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return base.NewID(sha)
}

func newMessage(from sdkTypes.AccAddress, maintainersID types.ID, classificationID types.ID, properties types.Properties, lock types.Height, burn types.Height,
	takerAddress sdkTypes.AccAddress, senderAddress sdkTypes.AccAddress, feeRecipientAddress sdkTypes.AccAddress,
	makerAssetAmount sdkTypes.Dec, makerAssetData types.ID, makerFee sdkTypes.Dec, makerFeeAssetData types.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData types.ID, takerFee sdkTypes.Dec, takerFeeAssetData types.ID,
	expirationTime types.Height, salt types.Height) sdkTypes.Msg {
	return Message{
		From:                from,
		MaintainersID:       maintainersID,
		ClassificationID:    classificationID,
		Properties:          properties,
		Lock:                lock,
		Burn:                burn,
		TakerAddress:        takerAddress,
		SenderAddress:       senderAddress,
		FeeRecipientAddress: feeRecipientAddress,
		MakerAssetAmount:    makerAssetAmount,
		MakerAssetData:      makerAssetData,
		MakerFee:            makerFee,
		MakerFeeAssetData:   makerFeeAssetData,
		TakerAssetAmount:    takerAssetAmount,
		TakerAssetData:      takerAssetData,
		TakerFee:            takerFee,
		TakerFeeAssetData:   takerFeeAssetData,
		ExpirationTime:      expirationTime,
		Salt:                salt,
	}
}
