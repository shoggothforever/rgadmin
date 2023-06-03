// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.19.4
// source: wage.proto

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

type CalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Jwt      string  `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	WorkHour float32 `protobuf:"fixed32,3,opt,name=workHour,proto3" json:"workHour,omitempty"`
}

func (x *CalReq) Reset() {
	*x = CalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalReq) ProtoMessage() {}

func (x *CalReq) ProtoReflect() protoreflect.Message {
	mi := &file_wage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalReq.ProtoReflect.Descriptor instead.
func (*CalReq) Descriptor() ([]byte, []int) {
	return file_wage_proto_rawDescGZIP(), []int{0}
}

func (x *CalReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CalReq) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *CalReq) GetWorkHour() float32 {
	if x != nil {
		return x.WorkHour
	}
	return 0
}

type CalResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WageBeforeTax float32 `protobuf:"fixed32,1,opt,name=wageBeforeTax,proto3" json:"wageBeforeTax,omitempty"`
	WageAfterTax  float32 `protobuf:"fixed32,2,opt,name=wageAfterTax,proto3" json:"wageAfterTax,omitempty"`
	ActualWage    float32 `protobuf:"fixed32,3,opt,name=actualWage,proto3" json:"actualWage,omitempty"`
}

func (x *CalResp) Reset() {
	*x = CalResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalResp) ProtoMessage() {}

func (x *CalResp) ProtoReflect() protoreflect.Message {
	mi := &file_wage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalResp.ProtoReflect.Descriptor instead.
func (*CalResp) Descriptor() ([]byte, []int) {
	return file_wage_proto_rawDescGZIP(), []int{1}
}

func (x *CalResp) GetWageBeforeTax() float32 {
	if x != nil {
		return x.WageBeforeTax
	}
	return 0
}

func (x *CalResp) GetWageAfterTax() float32 {
	if x != nil {
		return x.WageAfterTax
	}
	return 0
}

func (x *CalResp) GetActualWage() float32 {
	if x != nil {
		return x.ActualWage
	}
	return 0
}

var File_wage_proto protoreflect.FileDescriptor

var file_wage_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x46, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x22, 0x73, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x61, 0x67, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x54, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x77, 0x61, 0x67, 0x65,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x61, 0x67,
	0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0c, 0x77, 0x61, 0x67, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54, 0x61, 0x78, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x57, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x57, 0x61, 0x67, 0x65, 0x32, 0x2a, 0x0a,
	0x04, 0x77, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x57, 0x61, 0x67, 0x65,
	0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wage_proto_rawDescOnce sync.Once
	file_wage_proto_rawDescData = file_wage_proto_rawDesc
)

func file_wage_proto_rawDescGZIP() []byte {
	file_wage_proto_rawDescOnce.Do(func() {
		file_wage_proto_rawDescData = protoimpl.X.CompressGZIP(file_wage_proto_rawDescData)
	})
	return file_wage_proto_rawDescData
}

var file_wage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wage_proto_goTypes = []interface{}{
	(*CalReq)(nil),  // 0: pb.CalReq
	(*CalResp)(nil), // 1: pb.CalResp
}
var file_wage_proto_depIdxs = []int32{
	0, // 0: pb.wage.calWage:input_type -> pb.CalReq
	1, // 1: pb.wage.calWage:output_type -> pb.CalResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wage_proto_init() }
func file_wage_proto_init() {
	if File_wage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalReq); i {
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
		file_wage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalResp); i {
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
			RawDescriptor: file_wage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wage_proto_goTypes,
		DependencyIndexes: file_wage_proto_depIdxs,
		MessageInfos:      file_wage_proto_msgTypes,
	}.Build()
	File_wage_proto = out.File
	file_wage_proto_rawDesc = nil
	file_wage_proto_goTypes = nil
	file_wage_proto_depIdxs = nil
}