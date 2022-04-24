// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/order.proto

package order

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

type ProductDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32   `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Points    int32   `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
	Reward    int32   `protobuf:"varint,3,opt,name=reward,proto3" json:"reward,omitempty"`
	Quantity  int32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductDesc) Reset() {
	*x = ProductDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDesc) ProtoMessage() {}

func (x *ProductDesc) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDesc.ProtoReflect.Descriptor instead.
func (*ProductDesc) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{0}
}

func (x *ProductDesc) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductDesc) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *ProductDesc) GetReward() int32 {
	if x != nil {
		return x.Reward
	}
	return 0
}

func (x *ProductDesc) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductDesc) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  string         `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Status      int32          `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	AddressId   int32          `protobuf:"varint,3,opt,name=addressId,proto3" json:"addressId,omitempty"`
	TotalPrice  float32        `protobuf:"fixed32,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	PayedPrice  float32        `protobuf:"fixed32,5,opt,name=payedPrice,proto3" json:"payedPrice,omitempty"`
	ProductDesc []*ProductDesc `protobuf:"bytes,6,rep,name=productDesc,proto3" json:"productDesc,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Order) GetAddressId() int32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Order) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetPayedPrice() float32 {
	if x != nil {
		return x.PayedPrice
	}
	return 0
}

func (x *Order) GetProductDesc() []*ProductDesc {
	if x != nil {
		return x.ProductDesc
	}
	return nil
}

type OrderStatusUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderStatusUpdateRequest) Reset() {
	*x = OrderStatusUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusUpdateRequest) ProtoMessage() {}

func (x *OrderStatusUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrderStatusUpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderStatusUpdateRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatusUpdateRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type OrderStatusUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderStatusUpdateResponse) Reset() {
	*x = OrderStatusUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusUpdateResponse) ProtoMessage() {}

func (x *OrderStatusUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderStatusUpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderStatusUpdateResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type OrderPayedPriceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string  `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	PayedPrice float32 `protobuf:"fixed32,2,opt,name=payedPrice,proto3" json:"payedPrice,omitempty"`
}

func (x *OrderPayedPriceUpdateRequest) Reset() {
	*x = OrderPayedPriceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedPriceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedPriceUpdateRequest) ProtoMessage() {}

func (x *OrderPayedPriceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedPriceUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrderPayedPriceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderPayedPriceUpdateRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderPayedPriceUpdateRequest) GetPayedPrice() float32 {
	if x != nil {
		return x.PayedPrice
	}
	return 0
}

type OrderPayedPriceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderPayedPriceUpdateResponse) Reset() {
	*x = OrderPayedPriceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedPriceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedPriceUpdateResponse) ProtoMessage() {}

func (x *OrderPayedPriceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedPriceUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderPayedPriceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderPayedPriceUpdateResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type OrderAddressUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	AddressId int32  `protobuf:"varint,2,opt,name=addressId,proto3" json:"addressId,omitempty"`
}

func (x *OrderAddressUpdateRequest) Reset() {
	*x = OrderAddressUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddressUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddressUpdateRequest) ProtoMessage() {}

func (x *OrderAddressUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddressUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrderAddressUpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderAddressUpdateRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderAddressUpdateRequest) GetAddressId() int32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

type OrderAddressUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderAddressUpdateResponse) Reset() {
	*x = OrderAddressUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddressUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddressUpdateResponse) ProtoMessage() {}

func (x *OrderAddressUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddressUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderAddressUpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderAddressUpdateResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  string         `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Status      int32          `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	AddressId   int32          `protobuf:"varint,3,opt,name=addressId,proto3" json:"addressId,omitempty"`
	TotalPrice  float32        `protobuf:"fixed32,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	PayedPrice  float32        `protobuf:"fixed32,5,opt,name=payedPrice,proto3" json:"payedPrice,omitempty"`
	ProductDesc []*ProductDesc `protobuf:"bytes,6,rep,name=productDesc,proto3" json:"productDesc,omitempty"`
	OrderId     string         `protobuf:"bytes,7,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *OrderResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderResponse) GetAddressId() int32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *OrderResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderResponse) GetPayedPrice() float32 {
	if x != nil {
		return x.PayedPrice
	}
	return 0
}

func (x *OrderResponse) GetProductDesc() []*ProductDesc {
	if x != nil {
		return x.ProductDesc
	}
	return nil
}

func (x *OrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{9}
}

func (x *OrderCreateRequest) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *OrderResponse `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{10}
}

func (x *OrderCreateResponse) GetOrder() *OrderResponse {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{11}
}

func (x *GetOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *OrderResponse `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{12}
}

func (x *GetOrderResponse) GetOrder() *OrderResponse {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_protos_order_proto protoreflect.FileDescriptor

var file_protos_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x70, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x22, 0x4c, 0x0a, 0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x37, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x1c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x3b, 0x0a, 0x1d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x53, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf5,
	0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x79, 0x65,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x41, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x32, 0xd0, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x12,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x18, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x1c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_proto_rawDescOnce sync.Once
	file_protos_order_proto_rawDescData = file_protos_order_proto_rawDesc
)

func file_protos_order_proto_rawDescGZIP() []byte {
	file_protos_order_proto_rawDescOnce.Do(func() {
		file_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_proto_rawDescData)
	})
	return file_protos_order_proto_rawDescData
}

var file_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protos_order_proto_goTypes = []interface{}{
	(*ProductDesc)(nil),                   // 0: proto.ProductDesc
	(*Order)(nil),                         // 1: proto.Order
	(*OrderStatusUpdateRequest)(nil),      // 2: proto.OrderStatusUpdateRequest
	(*OrderStatusUpdateResponse)(nil),     // 3: proto.OrderStatusUpdateResponse
	(*OrderPayedPriceUpdateRequest)(nil),  // 4: proto.OrderPayedPriceUpdateRequest
	(*OrderPayedPriceUpdateResponse)(nil), // 5: proto.OrderPayedPriceUpdateResponse
	(*OrderAddressUpdateRequest)(nil),     // 6: proto.OrderAddressUpdateRequest
	(*OrderAddressUpdateResponse)(nil),    // 7: proto.OrderAddressUpdateResponse
	(*OrderResponse)(nil),                 // 8: proto.OrderResponse
	(*OrderCreateRequest)(nil),            // 9: proto.OrderCreateRequest
	(*OrderCreateResponse)(nil),           // 10: proto.OrderCreateResponse
	(*GetOrderRequest)(nil),               // 11: proto.GetOrderRequest
	(*GetOrderResponse)(nil),              // 12: proto.GetOrderResponse
}
var file_protos_order_proto_depIdxs = []int32{
	0,  // 0: proto.Order.productDesc:type_name -> proto.ProductDesc
	0,  // 1: proto.OrderResponse.productDesc:type_name -> proto.ProductDesc
	1,  // 2: proto.OrderCreateRequest.order:type_name -> proto.Order
	8,  // 3: proto.OrderCreateResponse.order:type_name -> proto.OrderResponse
	8,  // 4: proto.GetOrderResponse.order:type_name -> proto.OrderResponse
	9,  // 5: proto.Service.OrderCreateService:input_type -> proto.OrderCreateRequest
	2,  // 6: proto.Service.OrderStatusUpdateService:input_type -> proto.OrderStatusUpdateRequest
	11, // 7: proto.Service.GetOrderService:input_type -> proto.GetOrderRequest
	6,  // 8: proto.Service.OrderAddressUpdateService:input_type -> proto.OrderAddressUpdateRequest
	4,  // 9: proto.Service.OrderPayedPriceUpdateService:input_type -> proto.OrderPayedPriceUpdateRequest
	10, // 10: proto.Service.OrderCreateService:output_type -> proto.OrderCreateResponse
	3,  // 11: proto.Service.OrderStatusUpdateService:output_type -> proto.OrderStatusUpdateResponse
	12, // 12: proto.Service.GetOrderService:output_type -> proto.GetOrderResponse
	7,  // 13: proto.Service.OrderAddressUpdateService:output_type -> proto.OrderAddressUpdateResponse
	5,  // 14: proto.Service.OrderPayedPriceUpdateService:output_type -> proto.OrderPayedPriceUpdateResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_order_proto_init() }
func file_protos_order_proto_init() {
	if File_protos_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDesc); i {
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
		file_protos_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_protos_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusUpdateRequest); i {
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
		file_protos_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusUpdateResponse); i {
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
		file_protos_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedPriceUpdateRequest); i {
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
		file_protos_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedPriceUpdateResponse); i {
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
		file_protos_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddressUpdateRequest); i {
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
		file_protos_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddressUpdateResponse); i {
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
		file_protos_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_protos_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_protos_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
		file_protos_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_protos_order_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResponse); i {
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
			RawDescriptor: file_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_order_proto_goTypes,
		DependencyIndexes: file_protos_order_proto_depIdxs,
		MessageInfos:      file_protos_order_proto_msgTypes,
	}.Build()
	File_protos_order_proto = out.File
	file_protos_order_proto_rawDesc = nil
	file_protos_order_proto_goTypes = nil
	file_protos_order_proto_depIdxs = nil
}
