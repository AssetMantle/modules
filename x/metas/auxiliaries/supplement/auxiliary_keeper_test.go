// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/constants"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/record"
	baseData "github.com/AssetMantle/schema/go/data/base"
	dataConstants "github.com/AssetMantle/schema/go/data/constants"
	"github.com/AssetMantle/schema/go/errors"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

type mockAuxiliaryRequest struct {
	mock.Mock
}

func (*mockAuxiliaryRequest) Validate() error {
	return nil
}

const (
	ChainID = "testChain"
)

var (
	randomMetaPropertyGenerator = func() properties.MetaProperty {
		return baseProperties.NewMetaProperty(baseIDs.NewStringID(strconv.Itoa(rand.Intn(99999999999999999))), baseData.NewStringData(strconv.Itoa(rand.Intn(rand.Intn(99999999999999999)))))
	}

	randomPropertiesGenerator = func(n int) (unScrubbed []properties.Property, scrubbed []properties.Property) {
		unScrubbed = make([]properties.Property, n)
		scrubbed = make([]properties.Property, n)
		for i := 0; i < n; i++ {
			unScrubbed[i] = randomMetaPropertyGenerator()
			scrubbed[i] = unScrubbed[i].(properties.MetaProperty).ScrubData().ToAnyProperty().(*baseProperties.AnyProperty)
		}

		return unScrubbed, scrubbed
	}

	testPropertiesCount                              = 100
	testUnScrubbedProperties, testScrubbedProperties = randomPropertiesGenerator(testPropertiesCount)
	randomIndex                                      = rand.Intn(testPropertiesCount)

	invalidMetaProperty = &baseProperties.MetaProperty{
		ID:   baseIDs.NewPropertyID(baseIDs.NewStringID("invalid"), baseIDs.NewStringID("invalid")).(*baseIDs.PropertyID),
		Data: baseData.NewStringData("invalid").ToAnyData().(*baseData.AnyData),
	}

	moduleStoreKey  = sdkTypes.NewKVStoreKey(constants.ModuleName)
	AuxiliaryKeeper = auxiliaryKeeper{mapper.Prototype().Initialize(moduleStoreKey)}

	setContext = func() sdkTypes.Context {
		memDB := tendermintDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB)
		commitMultiStore.MountStoreWithDB(moduleStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	}

	Context = setContext()

	setMeta = func() error {
		for _, property := range testUnScrubbedProperties {
			AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
				Add(record.NewRecord(property.(properties.MetaProperty).GetData()))
		}
		return nil
	}

	_ = setMeta()
)

func Test_auxiliaryKeeper_Help(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		request helpers.AuxiliaryRequest
		want    helpers.AuxiliaryResponse
		wantErr errors.Error
	}{
		{
			"valid request",
			func() {},
			NewAuxiliaryRequest(testScrubbedProperties...),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties...)),
			nil,
		},
		{
			"empty request",
			func() {},
			NewAuxiliaryRequest(),
			NewAuxiliaryResponse(base.NewPropertyList()),
			nil,
		},
		{
			"invalid request",
			func() {},
			NewAuxiliaryRequest(invalidMetaProperty),
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid request type",
			func() {},
			&mockAuxiliaryRequest{},
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"nil properties",
			func() {},
			NewAuxiliaryRequest(nil),
			NewAuxiliaryResponse(base.NewPropertyList()),
			nil,
		},
		{
			"one property",
			func() {},
			NewAuxiliaryRequest(testScrubbedProperties[randomIndex]),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties[randomIndex])),
			nil,
		},
		{
			"two properties",
			func() {},
			NewAuxiliaryRequest(testScrubbedProperties[0], testScrubbedProperties[1]),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties[0], testUnScrubbedProperties[1])),
			nil,
		},
		{
			"prototype property",
			func() {},
			NewAuxiliaryRequest(baseProperties.PrototypeMetaProperty().ScrubData()),
			NewAuxiliaryResponse(base.NewPropertyList()),
			nil,
		},
		{
			"nil with properties",
			func() {},
			NewAuxiliaryRequest(nil, testScrubbedProperties[0], nil, testScrubbedProperties[1], nil),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties[0], testUnScrubbedProperties[1])),
			nil,
		},
		{
			"prototype property with properties",
			func() {},
			NewAuxiliaryRequest(baseProperties.PrototypeMetaProperty().ScrubData(), testScrubbedProperties[0], baseProperties.PrototypeMetaProperty().ScrubData(), testScrubbedProperties[1]),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties[0], testUnScrubbedProperties[1])),
			nil,
		},
		{
			"property not present",
			func() {},
			NewAuxiliaryRequest(randomMetaPropertyGenerator().ScrubData()),
			NewAuxiliaryResponse(base.NewPropertyList()),
			nil,
		},
		{
			"meta property",
			func() {},
			NewAuxiliaryRequest(testUnScrubbedProperties[randomIndex]),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties[randomIndex])),
			nil,
		},
		{
			"zero value number",
			func() {},
			NewAuxiliaryRequest(&baseProperties.MesaProperty{
				ID: baseIDs.NewPropertyID(baseIDs.NewStringID("zero"), dataConstants.NumberDataTypeID).(*baseIDs.PropertyID),
				DataID: &baseIDs.DataID{
					TypeID: dataConstants.NumberDataTypeID.(*baseIDs.StringID),
					HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
				},
			}),
			NewAuxiliaryResponse(base.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("zero"), baseData.PrototypeNumberData().ZeroValue()))),
			nil,
		},
		{
			"zero value string",
			func() {},
			NewAuxiliaryRequest(&baseProperties.MesaProperty{
				ID: baseIDs.NewPropertyID(baseIDs.NewStringID("zero"), dataConstants.StringDataTypeID).(*baseIDs.PropertyID),
				DataID: &baseIDs.DataID{
					TypeID: dataConstants.StringDataTypeID.(*baseIDs.StringID),
					HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
				},
			}),
			NewAuxiliaryResponse(base.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("zero"), baseData.PrototypeStringData().ZeroValue()))),
			nil,
		},
		{
			"zero value boolean",
			func() {},
			NewAuxiliaryRequest(&baseProperties.MesaProperty{
				ID: baseIDs.NewPropertyID(baseIDs.NewStringID("zero"), dataConstants.BooleanDataTypeID).(*baseIDs.PropertyID),
				DataID: &baseIDs.DataID{
					TypeID: dataConstants.BooleanDataTypeID.(*baseIDs.StringID),
					HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
				},
			},
			),
			NewAuxiliaryResponse(base.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("zero"), baseData.PrototypeBooleanData().ZeroValue()))),
			nil,
		},
		{
			"zero value list",
			func() {},
			NewAuxiliaryRequest(&baseProperties.MesaProperty{
				ID: baseIDs.NewPropertyID(baseIDs.NewStringID("zero"), dataConstants.ListDataTypeID).(*baseIDs.PropertyID),
				DataID: &baseIDs.DataID{
					TypeID: dataConstants.ListDataTypeID.(*baseIDs.StringID),
					HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
				},
			},
			),
			NewAuxiliaryResponse(base.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("zero"), baseData.PrototypeListData().ZeroValue()))),
			nil,
		},
		{
			"zero value linked",
			func() {},
			NewAuxiliaryRequest(&baseProperties.MesaProperty{
				ID: baseIDs.NewPropertyID(baseIDs.NewStringID("zero"), dataConstants.LinkedDataTypeID).(*baseIDs.PropertyID),
				DataID: &baseIDs.DataID{
					TypeID: dataConstants.LinkedDataTypeID.(*baseIDs.StringID),
					HashID: baseIDs.PrototypeHashID().(*baseIDs.HashID),
				},
			},
			),
			NewAuxiliaryResponse(base.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("zero"), baseData.PrototypeLinkedData().ZeroValue()))),
			nil,
		},
		{
			"very large number of properties",
			func() {
				for i := 0; i < 100000; i++ {
					AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
						Add(record.NewRecord(randomMetaPropertyGenerator().GetData()))
				}
			},
			NewAuxiliaryRequest(testScrubbedProperties...),
			NewAuxiliaryResponse(base.NewPropertyList(testUnScrubbedProperties...)),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := AuxiliaryKeeper.Help(sdkTypes.WrapSDKContext(Context), tt.request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n got: \n %v \n want: \n %v", got, tt.want)
			}

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("\n want error: \n %v \n got error: \n %v", err, tt.wantErr)
			}
		})
	}
}
