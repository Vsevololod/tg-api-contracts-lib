// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: messages/messages.proto

package messagesv1

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

type MessageType int32

const (
	MessageType_UNKNOWN MessageType = 0 // Неизвестный тип сообщения (по умолчанию)
	MessageType_TEXT    MessageType = 1 // Текстовое сообщение
	MessageType_VIDEO   MessageType = 2 // Видео сообщение
	MessageType_IMAGE   MessageType = 3 // Сообщение с изображением
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "UNKNOWN",
		1: "TEXT",
		2: "VIDEO",
		3: "IMAGE",
	}
	MessageType_value = map[string]int32{
		"UNKNOWN": 0,
		"TEXT":    1,
		"VIDEO":   2,
		"IMAGE":   3,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_messages_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_messages_messages_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_messages_messages_proto_rawDescGZIP(), []int{0}
}

// Сообщение
type TgSendMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`                                                                               // Текст сообщения
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                                            // Идентификатор пользователя
	Type          MessageType            `protobuf:"varint,3,opt,name=type,proto3,enum=messages.MessageType" json:"type,omitempty"`                                                    // Тип сообщения (используем enum)
	Params        map[string]string      `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Параметры в виде ключ-значение
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TgSendMessage) Reset() {
	*x = TgSendMessage{}
	mi := &file_messages_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TgSendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TgSendMessage) ProtoMessage() {}

func (x *TgSendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TgSendMessage.ProtoReflect.Descriptor instead.
func (*TgSendMessage) Descriptor() ([]byte, []int) {
	return file_messages_messages_proto_rawDescGZIP(), []int{0}
}

func (x *TgSendMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TgSendMessage) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TgSendMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_UNKNOWN
}

func (x *TgSendMessage) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_messages_messages_proto protoreflect.FileDescriptor

var file_messages_messages_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0d, 0x54, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x3a, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56,
	0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10,
	0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x76, 0x73, 0x65, 0x76, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_messages_messages_proto_rawDescOnce sync.Once
	file_messages_messages_proto_rawDescData []byte
)

func file_messages_messages_proto_rawDescGZIP() []byte {
	file_messages_messages_proto_rawDescOnce.Do(func() {
		file_messages_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_messages_messages_proto_rawDesc), len(file_messages_messages_proto_rawDesc)))
	})
	return file_messages_messages_proto_rawDescData
}

var file_messages_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_messages_proto_goTypes = []any{
	(MessageType)(0),      // 0: messages.MessageType
	(*TgSendMessage)(nil), // 1: messages.TgSendMessage
	nil,                   // 2: messages.TgSendMessage.ParamsEntry
}
var file_messages_messages_proto_depIdxs = []int32{
	0, // 0: messages.TgSendMessage.type:type_name -> messages.MessageType
	2, // 1: messages.TgSendMessage.params:type_name -> messages.TgSendMessage.ParamsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_messages_proto_init() }
func file_messages_messages_proto_init() {
	if File_messages_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_messages_messages_proto_rawDesc), len(file_messages_messages_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_messages_proto_goTypes,
		DependencyIndexes: file_messages_messages_proto_depIdxs,
		EnumInfos:         file_messages_messages_proto_enumTypes,
		MessageInfos:      file_messages_messages_proto_msgTypes,
	}.Build()
	File_messages_messages_proto = out.File
	file_messages_messages_proto_goTypes = nil
	file_messages_messages_proto_depIdxs = nil
}
