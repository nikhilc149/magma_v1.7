//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: protos/eap_auth.proto

package protos

import (
	context "context"
	reflect "reflect"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	protos "magma/feg/gateway/services/aaa/protos"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_eap_auth_proto protoreflect.FileDescriptor

var file_protos_eap_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x61, 0x70, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x61, 0x70, 0x1a, 0x09, 0x65, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3b, 0x0a, 0x0b, 0x65, 0x61, 0x70, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61,
	0x70, 0x1a, 0x0f, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65,
	0x61, 0x70, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66, 0x65,
	0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x65, 0x61, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_eap_auth_proto_goTypes = []interface{}{
	(*protos.Eap)(nil), // 0: aaa.protos.eap
}
var file_protos_eap_auth_proto_depIdxs = []int32{
	0, // 0: eap.eap_service.handle:input_type -> aaa.protos.eap
	0, // 1: eap.eap_service.handle:output_type -> aaa.protos.eap
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_eap_auth_proto_init() }
func file_protos_eap_auth_proto_init() {
	if File_protos_eap_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_eap_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_eap_auth_proto_goTypes,
		DependencyIndexes: file_protos_eap_auth_proto_depIdxs,
	}.Build()
	File_protos_eap_auth_proto = out.File
	file_protos_eap_auth_proto_rawDesc = nil
	file_protos_eap_auth_proto_goTypes = nil
	file_protos_eap_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EapServiceClient is the client API for EapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EapServiceClient interface {
	Handle(ctx context.Context, in *protos.Eap, opts ...grpc.CallOption) (*protos.Eap, error)
}

type eapServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEapServiceClient(cc grpc.ClientConnInterface) EapServiceClient {
	return &eapServiceClient{cc}
}

func (c *eapServiceClient) Handle(ctx context.Context, in *protos.Eap, opts ...grpc.CallOption) (*protos.Eap, error) {
	out := new(protos.Eap)
	err := c.cc.Invoke(ctx, "/eap.eap_service/handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EapServiceServer is the server API for EapService service.
type EapServiceServer interface {
	Handle(context.Context, *protos.Eap) (*protos.Eap, error)
}

// UnimplementedEapServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEapServiceServer struct {
}

func (*UnimplementedEapServiceServer) Handle(context.Context, *protos.Eap) (*protos.Eap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterEapServiceServer(s *grpc.Server, srv EapServiceServer) {
	s.RegisterService(&_EapService_serviceDesc, srv)
}

func _EapService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Eap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EapServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eap.eap_service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EapServiceServer).Handle(ctx, req.(*protos.Eap))
	}
	return interceptor(ctx, in, info, handler)
}

var _EapService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eap.eap_service",
	HandlerType: (*EapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handle",
			Handler:    _EapService_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/eap_auth.proto",
}
