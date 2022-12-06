package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type orderIDI ids.OrderID

var _ ids2.OrderID = (*orderIDI)(nil)

func (orderIDI *orderIDI) Compare(listable traits.Listable) int {
	return orderIDI.Impl.(ids2.OrderID).Compare(listable)
}

func (orderIDI *orderIDI) String() string {
	return orderIDI.Impl.(ids2.OrderID).String()
}

func (orderIDI *orderIDI) Bytes() []byte {
	return orderIDI.Impl.(ids2.OrderID).Bytes()
}

func (orderIDI *orderIDI) IsOrderID() {

}
