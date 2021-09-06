package qualified

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Document struct {
	ID               types.ID `json:"id" valid:"required~required field id missing"`
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	baseTraits.HasImmutables
	baseTraits.HasMutables //nolint:govet
}

var _ mappables.Document = (*Document)(nil)

func (document Document) GetID() types.ID               { return document.ID }
func (document Document) GetClassificationID() types.ID { return document.ClassificationID }
func (document Document) GetProperty(id types.ID) types.Property {
	if property := document.HasImmutables.GetImmutableProperties().Get(id); property != nil {
		return property
	} else if property := document.HasMutables.GetMutableProperties().Get(id); property != nil {
		return property
	} else {
		return nil
	}
}
func (document Document) GetKey() helpers.Key {
	panic("implement me")
}
func (document Document) RegisterCodec(codec *codec.Codec) {
	panic("implement me")
}
