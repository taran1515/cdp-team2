// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protos/reward.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Points      int32  `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

func (x *Reward) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_protos_reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Reward) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Reward) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

type RewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *RewardRequest) Reset() {
	*x = RewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardRequest) ProtoMessage() {}

func (x *RewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardRequest.ProtoReflect.Descriptor instead.
func (*RewardRequest) Descriptor() ([]byte, []int) {
	return file_protos_reward_proto_rawDescGZIP(), []int{1}
}

func (x *RewardRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type RewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward *Reward `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *RewardResponse) Reset() {
	*x = RewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardResponse) ProtoMessage() {}

func (x *RewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardResponse.ProtoReflect.Descriptor instead.
func (*RewardResponse) Descriptor() ([]byte, []int) {
	return file_protos_reward_proto_rawDescGZIP(), []int{2}
}

func (x *RewardResponse) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// product id associated with a particular product
	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	// reward associated with a particular product
	Reward int32 `protobuf:"varint,2,opt,name=reward,proto3" json:"reward,omitempty"`
	// points associated with a particular product
	Points int32 `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
	// count is the total number of product item with same product id in the cart
	Count int32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[3]
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
	return file_protos_reward_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetReward() int32 {
	if x != nil {
		return x.Reward
	}
	return 0
}

func (x *Product) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *Product) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CalculateRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total price of the current order on which reward is to be calculated
	TotalPrice int32 `protobuf:"varint,1,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	// array of products with information related to Product
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CalculateRewardRequest) Reset() {
	*x = CalculateRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRewardRequest) ProtoMessage() {}

func (x *CalculateRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRewardRequest.ProtoReflect.Descriptor instead.
func (*CalculateRewardRequest) Descriptor() ([]byte, []int) {
	return file_protos_reward_proto_rawDescGZIP(), []int{4}
}

func (x *CalculateRewardRequest) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *CalculateRewardRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type CalculateRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// final price is the total price of the order after deducting the reward/points
	FinalPrice int32 `protobuf:"varint,1,opt,name=finalPrice,proto3" json:"finalPrice,omitempty"`
	// error represents the error occurred while creating the response
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CalculateRewardResponse) Reset() {
	*x = CalculateRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reward_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRewardResponse) ProtoMessage() {}

func (x *CalculateRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reward_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRewardResponse.ProtoReflect.Descriptor instead.
func (*CalculateRewardResponse) Descriptor() ([]byte, []int) {
	return file_protos_reward_proto_rawDescGZIP(), []int{5}
}

func (x *CalculateRewardResponse) GetFinalPrice() int32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

func (x *CalculateRewardResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protos_reward_proto protoreflect.FileDescriptor

var file_protos_reward_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x62, 0x0a,
	0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x22, 0x2f, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x6d, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x16, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0x9f, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_reward_proto_rawDescOnce sync.Once
	file_protos_reward_proto_rawDescData = file_protos_reward_proto_rawDesc
)

func file_protos_reward_proto_rawDescGZIP() []byte {
	file_protos_reward_proto_rawDescOnce.Do(func() {
		file_protos_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_reward_proto_rawDescData)
	})
	return file_protos_reward_proto_rawDescData
}

var file_protos_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_reward_proto_goTypes = []interface{}{
	(*Reward)(nil),                  // 0: reward.Reward
	(*RewardRequest)(nil),           // 1: reward.RewardRequest
	(*RewardResponse)(nil),          // 2: reward.RewardResponse
	(*Product)(nil),                 // 3: reward.Product
	(*CalculateRewardRequest)(nil),  // 4: reward.CalculateRewardRequest
	(*CalculateRewardResponse)(nil), // 5: reward.CalculateRewardResponse
}
var file_protos_reward_proto_depIdxs = []int32{
	0, // 0: reward.RewardResponse.reward:type_name -> reward.Reward
	3, // 1: reward.CalculateRewardRequest.products:type_name -> reward.Product
	1, // 2: reward.RewardService.GetReward:input_type -> reward.RewardRequest
	4, // 3: reward.RewardService.CalculateReward:input_type -> reward.CalculateRewardRequest
	2, // 4: reward.RewardService.GetReward:output_type -> reward.RewardResponse
	5, // 5: reward.RewardService.CalculateReward:output_type -> reward.CalculateRewardResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_reward_proto_init() }
func file_protos_reward_proto_init() {
	if File_protos_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
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
		file_protos_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardRequest); i {
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
		file_protos_reward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardResponse); i {
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
		file_protos_reward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_reward_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRewardRequest); i {
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
		file_protos_reward_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRewardResponse); i {
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
			RawDescriptor: file_protos_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_reward_proto_goTypes,
		DependencyIndexes: file_protos_reward_proto_depIdxs,
		MessageInfos:      file_protos_reward_proto_msgTypes,
	}.Build()
	File_protos_reward_proto = out.File
	file_protos_reward_proto_rawDesc = nil
	file_protos_reward_proto_goTypes = nil
	file_protos_reward_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RewardServiceClient is the client API for RewardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RewardServiceClient interface {
	// Get the total reward currently associated with a customer with given customerId
	GetReward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error)
	// Calculates the final pricing after deducting the reward/points from the total price for the user
	CalculateReward(ctx context.Context, in *CalculateRewardRequest, opts ...grpc.CallOption) (*CalculateRewardResponse, error)
}

type rewardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardServiceClient(cc grpc.ClientConnInterface) RewardServiceClient {
	return &rewardServiceClient{cc}
}

func (c *rewardServiceClient) GetReward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error) {
	out := new(RewardResponse)
	err := c.cc.Invoke(ctx, "/reward.RewardService/GetReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardServiceClient) CalculateReward(ctx context.Context, in *CalculateRewardRequest, opts ...grpc.CallOption) (*CalculateRewardResponse, error) {
	out := new(CalculateRewardResponse)
	err := c.cc.Invoke(ctx, "/reward.RewardService/CalculateReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardServiceServer is the server API for RewardService service.
type RewardServiceServer interface {
	// Get the total reward currently associated with a customer with given customerId
	GetReward(context.Context, *RewardRequest) (*RewardResponse, error)
	// Calculates the final pricing after deducting the reward/points from the total price for the user
	CalculateReward(context.Context, *CalculateRewardRequest) (*CalculateRewardResponse, error)
}

// UnimplementedRewardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRewardServiceServer struct {
}

func (*UnimplementedRewardServiceServer) GetReward(context.Context, *RewardRequest) (*RewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReward not implemented")
}
func (*UnimplementedRewardServiceServer) CalculateReward(context.Context, *CalculateRewardRequest) (*CalculateRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateReward not implemented")
}

func RegisterRewardServiceServer(s *grpc.Server, srv RewardServiceServer) {
	s.RegisterService(&_RewardService_serviceDesc, srv)
}

func _RewardService_GetReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServiceServer).GetReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reward.RewardService/GetReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServiceServer).GetReward(ctx, req.(*RewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RewardService_CalculateReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServiceServer).CalculateReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reward.RewardService/CalculateReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServiceServer).CalculateReward(ctx, req.(*CalculateRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RewardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reward.RewardService",
	HandlerType: (*RewardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReward",
			Handler:    _RewardService_GetReward_Handler,
		},
		{
			MethodName: "CalculateReward",
			Handler:    _RewardService_CalculateReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/reward.proto",
}
