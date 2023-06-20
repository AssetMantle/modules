package classifications

var (
	AccountAssetClassificationIDMap          map[string]string = make(map[string]string)
	AccountIdentityClassificationIDMap       map[string]string = make(map[string]string)
	AssetClassificationIDMappableBytesMap    map[string][]byte = make(map[string][]byte)
	IdentityClassificationIDMappableBytesMap map[string][]byte = make(map[string][]byte)
)

func AddIdentityClassificationID(address string, classificationID string) {
	AccountIdentityClassificationIDMap[address] = classificationID
}
func AddAssetClassificationID(address string, classificationID string) {
	AccountAssetClassificationIDMap[address] = classificationID
}
func AddAssetMappableBytes(classificationID string, mappableBytes []byte) {
	AssetClassificationIDMappableBytesMap[classificationID] = mappableBytes
}
func AddIdentityMappableBytes(classificationID string, mappableBytes []byte) {
	IdentityClassificationIDMappableBytesMap[classificationID] = mappableBytes
}
func ClearAll() {
	AccountIdentityClassificationIDMap = make(map[string]string)
	IdentityClassificationIDMappableBytesMap = make(map[string][]byte)
	AccountAssetClassificationIDMap = make(map[string]string)
	AssetClassificationIDMappableBytesMap = make(map[string][]byte)
}
