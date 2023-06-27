package orders

var (
	AccountClassificationIDOrderIDMap map[string]map[string]string = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap  map[string][]byte            = make(map[string][]byte)
)

func AddOrderData(address, classificationID, orderID string) {
	classificationIdentityMap := make(map[string]string)
	classificationIdentityMap[classificationID] = orderID
	AccountClassificationIDOrderIDMap[address] = classificationIdentityMap
}
func AddMappableBytes(classificationID string, mappableBytes []byte) {
	ClassificationIDMappableBytesMap[classificationID] = mappableBytes
}
func GetMappableBytes(classificationID string) []byte {
	return ClassificationIDMappableBytesMap[classificationID]
}
func GetOrderData(address string) map[string]string {
	return AccountClassificationIDOrderIDMap[address]
}
func ClearAll() {
	AccountClassificationIDOrderIDMap = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap = make(map[string][]byte)
}
