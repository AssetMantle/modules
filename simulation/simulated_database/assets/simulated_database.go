// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package assets

var (
	AccountClassificationIDAssetIDMap map[string]map[string]string = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap  map[string][]byte            = make(map[string][]byte)
)

func AddAssetData(address, classificationID, assetID string) {
	classificationAssetMap := make(map[string]string)
	classificationAssetMap[classificationID] = assetID
	AccountClassificationIDAssetIDMap[address] = classificationAssetMap
}
func AddMappableBytes(classificationID string, mappableBytes []byte) {
	ClassificationIDMappableBytesMap[classificationID] = mappableBytes
}
func GetMappableBytes(classificationID string) []byte {
	return ClassificationIDMappableBytesMap[classificationID]
}
func GetAssetData(address string) map[string]string {
	return AccountClassificationIDAssetIDMap[address]
}
func ClearAll() {
	AccountClassificationIDAssetIDMap = make(map[string]map[string]string)
	ClassificationIDMappableBytesMap = make(map[string][]byte)
}
