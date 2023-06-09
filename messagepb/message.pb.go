// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: messagepb/message.proto

package messagepb

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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message1 string `protobuf:"bytes,1,opt,name=message1,proto3" json:"message1,omitempty"`
	Message2 string `protobuf:"bytes,2,opt,name=message2,proto3" json:"message2,omitempty"` //  string message3 = 10;
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetMessage1() string {
	if x != nil {
		return x.Message1
	}
	return ""
}

func (x *MessageRequest) GetMessage2() string {
	if x != nil {
		return x.Message2
	}
	return ""
}

type CombineMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CombineMessageResponse) Reset() {
	*x = CombineMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombineMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombineMessageResponse) ProtoMessage() {}

func (x *CombineMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombineMessageResponse.ProtoReflect.Descriptor instead.
func (*CombineMessageResponse) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{1}
}

func (x *CombineMessageResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type SplitMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SplitMessage) Reset() {
	*x = SplitMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitMessage) ProtoMessage() {}

func (x *SplitMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitMessage.ProtoReflect.Descriptor instead.
func (*SplitMessage) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{2}
}

func (x *SplitMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseWords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *ResponseWords) Reset() {
	*x = ResponseWords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseWords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseWords) ProtoMessage() {}

func (x *ResponseWords) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseWords.ProtoReflect.Descriptor instead.
func (*ResponseWords) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseWords) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type ParagraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *ParagraphRequest) Reset() {
	*x = ParagraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParagraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParagraphRequest) ProtoMessage() {}

func (x *ParagraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParagraphRequest.ProtoReflect.Descriptor instead.
func (*ParagraphRequest) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{4}
}

func (x *ParagraphRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type ParagraphResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ParagraphResponse) Reset() {
	*x = ParagraphResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParagraphResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParagraphResponse) ProtoMessage() {}

func (x *ParagraphResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParagraphResponse.ProtoReflect.Descriptor instead.
func (*ParagraphResponse) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{5}
}

func (x *ParagraphResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_messagepb_message_proto protoreflect.FileDescriptor

var file_messagepb_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x22, 0x30, 0x0a, 0x16,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28,
	0x0a, 0x0c, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a,
	0x10, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xf6, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x15, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x74, 0x6f, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x13, 0x4d,
	0x61, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x42, 0x79, 0x57, 0x6f,
	0x72, 0x64, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_messagepb_message_proto_rawDescOnce sync.Once
	file_messagepb_message_proto_rawDescData = file_messagepb_message_proto_rawDesc
)

func file_messagepb_message_proto_rawDescGZIP() []byte {
	file_messagepb_message_proto_rawDescOnce.Do(func() {
		file_messagepb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagepb_message_proto_rawDescData)
	})
	return file_messagepb_message_proto_rawDescData
}

var file_messagepb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_messagepb_message_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),         // 0: message.MessageRequest
	(*CombineMessageResponse)(nil), // 1: message.CombineMessageResponse
	(*SplitMessage)(nil),           // 2: message.SplitMessage
	(*ResponseWords)(nil),          // 3: message.ResponseWords
	(*ParagraphRequest)(nil),       // 4: message.ParagraphRequest
	(*ParagraphResponse)(nil),      // 5: message.ParagraphResponse
}
var file_messagepb_message_proto_depIdxs = []int32{
	0, // 0: message.MessageService.CombineMessage:input_type -> message.MessageRequest
	2, // 1: message.MessageService.SplitMessageIntoWords:input_type -> message.SplitMessage
	4, // 2: message.MessageService.MakeParagraphByWord:input_type -> message.ParagraphRequest
	1, // 3: message.MessageService.CombineMessage:output_type -> message.CombineMessageResponse
	3, // 4: message.MessageService.SplitMessageIntoWords:output_type -> message.ResponseWords
	5, // 5: message.MessageService.MakeParagraphByWord:output_type -> message.ParagraphResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messagepb_message_proto_init() }
func file_messagepb_message_proto_init() {
	if File_messagepb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagepb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_messagepb_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombineMessageResponse); i {
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
		file_messagepb_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitMessage); i {
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
		file_messagepb_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseWords); i {
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
		file_messagepb_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParagraphRequest); i {
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
		file_messagepb_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParagraphResponse); i {
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
			RawDescriptor: file_messagepb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messagepb_message_proto_goTypes,
		DependencyIndexes: file_messagepb_message_proto_depIdxs,
		MessageInfos:      file_messagepb_message_proto_msgTypes,
	}.Build()
	File_messagepb_message_proto = out.File
	file_messagepb_message_proto_rawDesc = nil
	file_messagepb_message_proto_goTypes = nil
	file_messagepb_message_proto_depIdxs = nil
}
