package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

func IsProvisioned(context sdkTypes.Context, supplementAuxiliary helpers.Auxiliary, identity mappables.Identity, accAddress sdkTypes.AccAddress) (bool, error) {
	if metaPropertyList, err := supplement.GetMetaPropertiesFromResponse(supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.GetAuthentication()))); err != nil {
		return false, err
	} else {
		_, found := metaPropertyList.GetMetaProperty(constants.AuthenticationProperty).GetData().(data.ListData).Get().Search(baseData.NewAccAddressData(accAddress))
		return found, nil
	}
}

func ProvisionAddress(context sdkTypes.Context, supplementAuxiliary helpers.Auxiliary, identity mappables.Identity, accAddress sdkTypes.AccAddress) (mappables.Identity, error) {
	if metaPropertyList, err := supplement.GetMetaPropertiesFromResponse(supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.GetAuthentication()))); err != nil {
		return identity, err
	} else {
		identity.Mutate(base.NewProperty(constants.AuthenticationProperty, baseData.NewListData(metaPropertyList.GetMetaProperty(constants.AuthenticationProperty).GetData().(data.ListData).Get().Add(baseData.NewAccAddressData(accAddress)))))
		return identity, nil
	}
}

func UnprovisionAddress(context sdkTypes.Context, supplementAuxiliary helpers.Auxiliary, identity mappables.Identity, accAddress sdkTypes.AccAddress) (mappables.Identity, error) {
	if metaPropertyList, err := supplement.GetMetaPropertiesFromResponse(supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.GetAuthentication()))); err != nil {
		return identity, err
	} else {
		identity.Mutate(base.NewProperty(constants.AuthenticationProperty, baseData.NewListData(metaPropertyList.GetMetaProperty(constants.AuthenticationProperty).GetData().(data.ListData).Get().Delete(baseData.NewAccAddressData(accAddress)))))
		return identity, nil
	}
}
