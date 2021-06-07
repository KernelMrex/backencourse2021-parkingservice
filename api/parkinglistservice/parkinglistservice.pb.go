// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: parkinglistservice.proto

package parkinglistservice

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

type GetParkingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetParkingRequest) Reset() {
	*x = GetParkingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglistservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParkingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParkingRequest) ProtoMessage() {}

func (x *GetParkingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglistservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParkingRequest.ProtoReflect.Descriptor instead.
func (*GetParkingRequest) Descriptor() ([]byte, []int) {
	return file_parkinglistservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetParkingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetParkingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Location       string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	AvailableSlots int64  `protobuf:"varint,4,opt,name=available_slots,json=availableSlots,proto3" json:"available_slots,omitempty"`
	ReservedSlots  int64  `protobuf:"varint,5,opt,name=reserved_slots,json=reservedSlots,proto3" json:"reserved_slots,omitempty"`
}

func (x *GetParkingResponse) Reset() {
	*x = GetParkingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglistservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParkingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParkingResponse) ProtoMessage() {}

func (x *GetParkingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglistservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParkingResponse.ProtoReflect.Descriptor instead.
func (*GetParkingResponse) Descriptor() ([]byte, []int) {
	return file_parkinglistservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetParkingResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetParkingResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetParkingResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetParkingResponse) GetAvailableSlots() int64 {
	if x != nil {
		return x.AvailableSlots
	}
	return 0
}

func (x *GetParkingResponse) GetReservedSlots() int64 {
	if x != nil {
		return x.ReservedSlots
	}
	return 0
}

var File_parkinglistservice_proto protoreflect.FileDescriptor

var file_parkinglistservice_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x32, 0x73,
	0x0a, 0x12, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parkinglistservice_proto_rawDescOnce sync.Once
	file_parkinglistservice_proto_rawDescData = file_parkinglistservice_proto_rawDesc
)

func file_parkinglistservice_proto_rawDescGZIP() []byte {
	file_parkinglistservice_proto_rawDescOnce.Do(func() {
		file_parkinglistservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_parkinglistservice_proto_rawDescData)
	})
	return file_parkinglistservice_proto_rawDescData
}

var file_parkinglistservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_parkinglistservice_proto_goTypes = []interface{}{
	(*GetParkingRequest)(nil),  // 0: parkinglistservice.GetParkingRequest
	(*GetParkingResponse)(nil), // 1: parkinglistservice.GetParkingResponse
}
var file_parkinglistservice_proto_depIdxs = []int32{
	0, // 0: parkinglistservice.ParkingListService.GetParking:input_type -> parkinglistservice.GetParkingRequest
	1, // 1: parkinglistservice.ParkingListService.GetParking:output_type -> parkinglistservice.GetParkingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_parkinglistservice_proto_init() }
func file_parkinglistservice_proto_init() {
	if File_parkinglistservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parkinglistservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParkingRequest); i {
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
		file_parkinglistservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParkingResponse); i {
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
			RawDescriptor: file_parkinglistservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parkinglistservice_proto_goTypes,
		DependencyIndexes: file_parkinglistservice_proto_depIdxs,
		MessageInfos:      file_parkinglistservice_proto_msgTypes,
	}.Build()
	File_parkinglistservice_proto = out.File
	file_parkinglistservice_proto_rawDesc = nil
	file_parkinglistservice_proto_goTypes = nil
	file_parkinglistservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParkingListServiceClient is the client API for ParkingListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParkingListServiceClient interface {
	GetParking(ctx context.Context, in *GetParkingRequest, opts ...grpc.CallOption) (*GetParkingResponse, error)
}

type parkingListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParkingListServiceClient(cc grpc.ClientConnInterface) ParkingListServiceClient {
	return &parkingListServiceClient{cc}
}

func (c *parkingListServiceClient) GetParking(ctx context.Context, in *GetParkingRequest, opts ...grpc.CallOption) (*GetParkingResponse, error) {
	out := new(GetParkingResponse)
	err := c.cc.Invoke(ctx, "/parkinglistservice.ParkingListService/GetParking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkingListServiceServer is the server API for ParkingListService service.
type ParkingListServiceServer interface {
	GetParking(context.Context, *GetParkingRequest) (*GetParkingResponse, error)
}

// UnimplementedParkingListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParkingListServiceServer struct {
}

func (*UnimplementedParkingListServiceServer) GetParking(context.Context, *GetParkingRequest) (*GetParkingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParking not implemented")
}

func RegisterParkingListServiceServer(s *grpc.Server, srv ParkingListServiceServer) {
	s.RegisterService(&_ParkingListService_serviceDesc, srv)
}

func _ParkingListService_GetParking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParkingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingListServiceServer).GetParking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkinglistservice.ParkingListService/GetParking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingListServiceServer).GetParking(ctx, req.(*GetParkingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParkingListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parkinglistservice.ParkingListService",
	HandlerType: (*ParkingListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParking",
			Handler:    _ParkingListService_GetParking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parkinglistservice.proto",
}