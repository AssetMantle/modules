package docs

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	propertiesUtilities "github.com/AssetMantle/modules/schema/properties/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, request{})
}

func handler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := Prototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			return
			//} else if reflect.TypeOf(Prototype()) != reflect.TypeOf(transactionRequest) {
			//	rest.CheckBadRequestError(responseWriter, errorConstants.InvalidRequest)
			//	return
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			return
		}

		req := transactionRequest.(request)

		immutableMetaProperties, _ := utilities.ReadMetaPropertyList(req.ImmutableMetaProperties)

		immutableProperties, _ := utilities.ReadMetaPropertyList(req.ImmutableProperties)

		immutableProperties = immutableProperties.ScrubData()

		mutableMetaProperties, _ := utilities.ReadMetaPropertyList(req.MutableMetaProperties)

		mutableProperties, _ := utilities.ReadMetaPropertyList(req.MutableProperties)
		mutableProperties = mutableProperties.ScrubData()

		immutables := base.NewImmutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(immutableMetaProperties.GetList(), immutableProperties.GetList()...)...)...))

		mutables := base.NewMutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(mutableMetaProperties.GetList(), mutableProperties.GetList()...)...)...))

		x := GetID(immutables, mutables)

		rest.PostProcessResponse(responseWriter, context, newResponse(x, nil))
	}
}

func GenerateHashID(toHashList ...[]byte) []byte {
	var nonEmptyByteList [][]byte

	for _, value := range toHashList {
		if len(value) != 0 {
			nonEmptyByteList = append(nonEmptyByteList, value)
		}
	}

	if len(nonEmptyByteList) == 0 {
		return nil
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return hash.Sum(nil)
}

func GetID(immutables qualified.Immutables, mutables qualified.Mutables) string {
	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))
	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().GetList()))
	for i, property := range mutables.GetMutablePropertyList().GetList() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}
	id := GenerateHashID(GenerateHashID(immutableIDByteList...), GenerateHashID(mutableIDByteList...), immutables.GenerateHashID().Bytes())
	return base64.URLEncoding.EncodeToString(GenerateHashID(id, immutables.GenerateHashID().Bytes()))
}

func RegisterRESTRoutes(context client.Context, router *mux.Router) {
	router.HandleFunc("/docs/get", handler(context)).Methods("POST")
}
