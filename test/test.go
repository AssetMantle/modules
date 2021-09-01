package test

import "github.com/persistenceOne/persistenceSDK/schema/types/base"

//
//import (
//	"fmt"
//	sdkTypes "github.com/cosmos/cosmos-sdk/types"
//	"github.com/persistenceOne/persistenceSDK/modules/metas"
//	"reflect"
//)
//
//func CaseSwitch(msg sdkTypes.Msg)  {
//	fmt.Println(reflect.TypeOf(msg))
//	fmt.Println(reflect.TypeOf(metas.Message()))
//	//switch msg {
//	//case metas.Message() :
//	//
//	//}
//}

func NewRevealMessage(address base.AccAddress, fact base.MetaFact) *Message {
	return &Message{
		From:     address,
		MetaFact: fact,
	}
}
