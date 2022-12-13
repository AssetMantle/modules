package base

//
//import (
//	"github.com/AssetMantle/modules/schema/documents"
//	"github.com/AssetMantle/modules/schema/ids"
//	"github.com/AssetMantle/modules/schema/properties"
//	"github.com/AssetMantle/modules/schema/qualified"
//)
//
//var _ documents.Document = (*Documents)(nil)
//
//func (x *Documents) GenerateHashID() ids.HashID {
//	return x.Impl.(documents.Document).GenerateHashID()
//}
//
//func (x *Documents) GetClassificationID() ids.ClassificationID {
//	return x.Impl.(documents.Document).GetClassificationID()
//}
//
//func (x *Documents) GetProperty(id ids.PropertyID) properties.Property {
//	return x.Impl.(documents.Document).GetProperty(id)
//}
//
//func (x *Documents) GetImmutables() qualified.Immutables {
//	return x.Impl.(documents.Document).GetImmutables()
//}
//
//func (x *Documents) GetMutables() qualified.Mutables {
//	return x.Impl.(documents.Document).GetMutables()
//}
//
//func (x *Documents) Mutate(property ...properties.Property) documents.Document {
//	return x.Impl.(documents.Document).Mutate(property...)
//}
