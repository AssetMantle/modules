// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type fields struct {
	From string
	To   string
	*baseIDs.IdentityID
}

func CreateTestInputForMessage(t *testing.T) (*baseIDs.IdentityID, sdkTypes.AccAddress, sdkTypes.AccAddress, sdkTypes.Msg) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	toAccAddress, err := sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, toAccAddress, testIdentityID)

	return testIdentityID.(*baseIDs.IdentityID), fromAccAddress, toAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, testMessage := CreateTestInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), toAccAddress.String(), testIdentityID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messageFromInterface(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{
		{"+ve", &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSignBytes(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, testMessage := CreateTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type args struct {
		from sdkTypes.AccAddress
		to   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, toAccAddress, testIdentityID}, &Message{fromAccAddress.String(), toAccAddress.String(), testIdentityID}},
		{"-ve", args{}, &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.to, tt.args.IdentityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
