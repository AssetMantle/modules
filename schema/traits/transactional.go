package traits

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Transactional interface {
	Splittable
	Send(sdkTypes.Dec) Transactional
	Receive(sdkTypes.Dec) Transactional

	CanSend(sdkTypes.Dec) bool
}
