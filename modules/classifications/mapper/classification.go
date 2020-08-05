package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Classification struct {
	ID     types.ID     `json:"id" valid:"required~required field id missing"`
	Traits types.Traits `json:"traits" valid:"required~required field traits missing"`
}

var _ mappables.Classification = (*Classification)(nil)

func (classification Classification) GetID() types.ID { return classification.ID }

func (classification Classification) GetTraits() types.Traits { return classification.Traits }

//TODO
func (classification Classification) String() string { return "" }

func (classification Classification) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(classification)
}
func (classification Classification) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &classification)
	return classification
}

func classificationPrototype() traits.Mappable {
	return Classification{}
}

func NewClassification(classificationID types.ID, traits types.Traits) mappables.Classification {
	return Classification{
		ID:     classificationID,
		Traits: traits,
	}
}
