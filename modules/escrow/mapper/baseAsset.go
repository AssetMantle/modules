package mapper

import (
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type baseInterNFT struct {
	Address types.InterNFTAddress `json:"address" yaml:"address" valid:"required~address"`
	Owner   sdkTypes.AccAddress   `json:"owner" yaml:"owner" valid:"required~owner"`
	Lock    bool                  `json:"lock" yaml:"lock"`
}

func newInterNFT(address types.InterNFTAddress, owner sdkTypes.AccAddress, lock bool) types.InterNFT {
	return &baseInterNFT{
		Address: address,
		Owner:   owner,
		Lock:    lock,
	}
}

var _ types.InterNFT = (*baseInterNFT)(nil)

func (baseInterNFT baseInterNFT) GetAddress() types.InterNFTAddress   { return baseInterNFT.Address }
func (baseInterNFT baseInterNFT) GetOwner() sdkTypes.AccAddress       { return baseInterNFT.Owner }
func (baseInterNFT *baseInterNFT) SetOwner(Owner sdkTypes.AccAddress) { baseInterNFT.Owner = Owner }
func (baseInterNFT baseInterNFT) GetLock() bool                       { return baseInterNFT.Lock }
func (baseInterNFT *baseInterNFT) SetLock(Lock bool)                  { baseInterNFT.Lock = Lock }
func (baseInterNFT baseInterNFT) String() string {
	bytes, Error := json.Marshal(baseInterNFT)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}
