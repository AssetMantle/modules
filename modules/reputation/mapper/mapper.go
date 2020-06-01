package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/reputation/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var storeKeyPrefix = []byte{0x11}

func storeKey(interNFTAddress types.InterNFTAddress) []byte {
	return append(storeKeyPrefix, interNFTAddress.Bytes()...)
}

type Mapper interface {
	Create(sdkTypes.Context, types.InterNFTAddress, sdkTypes.AccAddress, bool) error
	Read(sdkTypes.Context, types.InterNFTAddress) (types.InterNFT, error)
	Update(sdkTypes.Context, types.InterNFT) error
	Delete(sdkTypes.Context, types.InterNFTAddress)
}

type baseMapper struct {
	storeKey sdkTypes.StoreKey
	codec    *codec.Codec
}

func NewMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) Mapper {
	return baseMapper{
		storeKey: storeKey,
		codec:    codec,
	}
}

var _ Mapper = (*baseMapper)(nil)

func (baseMapper baseMapper) Create(context sdkTypes.Context, address types.InterNFTAddress, owner sdkTypes.AccAddress, lock bool) error {
	interNFT := newInterNFT(address, owner, lock)
	bytes, err := baseMapper.codec.MarshalBinaryBare(interNFT)
	if err != nil {
		panic(err)
	}
	interNFTAddress := interNFT.GetAddress()
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(interNFTAddress), bytes)
	return nil
}
func (baseMapper baseMapper) Read(context sdkTypes.Context, address types.InterNFTAddress) (interNFT types.InterNFT, error error) {
	kvStore := context.KVStore(baseMapper.storeKey)
	bytes := kvStore.Get(storeKey(address))
	if bytes == nil {
		return nil, errors.Wrap(constants.EntityNotFoundCode, address.String())
	}
	err := baseMapper.codec.UnmarshalBinaryBare(bytes, &interNFT)
	if err != nil {
		panic(err)
	}
	return interNFT, nil
}
func (baseMapper baseMapper) Update(context sdkTypes.Context, interNFT types.InterNFT) error {
	bytes, err := baseMapper.codec.MarshalBinaryBare(interNFT)
	if err != nil {
		panic(err)
	}
	interNFTAddress := interNFT.GetAddress()
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(interNFTAddress), bytes)
	return nil
}
func (baseMapper baseMapper) Delete(context sdkTypes.Context, address types.InterNFTAddress) {
	bytes, err := baseMapper.codec.MarshalBinaryBare(&baseInterNFT{})
	if err != nil {
		panic(err)
	}
	kvStore := context.KVStore(baseMapper.storeKey)
	kvStore.Set(storeKey(address), bytes)
}
