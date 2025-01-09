// Copyright 2021-2024 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file tests varying casing of the service name and method name.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: v1beta1service.proto

package v1beta1service

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

type GetExample_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetExample_Request) Reset() {
	*x = GetExample_Request{}
	mi := &file_v1beta1service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExample_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExample_Request) ProtoMessage() {}

func (x *GetExample_Request) ProtoReflect() protoreflect.Message {
	mi := &file_v1beta1service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExample_Request.ProtoReflect.Descriptor instead.
func (*GetExample_Request) Descriptor() ([]byte, []int) {
	return file_v1beta1service_proto_rawDescGZIP(), []int{0}
}

type Get1ExampleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get1ExampleResponse) Reset() {
	*x = Get1ExampleResponse{}
	mi := &file_v1beta1service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get1ExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get1ExampleResponse) ProtoMessage() {}

func (x *Get1ExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1beta1service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get1ExampleResponse.ProtoReflect.Descriptor instead.
func (*Get1ExampleResponse) Descriptor() ([]byte, []int) {
	return file_v1beta1service_proto_rawDescGZIP(), []int{1}
}

var File_v1beta1service_proto protoreflect.FileDescriptor

var file_v1beta1service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x31, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x57, 0x0a,
	0x0d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x12, 0x46,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x31, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa8, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x13, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0xe2, 0x02, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1beta1service_proto_rawDescOnce sync.Once
	file_v1beta1service_proto_rawDescData = file_v1beta1service_proto_rawDesc
)

func file_v1beta1service_proto_rawDescGZIP() []byte {
	file_v1beta1service_proto_rawDescOnce.Do(func() {
		file_v1beta1service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1beta1service_proto_rawDescData)
	})
	return file_v1beta1service_proto_rawDescData
}

var file_v1beta1service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1beta1service_proto_goTypes = []any{
	(*GetExample_Request)(nil),  // 0: example.GetExample_Request
	(*Get1ExampleResponse)(nil), // 1: example.Get1example_response
}
var file_v1beta1service_proto_depIdxs = []int32{
	0, // 0: example.ExampleV1beta.Method:input_type -> example.GetExample_Request
	1, // 1: example.ExampleV1beta.Method:output_type -> example.Get1example_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1beta1service_proto_init() }
func file_v1beta1service_proto_init() {
	if File_v1beta1service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1beta1service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1beta1service_proto_goTypes,
		DependencyIndexes: file_v1beta1service_proto_depIdxs,
		MessageInfos:      file_v1beta1service_proto_msgTypes,
	}.Build()
	File_v1beta1service_proto = out.File
	file_v1beta1service_proto_rawDesc = nil
	file_v1beta1service_proto_goTypes = nil
	file_v1beta1service_proto_depIdxs = nil
}