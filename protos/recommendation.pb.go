// fork from https://github.com/grpc/grpc/blob/master/examples/protos/helloworld.proto
//
// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: protos/recommendation.proto

package protos

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

type Recommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid               uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PromotionMessages string `protobuf:"bytes,101,opt,name=promotion_messages,json=promotionMessages,proto3" json:"promotion_messages,omitempty"`
	Timestamp         int64  `protobuf:"varint,1000,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Recommendation) Reset() {
	*x = Recommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_recommendation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recommendation) ProtoMessage() {}

func (x *Recommendation) ProtoReflect() protoreflect.Message {
	mi := &file_protos_recommendation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recommendation.ProtoReflect.Descriptor instead.
func (*Recommendation) Descriptor() ([]byte, []int) {
	return file_protos_recommendation_proto_rawDescGZIP(), []int{0}
}

func (x *Recommendation) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Recommendation) GetPromotionMessages() string {
	if x != nil {
		return x.PromotionMessages
	}
	return ""
}

func (x *Recommendation) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetRecommendationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetRecommendationReq) Reset() {
	*x = GetRecommendationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_recommendation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecommendationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecommendationReq) ProtoMessage() {}

func (x *GetRecommendationReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_recommendation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecommendationReq.ProtoReflect.Descriptor instead.
func (*GetRecommendationReq) Descriptor() ([]byte, []int) {
	return file_protos_recommendation_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecommendationReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetRecommendationRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"` // 0 = successed, 1 = failed ...
	Reason string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`  // reason for status
	List   []*Recommendation `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetRecommendationRep) Reset() {
	*x = GetRecommendationRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_recommendation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecommendationRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecommendationRep) ProtoMessage() {}

func (x *GetRecommendationRep) ProtoReflect() protoreflect.Message {
	mi := &file_protos_recommendation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecommendationRep.ProtoReflect.Descriptor instead.
func (*GetRecommendationRep) Descriptor() ([]byte, []int) {
	return file_protos_recommendation_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecommendationRep) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetRecommendationRep) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetRecommendationRep) GetList() []*Recommendation {
	if x != nil {
		return x.List
	}
	return nil
}

var File_protos_recommendation_proto protoreflect.FileDescriptor

var file_protos_recommendation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a,
	0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0xe8, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x59, 0x0a, 0x12, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_recommendation_proto_rawDescOnce sync.Once
	file_protos_recommendation_proto_rawDescData = file_protos_recommendation_proto_rawDesc
)

func file_protos_recommendation_proto_rawDescGZIP() []byte {
	file_protos_recommendation_proto_rawDescOnce.Do(func() {
		file_protos_recommendation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_recommendation_proto_rawDescData)
	})
	return file_protos_recommendation_proto_rawDescData
}

var file_protos_recommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_recommendation_proto_goTypes = []interface{}{
	(*Recommendation)(nil),       // 0: Recommendation
	(*GetRecommendationReq)(nil), // 1: GetRecommendationReq
	(*GetRecommendationRep)(nil), // 2: GetRecommendationRep
}
var file_protos_recommendation_proto_depIdxs = []int32{
	0, // 0: GetRecommendationRep.list:type_name -> Recommendation
	1, // 1: RecommendationServ.GetRecommendation:input_type -> GetRecommendationReq
	2, // 2: RecommendationServ.GetRecommendation:output_type -> GetRecommendationRep
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_recommendation_proto_init() }
func file_protos_recommendation_proto_init() {
	if File_protos_recommendation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_recommendation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recommendation); i {
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
		file_protos_recommendation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecommendationReq); i {
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
		file_protos_recommendation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecommendationRep); i {
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
			RawDescriptor: file_protos_recommendation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_recommendation_proto_goTypes,
		DependencyIndexes: file_protos_recommendation_proto_depIdxs,
		MessageInfos:      file_protos_recommendation_proto_msgTypes,
	}.Build()
	File_protos_recommendation_proto = out.File
	file_protos_recommendation_proto_rawDesc = nil
	file_protos_recommendation_proto_goTypes = nil
	file_protos_recommendation_proto_depIdxs = nil
}
