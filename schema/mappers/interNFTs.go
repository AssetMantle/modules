package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFTs interface {
	GetID() types.ID

	Get(types.ID) mappables.InterNFT
	GetList() []mappables.InterNFT

	Fetch(types.ID) InterNFTs
	Add(mappables.InterNFT) InterNFTs
	Remove(mappables.InterNFT) InterNFTs
	Mutate(mappables.InterNFT) InterNFTs
}
