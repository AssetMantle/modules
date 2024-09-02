// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

type fields struct {
	From                    string
	FromID                  *baseIDs.IdentityID
	ImmutableMetaProperties *baseLists.PropertyList
	ImmutableProperties     *baseLists.PropertyList
	MutableMetaProperties   *baseLists.PropertyList
	MutableProperties       *baseLists.PropertyList
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

func Test_message_GetSigners(t *testing.T) {
	fromID, err := baseIDs.PrototypeIdentityID().FromString("CBepOLnJFnKO9NEyZlSv7r80nKNZFFXRqHfnsObZ_KU=")
	testFromID := fromID.(*baseIDs.IdentityID)
	require.NoError(t, err)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableMetaProperties := immutableMetaPropertiesInterface.(*baseLists.PropertyList)

	immutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	immutableProperties := immutablePropertiesInterface.(*baseLists.PropertyList)

	mutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableMetaProperties := mutableMetaPropertiesInterface.(*baseLists.PropertyList)

	mutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	mutableProperties := mutablePropertiesInterface.(*baseLists.PropertyList)

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{

		{"+ve", fields{fromAccAddress.String(), testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

func Test_message_ValidateBasic(t *testing.T) {
	fromID, err := baseIDs.PrototypeIdentityID().FromString("CBepOLnJFnKO9NEyZlSv7r80nKNZFFXRqHfnsObZ_KU=")
	testFromID := fromID.(*baseIDs.IdentityID)
	require.NoError(t, err)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableMetaProperties := immutableMetaPropertiesInterface.(*baseLists.PropertyList)

	immutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	immutableProperties := immutablePropertiesInterface.(*baseLists.PropertyList)

	mutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableMetaProperties := mutableMetaPropertiesInterface.(*baseLists.PropertyList)

	mutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	mutableProperties := mutablePropertiesInterface.(*baseLists.PropertyList)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{"+ve", fields{fromAccAddress.String(), testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

func Test_NewMessage(t *testing.T) {
	fromID, err := baseIDs.PrototypeIdentityID().FromString("CBepOLnJFnKO9NEyZlSv7r80nKNZFFXRqHfnsObZ_KU=")
	testFromID := fromID.(*baseIDs.IdentityID)
	require.NoError(t, err)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableMetaProperties := immutableMetaPropertiesInterface.(*baseLists.PropertyList)

	immutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	immutableProperties := immutablePropertiesInterface.(*baseLists.PropertyList)

	mutableMetaPropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableMetaProperties := mutableMetaPropertiesInterface.(*baseLists.PropertyList)

	mutablePropertiesInterface, err := baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	mutableProperties := mutablePropertiesInterface.(*baseLists.PropertyList)

	type args struct {
		from                    sdkTypes.AccAddress
		fromID                  ids.IdentityID
		immutableMetaProperties lists.PropertyList
		immutableProperties     lists.PropertyList
		mutableMetaProperties   lists.PropertyList
		mutableProperties       lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{

		{"+ve", args{fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, &Message{fromAccAddress.String(), testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
		{"+ve with nil", args{}, &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
