// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: person.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Person struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Street        string                 `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip           string                 `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

var File_person_proto protoreflect.FileDescriptor

const file_person_proto_rawDesc = "" +
	"\n" +
	"\fperson.proto\x12\aexample\".\n" +
	"\x06Person\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x02 \x01(\x05R\x03age\"]\n" +
	"\aAddress\x12\x16\n" +
	"\x06street\x18\x01 \x01(\tR\x06street\x12\x12\n" +
	"\x04city\x18\x02 \x01(\tR\x04city\x12\x14\n" +
	"\x05state\x18\x03 \x01(\tR\x05state\x12\x10\n" +
	"\x03zip\x18\x04 \x01(\tR\x03zipB+Z)github.com/natemollica-nm/protobuffin/genb\x06proto3"

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData []byte
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_person_proto_rawDesc), len(file_person_proto_rawDesc)))
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_person_proto_goTypes = []any{
	(*Person)(nil),  // 0: example.Person
	(*Address)(nil), // 1: example.Address
}
var file_person_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_person_proto_rawDesc), len(file_person_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
