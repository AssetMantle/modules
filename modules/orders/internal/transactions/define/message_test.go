package define

import (
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func CreateTestInputForMessage(t *testing.T) (ids.ID, sdkTypes.AccAddress, lists.MetaPropertyList, lists.PropertyList, lists.MetaPropertyList, lists.PropertyList, sdkTypes.Msg) {
	testFromID := baseIDs.NewID("fromID")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaProperties, err := utilities.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)

	return testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, testMessage := CreateTestInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
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
		// TODO: Add test cases.
		{"+ve", message{}},
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
	//testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, testMessage := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		//{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.ID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testFromID, fromAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _ := CreateTestInputForMessage(t)
	type args struct {
		from                    sdkTypes.AccAddress
		fromID                  ids.ID
		immutableMetaProperties lists.MetaPropertyList
		immutableProperties     lists.PropertyList
		mutableMetaProperties   lists.MetaPropertyList
		mutableProperties       lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, message{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
