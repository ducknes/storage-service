// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: database/pb/products.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string         `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	BrandName   string         `protobuf:"bytes,2,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	FactoryName string         `protobuf:"bytes,3,opt,name=FactoryName,proto3" json:"FactoryName,omitempty"`
	Name        string         `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string         `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32        `protobuf:"fixed32,6,opt,name=Price,proto3" json:"Price,omitempty"`
	Items       []*ProductItem `protobuf:"bytes,7,rep,name=Items,proto3" json:"Items,omitempty"`
	Materials   []string       `protobuf:"bytes,8,rep,name=Materials,proto3" json:"Materials,omitempty"`
	Images      []string       `protobuf:"bytes,9,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_pb_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_database_pb_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_database_pb_products_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *Product) GetFactoryName() string {
	if x != nil {
		return x.FactoryName
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetItems() []*ProductItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Product) GetMaterials() []string {
	if x != nil {
		return x.Materials
	}
	return nil
}

func (x *Product) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type ProductItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCount int32   `protobuf:"varint,1,opt,name=StockCount,proto3" json:"StockCount,omitempty"`
	Size       int32   `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Weight     float32 `protobuf:"fixed32,3,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Color      string  `protobuf:"bytes,4,opt,name=Color,proto3" json:"Color,omitempty"`
}

func (x *ProductItem) Reset() {
	*x = ProductItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_pb_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductItem) ProtoMessage() {}

func (x *ProductItem) ProtoReflect() protoreflect.Message {
	mi := &file_database_pb_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductItem.ProtoReflect.Descriptor instead.
func (*ProductItem) Descriptor() ([]byte, []int) {
	return file_database_pb_products_proto_rawDescGZIP(), []int{1}
}

func (x *ProductItem) GetStockCount() int32 {
	if x != nil {
		return x.StockCount
	}
	return 0
}

func (x *ProductItem) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProductItem) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ProductItem) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type MapProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*Product `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapProducts) Reset() {
	*x = MapProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_pb_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapProducts) ProtoMessage() {}

func (x *MapProducts) ProtoReflect() protoreflect.Message {
	mi := &file_database_pb_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapProducts.ProtoReflect.Descriptor instead.
func (*MapProducts) Descriptor() ([]byte, []int) {
	return file_database_pb_products_proto_rawDescGZIP(), []int{2}
}

func (x *MapProducts) GetItems() map[string]*Product {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_database_pb_products_proto protoreflect.FileDescriptor

var file_database_pb_products_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x6f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4b, 0x0a, 0x0a, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x62, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_pb_products_proto_rawDescOnce sync.Once
	file_database_pb_products_proto_rawDescData = file_database_pb_products_proto_rawDesc
)

func file_database_pb_products_proto_rawDescGZIP() []byte {
	file_database_pb_products_proto_rawDescOnce.Do(func() {
		file_database_pb_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_pb_products_proto_rawDescData)
	})
	return file_database_pb_products_proto_rawDescData
}

var file_database_pb_products_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_database_pb_products_proto_goTypes = []any{
	(*Product)(nil),     // 0: pbmodels.Product
	(*ProductItem)(nil), // 1: pbmodels.ProductItem
	(*MapProducts)(nil), // 2: pbmodels.MapProducts
	nil,                 // 3: pbmodels.MapProducts.ItemsEntry
}
var file_database_pb_products_proto_depIdxs = []int32{
	1, // 0: pbmodels.Product.Items:type_name -> pbmodels.ProductItem
	3, // 1: pbmodels.MapProducts.Items:type_name -> pbmodels.MapProducts.ItemsEntry
	0, // 2: pbmodels.MapProducts.ItemsEntry.value:type_name -> pbmodels.Product
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_database_pb_products_proto_init() }
func file_database_pb_products_proto_init() {
	if File_database_pb_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_pb_products_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_database_pb_products_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProductItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_database_pb_products_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MapProducts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_database_pb_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_database_pb_products_proto_goTypes,
		DependencyIndexes: file_database_pb_products_proto_depIdxs,
		MessageInfos:      file_database_pb_products_proto_msgTypes,
	}.Build()
	File_database_pb_products_proto = out.File
	file_database_pb_products_proto_rawDesc = nil
	file_database_pb_products_proto_goTypes = nil
	file_database_pb_products_proto_depIdxs = nil
}