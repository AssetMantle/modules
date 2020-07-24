package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type splits struct {
	ID   types.ID          `json:"id" valid:"required~required field id missing"`
	List []mappables.Split `json:"list" valid:"required~required field list missing"`

	mapper  utilities.Mapper `json:"mapper" valid:"required~required field mapper missing"`
	context sdkTypes.Context `json:"context" valid:"required~required field context missing"`
}

var _ mappers.Splits = (*splits)(nil)

func (splits splits) GetID() types.ID { return splits.ID }
func (splits splits) Get(id types.ID) mappables.Split {
	splitID := splitIDFromInterface(id)
	for _, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(splitID) == 0 {
			return oldSplit
		}
	}
	return nil
}
func (splits splits) GetList() []mappables.Split {
	return splits.List
}

func (splits splits) Fetch(id types.ID) mappers.Splits {
	var splitList []mappables.Split
	splitsID := splitIDFromInterface(id)
	if len(splitsID.OwnableID.Bytes()) > 0 {
		mappable := splits.mapper.Read(splits.context, splitsID)
		if mappable != nil {
			splitList = append(splitList, mappable.(split))
		}
	} else {
		appendSplitList := func(mappable traits.Mappable) bool {
			splitList = append(splitList, mappable.(split))
			return false
		}
		splits.mapper.Iterate(splits.context, splitsID, appendSplitList)
	}
	splits.ID, splits.List = id, splitList
	return splits
}
func (splits splits) Add(split mappables.Split) mappers.Splits {
	splits.ID = readSplitID("")
	splits.mapper.Create(splits.context, split)
	splits.List = append(splits.List, split)
	return splits
}
func (splits splits) Remove(split mappables.Split) mappers.Splits {
	splits.mapper.Delete(splits.context, split.GetID())
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List = append(splits.List[:i], splits.List[i+1:]...)
			break
		}
	}
	return splits
}
func (splits splits) Mutate(split mappables.Split) mappers.Splits {
	splits.mapper.Update(splits.context, split)
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List[i] = split
			break
		}
	}
	return splits
}

func NewSplits(mapper utilities.Mapper, context sdkTypes.Context) mappers.Splits {
	return splits{
		ID:      readSplitID(""),
		List:    []mappables.Split{},
		mapper:  mapper,
		context: context,
	}
}
