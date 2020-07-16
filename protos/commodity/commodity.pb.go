// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: commodity.proto

package commodity

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Price request defines the request pattern.
// Only one commodity can be queried per request for a price.
type CommodityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is a name of the searching commodity.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *CommodityRequest) Reset() {
	*x = CommodityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommodityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommodityRequest) ProtoMessage() {}

func (x *CommodityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommodityRequest.ProtoReflect.Descriptor instead.
func (*CommodityRequest) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

func (x *CommodityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CommodityResponse defines the response pattern.
type CommodityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the commodity.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Current price of the commodity.
	Price float32 `protobuf:"fixed32,2,opt,name=Price,proto3" json:"Price,omitempty"`
	// The currency of the price.
	Currency string `protobuf:"bytes,3,opt,name=Currency,proto3" json:"Currency,omitempty"`
	// The weight for which the price of the commodity is determined.
	WeightUnit string `protobuf:"bytes,4,opt,name=WeightUnit,proto3" json:"WeightUnit,omitempty"`
	// Last change in percentages.
	ChangeP float32 `protobuf:"fixed32,5,opt,name=ChangeP,proto3" json:"ChangeP,omitempty"`
	// Last Change in a number.
	ChangeN float32 `protobuf:"fixed32,6,opt,name=ChangeN,proto3" json:"ChangeN,omitempty"`
	// Last time updated.
	// The Unix time.
	LastUpdate int64 `protobuf:"varint,7,opt,name=LastUpdate,proto3" json:"LastUpdate,omitempty"`
}

func (x *CommodityResponse) Reset() {
	*x = CommodityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommodityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommodityResponse) ProtoMessage() {}

func (x *CommodityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommodityResponse.ProtoReflect.Descriptor instead.
func (*CommodityResponse) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

func (x *CommodityResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommodityResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CommodityResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CommodityResponse) GetWeightUnit() string {
	if x != nil {
		return x.WeightUnit
	}
	return ""
}

func (x *CommodityResponse) GetChangeP() float32 {
	if x != nil {
		return x.ChangeP
	}
	return 0
}

func (x *CommodityResponse) GetChangeN() float32 {
	if x != nil {
		return x.ChangeN
	}
	return 0
}

func (x *CommodityResponse) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

type StreamingCommodityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*StreamingCommodityResponse_CommodityResponse
	//	*StreamingCommodityResponse_Error
	Message isStreamingCommodityResponse_Message `protobuf_oneof:"message"`
}

func (x *StreamingCommodityResponse) Reset() {
	*x = StreamingCommodityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingCommodityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingCommodityResponse) ProtoMessage() {}

func (x *StreamingCommodityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingCommodityResponse.ProtoReflect.Descriptor instead.
func (*StreamingCommodityResponse) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{2}
}

func (m *StreamingCommodityResponse) GetMessage() isStreamingCommodityResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *StreamingCommodityResponse) GetCommodityResponse() *CommodityResponse {
	if x, ok := x.GetMessage().(*StreamingCommodityResponse_CommodityResponse); ok {
		return x.CommodityResponse
	}
	return nil
}

func (x *StreamingCommodityResponse) GetError() *status.Status {
	if x, ok := x.GetMessage().(*StreamingCommodityResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isStreamingCommodityResponse_Message interface {
	isStreamingCommodityResponse_Message()
}

type StreamingCommodityResponse_CommodityResponse struct {
	CommodityResponse *CommodityResponse `protobuf:"bytes,1,opt,name=commodity_response,json=commodityResponse,proto3,oneof"`
}

type StreamingCommodityResponse_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*StreamingCommodityResponse_CommodityResponse) isStreamingCommodityResponse_Message() {}

func (*StreamingCommodityResponse_Error) isStreamingCommodityResponse_Message() {}

var File_commodity_proto protoreflect.FileDescriptor

var file_commodity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8c, 0x01,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x11, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x11, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_commodity_proto_rawDescOnce sync.Once
	file_commodity_proto_rawDescData = file_commodity_proto_rawDesc
)

func file_commodity_proto_rawDescGZIP() []byte {
	file_commodity_proto_rawDescOnce.Do(func() {
		file_commodity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commodity_proto_rawDescData)
	})
	return file_commodity_proto_rawDescData
}

var file_commodity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commodity_proto_goTypes = []interface{}{
	(*CommodityRequest)(nil),           // 0: CommodityRequest
	(*CommodityResponse)(nil),          // 1: CommodityResponse
	(*StreamingCommodityResponse)(nil), // 2: StreamingCommodityResponse
	(*status.Status)(nil),              // 3: google.rpc.Status
}
var file_commodity_proto_depIdxs = []int32{
	1, // 0: StreamingCommodityResponse.commodity_response:type_name -> CommodityResponse
	3, // 1: StreamingCommodityResponse.error:type_name -> google.rpc.Status
	0, // 2: Commodity.GetCommodity:input_type -> CommodityRequest
	0, // 3: Commodity.SubscribeCommodity:input_type -> CommodityRequest
	1, // 4: Commodity.GetCommodity:output_type -> CommodityResponse
	2, // 5: Commodity.SubscribeCommodity:output_type -> StreamingCommodityResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commodity_proto_init() }
func file_commodity_proto_init() {
	if File_commodity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commodity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommodityRequest); i {
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
		file_commodity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommodityResponse); i {
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
		file_commodity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingCommodityResponse); i {
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
	file_commodity_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*StreamingCommodityResponse_CommodityResponse)(nil),
		(*StreamingCommodityResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commodity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commodity_proto_goTypes,
		DependencyIndexes: file_commodity_proto_depIdxs,
		MessageInfos:      file_commodity_proto_msgTypes,
	}.Build()
	File_commodity_proto = out.File
	file_commodity_proto_rawDesc = nil
	file_commodity_proto_goTypes = nil
	file_commodity_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommodityClient is the client API for Commodity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommodityClient interface {
	GetCommodity(ctx context.Context, in *CommodityRequest, opts ...grpc.CallOption) (*CommodityResponse, error)
	SubscribeCommodity(ctx context.Context, opts ...grpc.CallOption) (Commodity_SubscribeCommodityClient, error)
}

type commodityClient struct {
	cc grpc.ClientConnInterface
}

func NewCommodityClient(cc grpc.ClientConnInterface) CommodityClient {
	return &commodityClient{cc}
}

func (c *commodityClient) GetCommodity(ctx context.Context, in *CommodityRequest, opts ...grpc.CallOption) (*CommodityResponse, error) {
	out := new(CommodityResponse)
	err := c.cc.Invoke(ctx, "/Commodity/GetCommodity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commodityClient) SubscribeCommodity(ctx context.Context, opts ...grpc.CallOption) (Commodity_SubscribeCommodityClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Commodity_serviceDesc.Streams[0], "/Commodity/SubscribeCommodity", opts...)
	if err != nil {
		return nil, err
	}
	x := &commoditySubscribeCommodityClient{stream}
	return x, nil
}

type Commodity_SubscribeCommodityClient interface {
	Send(*CommodityRequest) error
	Recv() (*StreamingCommodityResponse, error)
	grpc.ClientStream
}

type commoditySubscribeCommodityClient struct {
	grpc.ClientStream
}

func (x *commoditySubscribeCommodityClient) Send(m *CommodityRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commoditySubscribeCommodityClient) Recv() (*StreamingCommodityResponse, error) {
	m := new(StreamingCommodityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommodityServer is the server API for Commodity service.
type CommodityServer interface {
	GetCommodity(context.Context, *CommodityRequest) (*CommodityResponse, error)
	SubscribeCommodity(Commodity_SubscribeCommodityServer) error
}

// UnimplementedCommodityServer can be embedded to have forward compatible implementations.
type UnimplementedCommodityServer struct {
}

func (*UnimplementedCommodityServer) GetCommodity(context.Context, *CommodityRequest) (*CommodityResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCommodity not implemented")
}
func (*UnimplementedCommodityServer) SubscribeCommodity(Commodity_SubscribeCommodityServer) error {
	return status1.Errorf(codes.Unimplemented, "method SubscribeCommodity not implemented")
}

func RegisterCommodityServer(s *grpc.Server, srv CommodityServer) {
	s.RegisterService(&_Commodity_serviceDesc, srv)
}

func _Commodity_GetCommodity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommodityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServer).GetCommodity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Commodity/GetCommodity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServer).GetCommodity(ctx, req.(*CommodityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commodity_SubscribeCommodity_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommodityServer).SubscribeCommodity(&commoditySubscribeCommodityServer{stream})
}

type Commodity_SubscribeCommodityServer interface {
	Send(*StreamingCommodityResponse) error
	Recv() (*CommodityRequest, error)
	grpc.ServerStream
}

type commoditySubscribeCommodityServer struct {
	grpc.ServerStream
}

func (x *commoditySubscribeCommodityServer) Send(m *StreamingCommodityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commoditySubscribeCommodityServer) Recv() (*CommodityRequest, error) {
	m := new(CommodityRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Commodity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Commodity",
	HandlerType: (*CommodityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommodity",
			Handler:    _Commodity_GetCommodity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeCommodity",
			Handler:       _Commodity_SubscribeCommodity_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "commodity.proto",
}
