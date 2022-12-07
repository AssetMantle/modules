package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ documents.Identity = (*IdentityI)(nil)

func (x *IdentityI) GetExpiry() types.Height {
	return x.Impl.(documents.Identity).GetExpiry()
}

func (x *IdentityI) GetAuthentication() lists.DataList {
	return x.Impl.(documents.Identity).GetAuthentication()
}

func (x *IdentityI) IsProvisioned(address sdkTypes.AccAddress) bool {
	return x.Impl.(documents.Identity).IsProvisioned(address)
}

func (x *IdentityI) ProvisionAddress(address ...sdkTypes.AccAddress) documents.Identity {
	return x.Impl.(documents.Identity).ProvisionAddress(address...)
}

func (x *IdentityI) UnprovisionAddress(address ...sdkTypes.AccAddress) documents.Identity {
	return x.Impl.(documents.Identity).UnprovisionAddress(address...)
}

func (x *IdentityI) GenerateHashID() ids.HashID {
	return x.Impl.(documents.Identity).GenerateHashID()
}

func (x *IdentityI) GetClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Identity).GetClassificationID()
}

func (x *IdentityI) GetProperty(id ids.PropertyID) properties.Property {
	return x.Impl.(documents.Identity).GetProperty(id)
}

func (x *IdentityI) GetImmutables() qualified.Immutables {
	return x.Impl.(documents.Identity).GetImmutables()
}

func (x *IdentityI) GetMutables() qualified.Mutables {
	return x.Impl.(documents.Identity).GetMutables()
}

func (x *IdentityI) Mutate(property ...properties.Property) documents.Document {
	return x.Impl.(documents.Identity).Mutate(property...)
}
