package identities

var (
	AccountClassificationIDIdentityIDMap map[string]map[string]string = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap     map[string][]byte            = make(map[string][]byte)
)

func AddIDData(address, classificationID, identityID string) {
	classificationIdentityMap := make(map[string]string)
	classificationIdentityMap[classificationID] = identityID
	AccountClassificationIDIdentityIDMap[address] = classificationIdentityMap
}
func AddMappableBytes(classificationID string, mappableBytes []byte) {
	ClassificationIDMappableBytesMap[classificationID] = mappableBytes
}
func GetMappableBytes(classificationID string) []byte {
	return ClassificationIDMappableBytesMap[classificationID]
}
func GetIDData(address string) map[string]string {
	return AccountClassificationIDIdentityIDMap[address]
}
func ClearAll() {
	AccountClassificationIDIdentityIDMap = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap = make(map[string][]byte)
}
