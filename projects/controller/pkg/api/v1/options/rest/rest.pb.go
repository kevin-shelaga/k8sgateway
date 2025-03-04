// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/rest/rest.proto

package rest

import (
	reflect "reflect"
	sync "sync"

	transformation "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/transformation"
	transformation1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/transformation"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transformations map[string]*transformation.TransformationTemplate `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SwaggerInfo     *ServiceSpec_SwaggerInfo                          `protobuf:"bytes,2,opt,name=swagger_info,json=swaggerInfo,proto3" json:"swagger_info,omitempty"`
}

func (x *ServiceSpec) Reset() {
	*x = ServiceSpec{}
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec) ProtoMessage() {}

func (x *ServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec.ProtoReflect.Descriptor instead.
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceSpec) GetTransformations() map[string]*transformation.TransformationTemplate {
	if x != nil {
		return x.Transformations
	}
	return nil
}

func (x *ServiceSpec) GetSwaggerInfo() *ServiceSpec_SwaggerInfo {
	if x != nil {
		return x.SwaggerInfo
	}
	return nil
}

// This is only for upstream with REST service spec
type DestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionName           string                                 `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Parameters             *transformation1.Parameters            `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	ResponseTransformation *transformation.TransformationTemplate `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
}

func (x *DestinationSpec) Reset() {
	*x = DestinationSpec{}
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationSpec) ProtoMessage() {}

func (x *DestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationSpec.ProtoReflect.Descriptor instead.
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescGZIP(), []int{1}
}

func (x *DestinationSpec) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *DestinationSpec) GetParameters() *transformation1.Parameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *DestinationSpec) GetResponseTransformation() *transformation.TransformationTemplate {
	if x != nil {
		return x.ResponseTransformation
	}
	return nil
}

type ServiceSpec_SwaggerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SwaggerSpec:
	//
	//	*ServiceSpec_SwaggerInfo_Url
	//	*ServiceSpec_SwaggerInfo_Inline
	SwaggerSpec isServiceSpec_SwaggerInfo_SwaggerSpec `protobuf_oneof:"swagger_spec"`
}

func (x *ServiceSpec_SwaggerInfo) Reset() {
	*x = ServiceSpec_SwaggerInfo{}
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceSpec_SwaggerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec_SwaggerInfo) ProtoMessage() {}

func (x *ServiceSpec_SwaggerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec_SwaggerInfo.ProtoReflect.Descriptor instead.
func (*ServiceSpec_SwaggerInfo) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ServiceSpec_SwaggerInfo) GetSwaggerSpec() isServiceSpec_SwaggerInfo_SwaggerSpec {
	if m != nil {
		return m.SwaggerSpec
	}
	return nil
}

func (x *ServiceSpec_SwaggerInfo) GetUrl() string {
	if x, ok := x.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Url); ok {
		return x.Url
	}
	return ""
}

func (x *ServiceSpec_SwaggerInfo) GetInline() string {
	if x, ok := x.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Inline); ok {
		return x.Inline
	}
	return ""
}

type isServiceSpec_SwaggerInfo_SwaggerSpec interface {
	isServiceSpec_SwaggerInfo_SwaggerSpec()
}

type ServiceSpec_SwaggerInfo_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type ServiceSpec_SwaggerInfo_Inline struct {
	Inline string `protobuf:"bytes,2,opt,name=inline,proto3,oneof"`
}

func (*ServiceSpec_SwaggerInfo_Url) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

func (*ServiceSpec_SwaggerInfo_Inline) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

var File_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x65,
	0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x65, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a,
	0x0c, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x74, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x46,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0b, 0x53, 0x77,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4f, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x69, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x51, 0xb8, 0xf5,
	0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescData = file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_goTypes = []any{
	(*ServiceSpec)(nil),                // 0: rest.options.gloo.solo.io.ServiceSpec
	(*DestinationSpec)(nil),            // 1: rest.options.gloo.solo.io.DestinationSpec
	nil,                                // 2: rest.options.gloo.solo.io.ServiceSpec.TransformationsEntry
	(*ServiceSpec_SwaggerInfo)(nil),    // 3: rest.options.gloo.solo.io.ServiceSpec.SwaggerInfo
	(*transformation1.Parameters)(nil), // 4: transformation.options.gloo.solo.io.Parameters
	(*transformation.TransformationTemplate)(nil), // 5: envoy.api.v2.filter.http.TransformationTemplate
}
var file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_depIdxs = []int32{
	2, // 0: rest.options.gloo.solo.io.ServiceSpec.transformations:type_name -> rest.options.gloo.solo.io.ServiceSpec.TransformationsEntry
	3, // 1: rest.options.gloo.solo.io.ServiceSpec.swagger_info:type_name -> rest.options.gloo.solo.io.ServiceSpec.SwaggerInfo
	4, // 2: rest.options.gloo.solo.io.DestinationSpec.parameters:type_name -> transformation.options.gloo.solo.io.Parameters
	5, // 3: rest.options.gloo.solo.io.DestinationSpec.response_transformation:type_name -> envoy.api.v2.filter.http.TransformationTemplate
	5, // 4: rest.options.gloo.solo.io.ServiceSpec.TransformationsEntry.value:type_name -> envoy.api.v2.filter.http.TransformationTemplate
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_init() }
func file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_init() {
	if File_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes[3].OneofWrappers = []any{
		(*ServiceSpec_SwaggerInfo_Url)(nil),
		(*ServiceSpec_SwaggerInfo_Inline)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto = out.File
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_controller_api_v1_options_rest_rest_proto_depIdxs = nil
}
