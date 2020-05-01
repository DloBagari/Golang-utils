// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.5.1
// source: protocol_buffer/person.proto

package person

import (
	proto "github.com/golang/protobuf/proto"
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

type Person_EyeColour int32

const (
	Person_NOT_DEFINED Person_EyeColour = 0
	Person_BLACK       Person_EyeColour = 1
	Person_GREEN       Person_EyeColour = 2
	Person_BLUE        Person_EyeColour = 3
)

// Enum value maps for Person_EyeColour.
var (
	Person_EyeColour_name = map[int32]string{
		0: "NOT_DEFINED",
		1: "BLACK",
		2: "GREEN",
		3: "BLUE",
	}
	Person_EyeColour_value = map[string]int32{
		"NOT_DEFINED": 0,
		"BLACK":       1,
		"GREEN":       2,
		"BLUE":        3,
	}
)

func (x Person_EyeColour) Enum() *Person_EyeColour {
	p := new(Person_EyeColour)
	*p = x
	return p
}

func (x Person_EyeColour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_EyeColour) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_buffer_person_proto_enumTypes[0].Descriptor()
}

func (Person_EyeColour) Type() protoreflect.EnumType {
	return &file_protocol_buffer_person_proto_enumTypes[0]
}

func (x Person_EyeColour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_EyeColour.Descriptor instead.
func (Person_EyeColour) EnumDescriptor() ([]byte, []int) {
	return file_protocol_buffer_person_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age       int32      `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Developer bool       `protobuf:"varint,3,opt,name=developer,proto3" json:"developer,omitempty"`
	MyAddress []*Address `protobuf:"bytes,4,rep,name=my_address,json=myAddress,proto3" json:"my_address,omitempty"`
	// default value is EyeColour tag 0
	EyeColour Person_EyeColour `protobuf:"varint,6,opt,name=eye_colour,json=eyeColour,proto3,enum=Person_EyeColour" json:"eye_colour,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_buffer_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_buffer_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_protocol_buffer_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetDeveloper() bool {
	if x != nil {
		return x.Developer
	}
	return false
}

func (x *Person) GetMyAddress() []*Address {
	if x != nil {
		return x.MyAddress
	}
	return nil
}

func (x *Person) GetEyeColour() Person_EyeColour {
	if x != nil {
		return x.EyeColour
	}
	return Person_NOT_DEFINED
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street      string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	HouseNumber int32  `protobuf:"varint,2,opt,name=house_number,json=houseNumber,proto3" json:"house_number,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_buffer_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_buffer_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_protocol_buffer_person_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetHouseNumber() int32 {
	if x != nil {
		return x.HouseNumber
	}
	return 0
}

var File_protocol_buffer_person_proto protoreflect.FileDescriptor

var file_protocol_buffer_person_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5,
	0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x0a, 0x6d, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x6d, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x65, 0x79, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x75, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x45, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x52, 0x09, 0x65,
	0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x22, 0x3c, 0x0a, 0x09, 0x45, 0x79, 0x65, 0x43,
	0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x22, 0x44, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x18, 0x5a, 0x16,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x3b,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_buffer_person_proto_rawDescOnce sync.Once
	file_protocol_buffer_person_proto_rawDescData = file_protocol_buffer_person_proto_rawDesc
)

func file_protocol_buffer_person_proto_rawDescGZIP() []byte {
	file_protocol_buffer_person_proto_rawDescOnce.Do(func() {
		file_protocol_buffer_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_buffer_person_proto_rawDescData)
	})
	return file_protocol_buffer_person_proto_rawDescData
}

var file_protocol_buffer_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_buffer_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_buffer_person_proto_goTypes = []interface{}{
	(Person_EyeColour)(0), // 0: Person.EyeColour
	(*Person)(nil),        // 1: Person
	(*Address)(nil),       // 2: Address
}
var file_protocol_buffer_person_proto_depIdxs = []int32{
	2, // 0: Person.my_address:type_name -> Address
	0, // 1: Person.eye_colour:type_name -> Person.EyeColour
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_buffer_person_proto_init() }
func file_protocol_buffer_person_proto_init() {
	if File_protocol_buffer_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_buffer_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_protocol_buffer_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_protocol_buffer_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_buffer_person_proto_goTypes,
		DependencyIndexes: file_protocol_buffer_person_proto_depIdxs,
		EnumInfos:         file_protocol_buffer_person_proto_enumTypes,
		MessageInfos:      file_protocol_buffer_person_proto_msgTypes,
	}.Build()
	File_protocol_buffer_person_proto = out.File
	file_protocol_buffer_person_proto_rawDesc = nil
	file_protocol_buffer_person_proto_goTypes = nil
	file_protocol_buffer_person_proto_depIdxs = nil
}