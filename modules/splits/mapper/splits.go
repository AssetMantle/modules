package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type splits struct {
	ID   schema.ID
	List []schema.Split

	mapper  splitsMapper
	context sdkTypes.Context
}

var _ schema.Splits = (*splits)(nil)

func (splits splits) GetID() schema.ID { return splits.ID }
func (splits splits) Get(id schema.ID) schema.Split {
	splitID := splitIDFromInterface(id)
	for _, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(splitID) == 0 {
			return oldSplit
		}
	}
	return nil
}
func (splits splits) GetList() []schema.Split {
	return splits.List
}

func (splits splits) Fetch(id schema.ID) schema.Splits {
	var splitList []schema.Split
	splitsID := splitIDFromInterface(id)
	if len(splitsID.OwnableID.Bytes()) > 0 {
		split := splits.mapper.read(splits.context, splitsID)
		if split != nil {
			splitList = append(splitList, split)
		}
	} else {
		appendSplitList := func(split schema.Split) bool {
			splitList = append(splitList, split)
			return false
		}
		splits.mapper.iterate(splits.context, splitsID, appendSplitList)
	}
	splits.ID, splits.List = id, splitList
	return splits
}
func (splits splits) Add(split schema.Split) schema.Splits {
	splits.ID = readSplitID("")
	splits.mapper.create(splits.context, split)
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) < 0 {
			splits.List = append(append(splits.List[:i], split), splits.List[i+1:]...)
			break
		}
	}
	return splits
}
func (splits splits) Remove(split schema.Split) schema.Splits {
	splits.mapper.delete(splits.context, split.GetID())
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List = append(splits.List[:i], splits.List[i+1:]...)
			break
		}
	}
	return splits
}
func (splits splits) Mutate(split schema.Split) schema.Splits {
	splits.mapper.update(splits.context, split)
	for i, oldSplit := range splits.List {
		if oldSplit.GetID().Compare(split.GetID()) == 0 {
			splits.List[i] = split
			break
		}
	}
	return splits
}

func NewSplits(Mapper utility.Mapper, context sdkTypes.Context) schema.Splits {
	switch mapper := Mapper.(type) {
	case splitsMapper:
		return splits{
			ID:      readSplitID(""),
			List:    []schema.Split{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
