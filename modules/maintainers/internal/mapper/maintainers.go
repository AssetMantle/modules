/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type maintainers struct {
	ID   types.ID               `json:"id" valid:"required~required field id missing"`
	List []mappables.Maintainer `json:"list" valid:"required~required field list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.Maintainers = (*maintainers)(nil)

func (maintainers maintainers) GetID() types.ID { return maintainers.ID }
func (maintainers maintainers) Get(id types.ID) mappables.Maintainer {
	maintainerID := maintainerIDFromInterface(id)
	for _, oldMaintainer := range maintainers.List {
		if oldMaintainer.GetID().Compare(maintainerID) == 0 {
			return oldMaintainer
		}
	}
	return nil
}
func (maintainers maintainers) GetList() []mappables.Maintainer {
	return maintainers.List
}

func (maintainers maintainers) Fetch(id types.ID) mappers.Maintainers {
	var maintainerList []mappables.Maintainer
	maintainersID := maintainerIDFromInterface(id)
	if len(maintainersID.IdentityID.Bytes()) > 0 {
		mappable := maintainers.mapper.Read(maintainers.context, maintainersID)
		if mappable != nil {
			maintainerList = append(maintainerList, mappable.(maintainer))
		}
	} else {
		appendMaintainerList := func(mappable traits.Mappable) bool {
			maintainerList = append(maintainerList, mappable.(maintainer))
			return false
		}
		maintainers.mapper.Iterate(maintainers.context, maintainersID, appendMaintainerList)
	}
	maintainers.ID, maintainers.List = id, maintainerList
	return maintainers
}
func (maintainers maintainers) Add(maintainer mappables.Maintainer) mappers.Maintainers {
	maintainers.ID = readMaintainerID("")
	maintainers.mapper.Create(maintainers.context, maintainer)
	maintainers.List = append(maintainers.List, maintainer)
	return maintainers
}
func (maintainers maintainers) Remove(maintainer mappables.Maintainer) mappers.Maintainers {
	maintainers.mapper.Delete(maintainers.context, maintainer.GetID())
	for i, oldMaintainer := range maintainers.List {
		if oldMaintainer.GetID().Compare(maintainer.GetID()) == 0 {
			maintainers.List = append(maintainers.List[:i], maintainers.List[i+1:]...)
			break
		}
	}
	return maintainers
}
func (maintainers maintainers) Mutate(maintainer mappables.Maintainer) mappers.Maintainers {
	maintainers.mapper.Update(maintainers.context, maintainer)
	for i, oldMaintainer := range maintainers.List {
		if oldMaintainer.GetID().Compare(maintainer.GetID()) == 0 {
			maintainers.List[i] = maintainer
			break
		}
	}
	return maintainers
}

func NewMaintainers(mapper helpers.Mapper, context sdkTypes.Context) mappers.Maintainers {
	return maintainers{
		ID:      readMaintainerID(""),
		List:    []mappables.Maintainer{},
		mapper:  mapper,
		context: context,
	}
}
