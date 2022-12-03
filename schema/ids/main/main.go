package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	base2 "github.com/AssetMantle/modules/schema/qualified/base"
)

func main() {
	fmt.Println("start")
	// fmt.Println("hash of deepanshutr:", base.GenerateHashID(base.NewStringID("deepanshutr1").Bytes(), base.NewStringID("deepanshutr2").Bytes()))
	accAddr, _ := types.AccAddressFromBech32("mantle1232v8efpehm3pr5eceueenfm3rnlc7zhlr9x8k")

	immutables := base2.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.NubIDProperty.GetKey(), baseData.NewIDData(base.NewStringID("")))))
	mutables := base2.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseLists.NewDataList(baseData.NewAccAddressData(accAddr))))))

	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))
	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}
	println(base.GenerateHashID(immutableIDByteList...).String())

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().GetList()))
	for i, property := range mutables.GetMutablePropertyList().GetList() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}
	println(base.GenerateHashID(mutableIDByteList...).String())

	println(immutables.GenerateHashID().String())

	println(base.GenerateHashID(base.GenerateHashID(immutableIDByteList...).Bytes(), base.GenerateHashID(mutableIDByteList...).Bytes(), immutables.GenerateHashID().Bytes()).String())
	fmt.Println(base.NewClassificationID(immutables, mutables).String())
	idImmutables := base2.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.NubIDProperty.GetKey(), baseData.NewIDData(base.NewStringID("deepanshutr")))))
	println(idImmutables.GenerateHashID().String())
	println(base.GenerateHashID(base.GenerateHashID(base.NewStringID("deepanshutr").Bytes()).Bytes()).String())
	println("end")
	for _, immutableProperty := range idImmutables.GetImmutablePropertyList().GetList() {
		println(immutableProperty.GetDataID().GetHashID().String())
	}
	// println(immutables.GenerateHashID().String())
	//
	// defaultImmutableByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))
	// for i, property := range immutables.GetImmutablePropertyList().GetList() {
	// 	// TODO test
	// 	if hashID := property.GetDataID().GetHashID(); !(hashID.Compare(base.GenerateHashID()) == 0) {
	// 		defaultImmutableByteList[i] = hashID.Bytes()
	// 	}
	// }

	// println(base.GenerateHashID(defaultImmutableByteList...).String())
}
