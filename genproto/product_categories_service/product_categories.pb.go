// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: product_categories.proto

package product_categories_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ProductCategoriesPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductCategoriesPrimaryKey) Reset() {
	*x = ProductCategoriesPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_categories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategoriesPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategoriesPrimaryKey) ProtoMessage() {}

func (x *ProductCategoriesPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_product_categories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategoriesPrimaryKey.ProtoReflect.Descriptor instead.
func (*ProductCategoriesPrimaryKey) Descriptor() ([]byte, []int) {
	return file_product_categories_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCategoriesPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateProductCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateProductCategories) Reset() {
	*x = CreateProductCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_categories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductCategories) ProtoMessage() {}

func (x *CreateProductCategories) ProtoReflect() protoreflect.Message {
	mi := &file_product_categories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductCategories.ProtoReflect.Descriptor instead.
func (*CreateProductCategories) Descriptor() ([]byte, []int) {
	return file_product_categories_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductCategories) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateProductCategories) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type GetProductCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId  string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *GetProductCategories) Reset() {
	*x = GetProductCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_categories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductCategories) ProtoMessage() {}

func (x *GetProductCategories) ProtoReflect() protoreflect.Message {
	mi := &file_product_categories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductCategories.ProtoReflect.Descriptor instead.
func (*GetProductCategories) Descriptor() ([]byte, []int) {
	return file_product_categories_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductCategories) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProductCategories) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *GetProductCategories) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type UpdateProductCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId  string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *UpdateProductCategories) Reset() {
	*x = UpdateProductCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_categories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductCategories) ProtoMessage() {}

func (x *UpdateProductCategories) ProtoReflect() protoreflect.Message {
	mi := &file_product_categories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductCategories.ProtoReflect.Descriptor instead.
func (*UpdateProductCategories) Descriptor() ([]byte, []int) {
	return file_product_categories_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductCategories) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductCategories) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateProductCategories) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

var File_product_categories_proto protoreflect.FileDescriptor

var file_product_categories_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x32, 0xea, 0x03, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x77, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x1a, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74,
	0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63,
	0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_categories_proto_rawDescOnce sync.Once
	file_product_categories_proto_rawDescData = file_product_categories_proto_rawDesc
)

func file_product_categories_proto_rawDescGZIP() []byte {
	file_product_categories_proto_rawDescOnce.Do(func() {
		file_product_categories_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_categories_proto_rawDescData)
	})
	return file_product_categories_proto_rawDescData
}

var file_product_categories_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_product_categories_proto_goTypes = []interface{}{
	(*ProductCategoriesPrimaryKey)(nil), // 0: product_catrgories_service_go.ProductCategoriesPrimaryKey
	(*CreateProductCategories)(nil),     // 1: product_catrgories_service_go.CreateProductCategories
	(*GetProductCategories)(nil),        // 2: product_catrgories_service_go.GetProductCategories
	(*UpdateProductCategories)(nil),     // 3: product_catrgories_service_go.UpdateProductCategories
	(*empty.Empty)(nil),                 // 4: google.protobuf.Empty
}
var file_product_categories_proto_depIdxs = []int32{
	1, // 0: product_catrgories_service_go.ProductCategoriesService.Create:input_type -> product_catrgories_service_go.CreateProductCategories
	0, // 1: product_catrgories_service_go.ProductCategoriesService.GetByID:input_type -> product_catrgories_service_go.ProductCategoriesPrimaryKey
	3, // 2: product_catrgories_service_go.ProductCategoriesService.Update:input_type -> product_catrgories_service_go.UpdateProductCategories
	0, // 3: product_catrgories_service_go.ProductCategoriesService.Delete:input_type -> product_catrgories_service_go.ProductCategoriesPrimaryKey
	2, // 4: product_catrgories_service_go.ProductCategoriesService.Create:output_type -> product_catrgories_service_go.GetProductCategories
	2, // 5: product_catrgories_service_go.ProductCategoriesService.GetByID:output_type -> product_catrgories_service_go.GetProductCategories
	2, // 6: product_catrgories_service_go.ProductCategoriesService.Update:output_type -> product_catrgories_service_go.GetProductCategories
	4, // 7: product_catrgories_service_go.ProductCategoriesService.Delete:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_categories_proto_init() }
func file_product_categories_proto_init() {
	if File_product_categories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_categories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategoriesPrimaryKey); i {
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
		file_product_categories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductCategories); i {
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
		file_product_categories_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductCategories); i {
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
		file_product_categories_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductCategories); i {
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
			RawDescriptor: file_product_categories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_categories_proto_goTypes,
		DependencyIndexes: file_product_categories_proto_depIdxs,
		MessageInfos:      file_product_categories_proto_msgTypes,
	}.Build()
	File_product_categories_proto = out.File
	file_product_categories_proto_rawDesc = nil
	file_product_categories_proto_goTypes = nil
	file_product_categories_proto_depIdxs = nil
}
