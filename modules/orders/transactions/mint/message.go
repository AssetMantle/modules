package mint

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

//TODO make private
type Message struct {
	From                sdkTypes.AccAddress
	MaintainersID       schema.ID
	ClassificationID    schema.ID
	Properties          schema.Properties
	Lock                schema.Height
	Burn                schema.Height
	TakerAddress        sdkTypes.AccAddress
	SenderAddress       sdkTypes.AccAddress
	FeeRecipientAddress sdkTypes.AccAddress
	MakerAssetAmount    sdkTypes.Dec
	MakerAssetData      schema.ID
	MakerFee            sdkTypes.Dec
	MakerFeeAssetData   schema.ID
	TakerAssetAmount    sdkTypes.Dec
	TakerAssetData      schema.ID
	TakerFee            sdkTypes.Dec
	TakerFeeAssetData   schema.ID
	ExpirationTime      schema.Height
	Salt                schema.Height
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

func (message Message) GenerateHash() schema.ID {
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
	return schema.NewID(sha)
}

func newMessage(from sdkTypes.AccAddress, maintainersID schema.ID, classificationID schema.ID, properties schema.Properties, lock schema.Height, burn schema.Height,
	takerAddress sdkTypes.AccAddress, senderAddress sdkTypes.AccAddress, feeRecipientAddress sdkTypes.AccAddress,
	makerAssetAmount sdkTypes.Dec, makerAssetData schema.ID, makerFee sdkTypes.Dec, makerFeeAssetData schema.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData schema.ID, takerFee sdkTypes.Dec, takerFeeAssetData schema.ID,
	expirationTime schema.Height, salt schema.Height) sdkTypes.Msg {
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
