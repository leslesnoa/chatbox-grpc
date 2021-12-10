// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/chatbox.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_chatbox_proto_rawDescGZIP(), []int{0}
}

type ChatData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ChatData) Reset() {
	*x = ChatData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatData) ProtoMessage() {}

func (x *ChatData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatData.ProtoReflect.Descriptor instead.
func (*ChatData) Descriptor() ([]byte, []int) {
	return file_proto_chatbox_proto_rawDescGZIP(), []int{1}
}

func (x *ChatData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ChatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatList []*ChatData `protobuf:"bytes,1,rep,name=chat_list,json=chatList,proto3" json:"chat_list,omitempty"`
}

func (x *ChatList) Reset() {
	*x = ChatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatList) ProtoMessage() {}

func (x *ChatList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatList.ProtoReflect.Descriptor instead.
func (*ChatList) Descriptor() ([]byte, []int) {
	return file_proto_chatbox_proto_rawDescGZIP(), []int{2}
}

func (x *ChatList) GetChatList() []*ChatData {
	if x != nil {
		return x.ChatList
	}
	return nil
}

var File_proto_chatbox_proto protoreflect.FileDescriptor

var file_proto_chatbox_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x6f, 0x78, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x97, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x42, 0x6f,
	0x78, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x11, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x78, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chatbox_proto_rawDescOnce sync.Once
	file_proto_chatbox_proto_rawDescData = file_proto_chatbox_proto_rawDesc
)

func file_proto_chatbox_proto_rawDescGZIP() []byte {
	file_proto_chatbox_proto_rawDescOnce.Do(func() {
		file_proto_chatbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chatbox_proto_rawDescData)
	})
	return file_proto_chatbox_proto_rawDescData
}

var file_proto_chatbox_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_chatbox_proto_goTypes = []interface{}{
	(*Empty)(nil),    // 0: chatbox.Empty
	(*ChatData)(nil), // 1: chatbox.ChatData
	(*ChatList)(nil), // 2: chatbox.ChatList
}
var file_proto_chatbox_proto_depIdxs = []int32{
	1, // 0: chatbox.ChatList.chat_list:type_name -> chatbox.ChatData
	1, // 1: chatbox.ChatBox.SendChat:input_type -> chatbox.ChatData
	0, // 2: chatbox.ChatBox.GetChat:input_type -> chatbox.Empty
	0, // 3: chatbox.ChatBox.SendChat:output_type -> chatbox.Empty
	2, // 4: chatbox.ChatBox.GetChat:output_type -> chatbox.ChatList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_chatbox_proto_init() }
func file_proto_chatbox_proto_init() {
	if File_proto_chatbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chatbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_chatbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatData); i {
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
		file_proto_chatbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatList); i {
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
			RawDescriptor: file_proto_chatbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chatbox_proto_goTypes,
		DependencyIndexes: file_proto_chatbox_proto_depIdxs,
		MessageInfos:      file_proto_chatbox_proto_msgTypes,
	}.Build()
	File_proto_chatbox_proto = out.File
	file_proto_chatbox_proto_rawDesc = nil
	file_proto_chatbox_proto_goTypes = nil
	file_proto_chatbox_proto_depIdxs = nil
}
