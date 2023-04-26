package identities

import (
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	AccountNubIDPairs            map[string]string            = make(map[string]string)
	AccountClassificationIDPairs map[string]string            = make(map[string]string)
	AccountIssuedIdentities      map[string]map[string]string = make(map[string]map[string]string)
)

func AddAccountNubIDPair(address sdkTypes.AccAddress, nubID string) {
	AccountNubIDPairs[address.String()] = nubID
}
func AddAccountClassificationIDPair(address sdkTypes.AccAddress, classificationID string) {
	AccountClassificationIDPairs[address.String()] = classificationID
}
func AddIssuedIdentity(from sdkTypes.AccAddress, identity string, to sdkTypes.AccAddress) {
	idToPair := make(map[string]string)
	idToPair[identity] = to.String()
	AccountIssuedIdentities[from.String()] = idToPair
}
func GetIssuedIdentityInfo(from sdkTypes.Address) map[string]string {
	return AccountIssuedIdentities[from.String()]
}
func GetClassificationID(address sdkTypes.Address) ids.ClassificationID {
	classificationID, _ := base.ReadClassificationID(AccountClassificationIDPairs[address.String()])
	return classificationID
}
func GetRandomAccNubIDPair() (sdkTypes.AccAddress, ids.IdentityID) {
	for x, i := range AccountNubIDPairs {
		id, _ := base.ReadIdentityID(i)
		return sdkTypes.AccAddress(x), id
	}
	return nil, nil
}
