// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

type fields struct {
	From                    string
	FromID                  *baseIDs.IdentityID
	ClassificationID        *baseIDs.ClassificationID
	ImmutableMetaProperties *baseLists.PropertyList
	ImmutableProperties     *baseLists.PropertyList
	MutableMetaProperties   *baseLists.PropertyList
	MutableProperties       *baseLists.PropertyList
}

func createTestInput(t *testing.T) (*baseIDs.IdentityID, *baseIDs.ClassificationID, string, sdkTypes.AccAddress, string, sdkTypes.AccAddress, *baseLists.PropertyList, *baseLists.PropertyList, *baseLists.PropertyList, *baseLists.PropertyList) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)

	const fromAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	const toAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	var immutableMetaProperties lists.PropertyList
	immutableMetaProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.PropertyList
	mutableMetaProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	return testFromID.(*baseIDs.IdentityID), testClassificationID.(*baseIDs.ClassificationID), fromAddress, fromAccAddress, toAddress, toAccAddress, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)
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
	testFromID, testClassificationID, _, fromAccAddress, _, _, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, _, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
		{"-ve with nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type args struct {
		from                    sdkTypes.AccAddress
		to                      sdkTypes.AccAddress
		fromID                  ids.IdentityID
		classificationID        ids.ClassificationID
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
		{"+ve", args{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, NewMessage(fromAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)},
		{"-ve with nil", args{}, &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
