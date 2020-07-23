package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type splits struct {
	ID   types.ID         `json:"id" valid:"required~Enter the ID"`
	List []entities.Split `json:"list" valid:"required~Enter the List"`

	mapper  splitsMapper     `json:"mapper" valid:"required~Enter the Mapper"`
	context sdkTypes.Context `json:"context" valid:"required~Enter the Context"`
}

var _ mappers.Splits = (*splits)(nil)

func (splits splits) GetID() types.ID { return splits.ID }
func (splits splits) Get(id types.ID) entities.Split {
	splitID := splitIDFromInterface(id)
	for _, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(splitID) == 0 {
			return oldSplit
		}
	}
	return nil
}
func (splits splits) GetList() []entities.Split {
	return splits.List
}

func (splits splits) Fetch(id types.ID) mappers.Splits {
	var splitList []entities.Split
	splitsID := splitIDFromInterface(id)
	if len(splitsID.OwnableID.Bytes()) > 0 {
		split := splits.mapper.read(splits.context, splitsID)
		if split != nil {
			splitList = append(splitList, split)
		}
	} else {
		appendSplitList := func(split entities.Split) bool {
			splitList = append(splitList, split)
			return false
		}
		splits.mapper.iterate(splits.context, splitsID, appendSplitList)
	}
	splits.ID, splits.List = id, splitList
	return splits
}
func (splits splits) Add(split entities.Split) mappers.Splits {
	splits.ID = readSplitID("")
	splits.mapper.create(splits.context, split)
	splits.List = append(splits.List, split)
	return splits
}
func (splits splits) Remove(split entities.Split) mappers.Splits {
	splits.mapper.delete(splits.context, split.GetID())
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List = append(splits.List[:i], splits.List[i+1:]...)
			break
		}
	}
	return splits
}
func (splits splits) Mutate(split entities.Split) mappers.Splits {
	splits.mapper.update(splits.context, split)
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List[i] = split
			break
		}
	}
	return splits
}

func NewSplits(Mapper utilities.Mapper, context sdkTypes.Context) mappers.Splits {
	switch mapper := Mapper.(type) {
	case splitsMapper:
		return splits{
			ID:      readSplitID(""),
			List:    []entities.Split{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
