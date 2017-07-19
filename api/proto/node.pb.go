// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	RegistrationEntry
	SvidEntry
	SpiffeMap
	SpiffeEntry
	AttestedData
	FetchBootstrapSVIDRequest
	FetchBootstrapSVIDResponse
	FetchNodeSVIDRequest
	FetchNodeSVIDResponse
	FetchSVIDRequest
	FetchSVIDResponse
	FetchCPBundleRequest
	FetchCPBundleResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RegistrationEntry struct {
	SelectorType    string `protobuf:"bytes,1,opt,name=selectorType" json:"selectorType,omitempty"`
	Selector        string `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
	Attestor        string `protobuf:"bytes,3,opt,name=attestor" json:"attestor,omitempty"`
	SpiffeId        string `protobuf:"bytes,4,opt,name=spiffeId" json:"spiffeId,omitempty"`
	FederatedBundle []byte `protobuf:"bytes,5,opt,name=federatedBundle,proto3" json:"federatedBundle,omitempty"`
	Ttl             int32  `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *RegistrationEntry) Reset()                    { *m = RegistrationEntry{} }
func (m *RegistrationEntry) String() string            { return proto1.CompactTextString(m) }
func (*RegistrationEntry) ProtoMessage()               {}
func (*RegistrationEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistrationEntry) GetSelectorType() string {
	if m != nil {
		return m.SelectorType
	}
	return ""
}

func (m *RegistrationEntry) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *RegistrationEntry) GetAttestor() string {
	if m != nil {
		return m.Attestor
	}
	return ""
}

func (m *RegistrationEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *RegistrationEntry) GetFederatedBundle() []byte {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

func (m *RegistrationEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type SvidEntry struct {
	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	Ttl  int32  `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *SvidEntry) Reset()                    { *m = SvidEntry{} }
func (m *SvidEntry) String() string            { return proto1.CompactTextString(m) }
func (*SvidEntry) ProtoMessage()               {}
func (*SvidEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SvidEntry) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *SvidEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type SpiffeMap struct {
	Map map[string]*SvidEntry `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SpiffeMap) Reset()                    { *m = SpiffeMap{} }
func (m *SpiffeMap) String() string            { return proto1.CompactTextString(m) }
func (*SpiffeMap) ProtoMessage()               {}
func (*SpiffeMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpiffeMap) GetMap() map[string]*SvidEntry {
	if m != nil {
		return m.Map
	}
	return nil
}

type SpiffeEntry struct {
	SpiffeMap             *SpiffeMap           `protobuf:"bytes,1,opt,name=spiffeMap" json:"spiffeMap,omitempty"`
	RegistrationEntryList []*RegistrationEntry `protobuf:"bytes,2,rep,name=registrationEntryList" json:"registrationEntryList,omitempty"`
}

func (m *SpiffeEntry) Reset()                    { *m = SpiffeEntry{} }
func (m *SpiffeEntry) String() string            { return proto1.CompactTextString(m) }
func (*SpiffeEntry) ProtoMessage()               {}
func (*SpiffeEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SpiffeEntry) GetSpiffeMap() *SpiffeMap {
	if m != nil {
		return m.SpiffeMap
	}
	return nil
}

func (m *SpiffeEntry) GetRegistrationEntryList() []*RegistrationEntry {
	if m != nil {
		return m.RegistrationEntryList
	}
	return nil
}

type AttestedData struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *AttestedData) Reset()                    { *m = AttestedData{} }
func (m *AttestedData) String() string            { return proto1.CompactTextString(m) }
func (*AttestedData) ProtoMessage()               {}
func (*AttestedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AttestedData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestedData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type FetchBootstrapSVIDRequest struct {
	AttestedData *AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	Csr          []byte        `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *FetchBootstrapSVIDRequest) Reset()                    { *m = FetchBootstrapSVIDRequest{} }
func (m *FetchBootstrapSVIDRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchBootstrapSVIDRequest) ProtoMessage()               {}
func (*FetchBootstrapSVIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchBootstrapSVIDRequest) GetAttestedData() *AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *FetchBootstrapSVIDRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type FetchBootstrapSVIDResponse struct {
	SpiffeEntry *SpiffeEntry `protobuf:"bytes,1,opt,name=spiffeEntry" json:"spiffeEntry,omitempty"`
	Nonce       int64        `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *FetchBootstrapSVIDResponse) Reset()                    { *m = FetchBootstrapSVIDResponse{} }
func (m *FetchBootstrapSVIDResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchBootstrapSVIDResponse) ProtoMessage()               {}
func (*FetchBootstrapSVIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchBootstrapSVIDResponse) GetSpiffeEntry() *SpiffeEntry {
	if m != nil {
		return m.SpiffeEntry
	}
	return nil
}

func (m *FetchBootstrapSVIDResponse) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type FetchNodeSVIDRequest struct {
	Nonce   int64    `protobuf:"varint,1,opt,name=nonce" json:"nonce,omitempty"`
	CsrList [][]byte `protobuf:"bytes,2,rep,name=csrList,proto3" json:"csrList,omitempty"`
}

func (m *FetchNodeSVIDRequest) Reset()                    { *m = FetchNodeSVIDRequest{} }
func (m *FetchNodeSVIDRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchNodeSVIDRequest) ProtoMessage()               {}
func (*FetchNodeSVIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FetchNodeSVIDRequest) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *FetchNodeSVIDRequest) GetCsrList() [][]byte {
	if m != nil {
		return m.CsrList
	}
	return nil
}

type FetchNodeSVIDResponse struct {
	SpiffeEntry *SpiffeEntry `protobuf:"bytes,1,opt,name=spiffeEntry" json:"spiffeEntry,omitempty"`
}

func (m *FetchNodeSVIDResponse) Reset()                    { *m = FetchNodeSVIDResponse{} }
func (m *FetchNodeSVIDResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchNodeSVIDResponse) ProtoMessage()               {}
func (*FetchNodeSVIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FetchNodeSVIDResponse) GetSpiffeEntry() *SpiffeEntry {
	if m != nil {
		return m.SpiffeEntry
	}
	return nil
}

type FetchSVIDRequest struct {
	CsrList [][]byte `protobuf:"bytes,2,rep,name=csrList,proto3" json:"csrList,omitempty"`
}

func (m *FetchSVIDRequest) Reset()                    { *m = FetchSVIDRequest{} }
func (m *FetchSVIDRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchSVIDRequest) ProtoMessage()               {}
func (*FetchSVIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FetchSVIDRequest) GetCsrList() [][]byte {
	if m != nil {
		return m.CsrList
	}
	return nil
}

type FetchSVIDResponse struct {
	SpiffeEntry *SpiffeEntry `protobuf:"bytes,1,opt,name=spiffeEntry" json:"spiffeEntry,omitempty"`
}

func (m *FetchSVIDResponse) Reset()                    { *m = FetchSVIDResponse{} }
func (m *FetchSVIDResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchSVIDResponse) ProtoMessage()               {}
func (*FetchSVIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FetchSVIDResponse) GetSpiffeEntry() *SpiffeEntry {
	if m != nil {
		return m.SpiffeEntry
	}
	return nil
}

type FetchCPBundleRequest struct {
}

func (m *FetchCPBundleRequest) Reset()                    { *m = FetchCPBundleRequest{} }
func (m *FetchCPBundleRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchCPBundleRequest) ProtoMessage()               {}
func (*FetchCPBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type FetchCPBundleResponse struct {
	CpBundle []byte `protobuf:"bytes,1,opt,name=cpBundle,proto3" json:"cpBundle,omitempty"`
}

func (m *FetchCPBundleResponse) Reset()                    { *m = FetchCPBundleResponse{} }
func (m *FetchCPBundleResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchCPBundleResponse) ProtoMessage()               {}
func (*FetchCPBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FetchCPBundleResponse) GetCpBundle() []byte {
	if m != nil {
		return m.CpBundle
	}
	return nil
}

func init() {
	proto1.RegisterType((*RegistrationEntry)(nil), "proto.RegistrationEntry")
	proto1.RegisterType((*SvidEntry)(nil), "proto.SvidEntry")
	proto1.RegisterType((*SpiffeMap)(nil), "proto.SpiffeMap")
	proto1.RegisterType((*SpiffeEntry)(nil), "proto.SpiffeEntry")
	proto1.RegisterType((*AttestedData)(nil), "proto.AttestedData")
	proto1.RegisterType((*FetchBootstrapSVIDRequest)(nil), "proto.FetchBootstrapSVIDRequest")
	proto1.RegisterType((*FetchBootstrapSVIDResponse)(nil), "proto.FetchBootstrapSVIDResponse")
	proto1.RegisterType((*FetchNodeSVIDRequest)(nil), "proto.FetchNodeSVIDRequest")
	proto1.RegisterType((*FetchNodeSVIDResponse)(nil), "proto.FetchNodeSVIDResponse")
	proto1.RegisterType((*FetchSVIDRequest)(nil), "proto.FetchSVIDRequest")
	proto1.RegisterType((*FetchSVIDResponse)(nil), "proto.FetchSVIDResponse")
	proto1.RegisterType((*FetchCPBundleRequest)(nil), "proto.FetchCPBundleRequest")
	proto1.RegisterType((*FetchCPBundleResponse)(nil), "proto.FetchCPBundleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	FetchBootstrapSVID(ctx context.Context, in *FetchBootstrapSVIDRequest, opts ...grpc.CallOption) (*FetchBootstrapSVIDResponse, error)
	FetchNodeSVID(ctx context.Context, in *FetchNodeSVIDRequest, opts ...grpc.CallOption) (*FetchNodeSVIDResponse, error)
	FetchSVID(ctx context.Context, in *FetchSVIDRequest, opts ...grpc.CallOption) (*FetchSVIDResponse, error)
	FetchCPBundle(ctx context.Context, in *FetchCPBundleRequest, opts ...grpc.CallOption) (*FetchCPBundleResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) FetchBootstrapSVID(ctx context.Context, in *FetchBootstrapSVIDRequest, opts ...grpc.CallOption) (*FetchBootstrapSVIDResponse, error) {
	out := new(FetchBootstrapSVIDResponse)
	err := grpc.Invoke(ctx, "/proto.node/FetchBootstrapSVID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) FetchNodeSVID(ctx context.Context, in *FetchNodeSVIDRequest, opts ...grpc.CallOption) (*FetchNodeSVIDResponse, error) {
	out := new(FetchNodeSVIDResponse)
	err := grpc.Invoke(ctx, "/proto.node/FetchNodeSVID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) FetchSVID(ctx context.Context, in *FetchSVIDRequest, opts ...grpc.CallOption) (*FetchSVIDResponse, error) {
	out := new(FetchSVIDResponse)
	err := grpc.Invoke(ctx, "/proto.node/FetchSVID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) FetchCPBundle(ctx context.Context, in *FetchCPBundleRequest, opts ...grpc.CallOption) (*FetchCPBundleResponse, error) {
	out := new(FetchCPBundleResponse)
	err := grpc.Invoke(ctx, "/proto.node/FetchCPBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	FetchBootstrapSVID(context.Context, *FetchBootstrapSVIDRequest) (*FetchBootstrapSVIDResponse, error)
	FetchNodeSVID(context.Context, *FetchNodeSVIDRequest) (*FetchNodeSVIDResponse, error)
	FetchSVID(context.Context, *FetchSVIDRequest) (*FetchSVIDResponse, error)
	FetchCPBundle(context.Context, *FetchCPBundleRequest) (*FetchCPBundleResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_FetchBootstrapSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBootstrapSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchBootstrapSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/FetchBootstrapSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchBootstrapSVID(ctx, req.(*FetchBootstrapSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_FetchNodeSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchNodeSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchNodeSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/FetchNodeSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchNodeSVID(ctx, req.(*FetchNodeSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_FetchSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/FetchSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchSVID(ctx, req.(*FetchSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_FetchCPBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCPBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchCPBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/FetchCPBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchCPBundle(ctx, req.(*FetchCPBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchBootstrapSVID",
			Handler:    _Node_FetchBootstrapSVID_Handler,
		},
		{
			MethodName: "FetchNodeSVID",
			Handler:    _Node_FetchNodeSVID_Handler,
		},
		{
			MethodName: "FetchSVID",
			Handler:    _Node_FetchSVID_Handler,
		},
		{
			MethodName: "FetchCPBundle",
			Handler:    _Node_FetchCPBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto1.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xf3, 0x53, 0x92, 0x89, 0x11, 0xe9, 0xd2, 0x82, 0x6b, 0x38, 0x84, 0x3d, 0xa0, 0x48,
	0xa0, 0x48, 0xa4, 0x08, 0x10, 0x07, 0x24, 0x4a, 0xa9, 0x08, 0x22, 0x15, 0xda, 0x20, 0x24, 0x8e,
	0x8b, 0x3d, 0xa1, 0x11, 0xc1, 0x36, 0xde, 0x4d, 0xa5, 0x5c, 0xb9, 0xf3, 0x40, 0x3c, 0x01, 0xaf,
	0x85, 0xf6, 0xc7, 0xce, 0xda, 0x49, 0xb9, 0xf4, 0x94, 0xf9, 0xfd, 0xbe, 0x99, 0x6f, 0xd6, 0x01,
	0x48, 0xd2, 0x18, 0x47, 0x59, 0x9e, 0xca, 0x94, 0xb4, 0xf5, 0x0f, 0xfd, 0xeb, 0xc1, 0x3e, 0xc3,
	0x6f, 0x0b, 0x21, 0x73, 0x2e, 0x17, 0x69, 0xf2, 0x36, 0x91, 0xf9, 0x9a, 0x50, 0xf0, 0x05, 0x2e,
	0x31, 0x92, 0x69, 0xfe, 0x69, 0x9d, 0x61, 0xe0, 0x0d, 0xbc, 0x61, 0x97, 0x55, 0x62, 0x24, 0x84,
	0x4e, 0xe1, 0x07, 0x0d, 0x9d, 0x2f, 0x7d, 0x95, 0xe3, 0x52, 0xa2, 0x50, 0xb9, 0xa6, 0xc9, 0x15,
	0xbe, 0xee, 0xcb, 0x16, 0xf3, 0x39, 0x4e, 0xe2, 0xa0, 0x65, 0xfb, 0xac, 0x4f, 0x86, 0x70, 0x6b,
	0x8e, 0x31, 0xe6, 0x5c, 0x62, 0x7c, 0xb2, 0x4a, 0xe2, 0x25, 0x06, 0xed, 0x81, 0x37, 0xf4, 0x59,
	0x3d, 0x4c, 0xfa, 0xd0, 0x94, 0x72, 0x19, 0xec, 0x0d, 0xbc, 0x61, 0x9b, 0x29, 0x93, 0x3e, 0x81,
	0xee, 0xec, 0x72, 0x11, 0x9b, 0x05, 0x08, 0xb4, 0x22, 0xcc, 0xa5, 0x1e, 0xdc, 0x67, 0xda, 0x2e,
	0x5a, 0x1a, 0x9b, 0x96, 0x5f, 0x1e, 0x74, 0x67, 0x9a, 0x7b, 0xca, 0x33, 0xf2, 0x08, 0x9a, 0x3f,
	0x78, 0x16, 0x78, 0x83, 0xe6, 0xb0, 0x37, 0x3e, 0x32, 0x32, 0x8d, 0xca, 0xf4, 0x68, 0xca, 0x33,
	0x8d, 0xcd, 0x54, 0x55, 0xf8, 0x0e, 0x3a, 0x45, 0x40, 0x01, 0x7f, 0xc7, 0xb5, 0x15, 0x49, 0x99,
	0xe4, 0x21, 0xb4, 0x2f, 0xf9, 0x72, 0x85, 0x9a, 0xac, 0x37, 0xee, 0x17, 0x60, 0xc5, 0x7c, 0xcc,
	0xa4, 0x5f, 0x36, 0x5e, 0x78, 0xf4, 0xb7, 0x07, 0x3d, 0xc3, 0x62, 0xd0, 0x46, 0xd0, 0x15, 0x05,
	0xa9, 0xc6, 0x74, 0xfa, 0x8b, 0x38, 0xdb, 0x94, 0x90, 0x73, 0x38, 0xcc, 0xeb, 0x07, 0xfc, 0xb0,
	0x10, 0x32, 0x68, 0xe8, 0x45, 0x02, 0xdb, 0xbb, 0x75, 0x64, 0xb6, 0xbb, 0x8d, 0x3e, 0x03, 0xff,
	0xb5, 0xbe, 0x15, 0xc6, 0xa7, 0x5c, 0x72, 0x25, 0xa5, 0xdc, 0xbc, 0x01, 0x6d, 0xab, 0x58, 0xcc,
	0x25, 0xb7, 0x77, 0xd7, 0x36, 0x9d, 0xc3, 0xd1, 0x19, 0xca, 0xe8, 0xe2, 0x24, 0x4d, 0xa5, 0x42,
	0xcd, 0x66, 0x9f, 0x27, 0xa7, 0x0c, 0x7f, 0xae, 0x50, 0x48, 0xf2, 0x1c, 0x7c, 0xee, 0x80, 0xda,
	0xbd, 0x6e, 0xdb, 0xd9, 0x5c, 0x3e, 0x56, 0x29, 0x54, 0xda, 0x46, 0xc2, 0x3c, 0x30, 0x9f, 0x29,
	0x93, 0x5e, 0x40, 0xb8, 0x8b, 0x47, 0x64, 0x69, 0x22, 0x90, 0x3c, 0x85, 0x9e, 0xd8, 0x88, 0x69,
	0x79, 0x48, 0x45, 0x3f, 0xb3, 0xbd, 0x5b, 0x46, 0x0e, 0xa0, 0x9d, 0xa4, 0x49, 0x84, 0xfa, 0xb1,
	0x36, 0x99, 0x71, 0xe8, 0x19, 0x1c, 0x68, 0xa6, 0xf3, 0x34, 0x46, 0x77, 0x99, 0xb2, 0xda, 0x73,
	0xaa, 0x49, 0x00, 0x37, 0x22, 0x91, 0x97, 0xca, 0xfb, 0xac, 0x70, 0xe9, 0x14, 0x0e, 0x6b, 0x38,
	0xd7, 0x19, 0x96, 0x3e, 0x86, 0xbe, 0x86, 0x73, 0x47, 0xba, 0x9a, 0x7c, 0x02, 0xfb, 0x4e, 0xf5,
	0xb5, 0x88, 0xef, 0x58, 0x3d, 0xde, 0x7c, 0x34, 0x1f, 0xa1, 0x25, 0xa7, 0xc7, 0x76, 0xbf, 0x4d,
	0xdc, 0xd2, 0x84, 0xd0, 0x89, 0x32, 0xfb, 0x1d, 0x9b, 0x2f, 0xb1, 0xf4, 0xc7, 0x7f, 0x1a, 0xd0,
	0x52, 0x7f, 0x47, 0xe4, 0x0b, 0x90, 0xed, 0x7b, 0x92, 0x81, 0x1d, 0xe6, 0xca, 0x27, 0x15, 0x3e,
	0xf8, 0x4f, 0x85, 0xe5, 0x7f, 0x0f, 0x37, 0x2b, 0xc2, 0x93, 0x7b, 0x6e, 0x4f, 0xed, 0xac, 0xe1,
	0xfd, 0xdd, 0x49, 0x8b, 0xf5, 0x0a, 0xba, 0xa5, 0x8e, 0xe4, 0xae, 0x5b, 0xea, 0x62, 0x04, 0xdb,
	0x89, 0xda, 0x2c, 0x85, 0x48, 0xd5, 0x59, 0x6a, 0x92, 0x56, 0x67, 0xa9, 0xeb, 0xfa, 0x75, 0x4f,
	0x27, 0x8f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x92, 0xc5, 0x66, 0xe9, 0xd0, 0x05, 0x00, 0x00,
}
