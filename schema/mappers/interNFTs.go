package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFTs interface {
	GetID() types.ID

	Get(types.ID) entities.InterNFT
	GetList() []entities.InterNFT

	Fetch(types.ID) InterNFTs
	Add(entities.InterNFT) InterNFTs
	Remove(entities.InterNFT) InterNFTs
	Mutate(entities.InterNFT) InterNFTs
}
