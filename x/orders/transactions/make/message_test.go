// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/types"
	"github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	expiresInHeight = base.NewHeight(60).(*base.Height)
	testMessage     = NewMessage(fromAccAddress, testFromID, testClassificationID, testFromID, makerAssetID, takerAssetID, expiresInHeight, makerSplit, takerSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties, mutableMetaProperties, mutableProperties)
)

type fields struct {
	From                    string
	FromID                  *baseIDs.IdentityID
	ClassificationID        *baseIDs.ClassificationID
	TakerID                 *baseIDs.IdentityID
	MakerAssetID            *baseIDs.AssetID
	TakerAssetID            *baseIDs.AssetID
	ExpiresIn               *base.Height
	MakerSplit              sdkTypes.Int
	TakerSplit              sdkTypes.Int
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

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerAssetID.(*baseIDs.AssetID), takerAssetID.(*baseIDs.AssetID), expiresInHeight, makerSplit, takerSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerAssetID:            tt.fields.MakerAssetID,
				TakerAssetID:            tt.fields.TakerAssetID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerSplit:              tt.fields.MakerSplit.String(),
				TakerSplit:              tt.fields.TakerSplit.String(),
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

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerAssetID.(*baseIDs.AssetID), takerAssetID.(*baseIDs.AssetID), expiresInHeight, makerSplit, takerSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, false},
		{"-ve for nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerAssetID:            tt.fields.MakerAssetID,
				TakerAssetID:            tt.fields.TakerAssetID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerSplit:              tt.fields.MakerSplit.String(),
				TakerSplit:              tt.fields.TakerSplit.String(),
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
	type args struct {
		from                    sdkTypes.AccAddress
		fromID                  ids.IdentityID
		classificationID        ids.ClassificationID
		takerID                 ids.IdentityID
		makerAssetID            ids.AssetID
		takerAssetID            ids.AssetID
		expiresIn               types.Height
		makerSplit              sdkTypes.Int
		takerSplit              sdkTypes.Int
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
		{"+ve", args{fromAccAddress, testFromID, testClassificationID, testFromID, makerAssetID, takerAssetID, expiresInHeight, makerSplit, takerSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, &Message{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerAssetID.(*baseIDs.AssetID), takerAssetID.(*baseIDs.AssetID), expiresInHeight, makerSplit.String(), takerSplit.String(), immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.classificationID, tt.args.takerID, tt.args.makerAssetID, tt.args.takerAssetID, tt.args.expiresIn, tt.args.makerSplit, tt.args.takerSplit, tt.args.immutableMetaProperties.(*baseLists.PropertyList), tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
