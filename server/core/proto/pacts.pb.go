// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pacts.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	pacts.proto

It has these top-level messages:
	Participant
	ParticipantVersion
	Pact
	PactPublication
	Tag
	PublishPactRequest
	Response
	PublishPactResponse
	GetAllProviderPactsRequest
	ConsumerInfo
	Links
	GetAllProviderPactsReponse
	GetProviderConsumerVersionPactRequest
	GetProviderConsumerVersionPactResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

/*
type Response_Code int32

const (
	Response_UNKNOWN Response_Code = 0
	Response_SUCCESS Response_Code = 1
	Response_FAIL    Response_Code = 2
)

var Response_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAIL",
}
var Response_Code_value = map[string]int32{
	"UNKNOWN": 0,
	"SUCCESS": 1,
	"FAIL":    2,
}

func (x Response_Code) String() string {
	return proto1.EnumName(Response_Code_name, int32(x))
}
func (Response_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }
*/

type Participant struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AppId       string `protobuf:"bytes,2,opt,name=appId" json:"appId,omitempty"`
	ServiceName string `protobuf:"bytes,3,opt,name=serviceName" json:"serviceName,omitempty"`
}

func (m *Participant) Reset()                    { *m = Participant{} }
func (m *Participant) String() string            { return proto1.CompactTextString(m) }
func (*Participant) ProtoMessage()               {}
func (*Participant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Participant) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Participant) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Participant) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type ParticipantVersion struct {
	Id            int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Number        string `protobuf:"bytes,2,opt,name=number" json:"number,omitempty"`
	ParticipantId int32  `protobuf:"varint,3,opt,name=participantId" json:"participantId,omitempty"`
	Order         int32  `protobuf:"varint,4,opt,name=order" json:"order,omitempty"`
}

func (m *ParticipantVersion) Reset()                    { *m = ParticipantVersion{} }
func (m *ParticipantVersion) String() string            { return proto1.CompactTextString(m) }
func (*ParticipantVersion) ProtoMessage()               {}
func (*ParticipantVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ParticipantVersion) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ParticipantVersion) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *ParticipantVersion) GetParticipantId() int32 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

func (m *ParticipantVersion) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

type Pact struct {
	Id                    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ConsumerParticipantId int32  `protobuf:"varint,2,opt,name=consumerParticipantId" json:"consumerParticipantId,omitempty"`
	ProviderParticipantId int32  `protobuf:"varint,3,opt,name=providerParticipantId" json:"providerParticipantId,omitempty"`
	Sha                   []byte `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty"`
	Content               []byte `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Pact) Reset()                    { *m = Pact{} }
func (m *Pact) String() string            { return proto1.CompactTextString(m) }
func (*Pact) ProtoMessage()               {}
func (*Pact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Pact) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pact) GetConsumerParticipantId() int32 {
	if m != nil {
		return m.ConsumerParticipantId
	}
	return 0
}

func (m *Pact) GetProviderParticipantId() int32 {
	if m != nil {
		return m.ProviderParticipantId
	}
	return 0
}

func (m *Pact) GetSha() []byte {
	if m != nil {
		return m.Sha
	}
	return nil
}

func (m *Pact) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PactPublication struct {
	Id                   int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ParticipantVersionId int32 `protobuf:"varint,2,opt,name=participantVersionId" json:"participantVersionId,omitempty"`
	PactId               int32 `protobuf:"varint,3,opt,name=pactId" json:"pactId,omitempty"`
	ParticipantId        int32 `protobuf:"varint,4,opt,name=participantId" json:"participantId,omitempty"`
}

func (m *PactPublication) Reset()                    { *m = PactPublication{} }
func (m *PactPublication) String() string            { return proto1.CompactTextString(m) }
func (*PactPublication) ProtoMessage()               {}
func (*PactPublication) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PactPublication) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PactPublication) GetParticipantVersionId() int32 {
	if m != nil {
		return m.ParticipantVersionId
	}
	return 0
}

func (m *PactPublication) GetPactId() int32 {
	if m != nil {
		return m.PactId
	}
	return 0
}

func (m *PactPublication) GetParticipantId() int32 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

type Tag struct {
	Name                 string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ParticipantVersionId int32  `protobuf:"varint,2,opt,name=participantVersionId" json:"participantVersionId,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto1.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetParticipantVersionId() int32 {
	if m != nil {
		return m.ParticipantVersionId
	}
	return 0
}

type PublishPactRequest struct {
	ProviderId string `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	ConsumerId string `protobuf:"bytes,2,opt,name=consumerId" json:"consumerId,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Pact       []byte `protobuf:"bytes,4,opt,name=pact,proto3" json:"pact,omitempty"`
}

func (m *PublishPactRequest) Reset()                    { *m = PublishPactRequest{} }
func (m *PublishPactRequest) String() string            { return proto1.CompactTextString(m) }
func (*PublishPactRequest) ProtoMessage()               {}
func (*PublishPactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublishPactRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *PublishPactRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *PublishPactRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PublishPactRequest) GetPact() []byte {
	if m != nil {
		return m.Pact
	}
	return nil
}

/*
type Response struct {
	Code    Response_Code `protobuf:"varint,1,opt,name=code,enum=com.huawei.paas.cse.serviceregistry.api.Response_Code" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetCode() Response_Code {
	if m != nil {
		return m.Code
	}
	return Response_UNKNOWN
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}
*/
type PublishPactResponse struct {
	Response *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *PublishPactResponse) Reset()                    { *m = PublishPactResponse{} }
func (m *PublishPactResponse) String() string            { return proto1.CompactTextString(m) }
func (*PublishPactResponse) ProtoMessage()               {}
func (*PublishPactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PublishPactResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetAllProviderPactsRequest struct {
	ProviderId string `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
}

func (m *GetAllProviderPactsRequest) Reset()                    { *m = GetAllProviderPactsRequest{} }
func (m *GetAllProviderPactsRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetAllProviderPactsRequest) ProtoMessage()               {}
func (*GetAllProviderPactsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetAllProviderPactsRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

type ConsumerInfo struct {
	Href string `protobuf:"bytes,1,opt,name=href" json:"href,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ConsumerInfo) Reset()                    { *m = ConsumerInfo{} }
func (m *ConsumerInfo) String() string            { return proto1.CompactTextString(m) }
func (*ConsumerInfo) ProtoMessage()               {}
func (*ConsumerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ConsumerInfo) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *ConsumerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Links struct {
	Pacts []*ConsumerInfo `protobuf:"bytes,1,rep,name=pacts" json:"pacts,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto1.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Links) GetPacts() []*ConsumerInfo {
	if m != nil {
		return m.Pacts
	}
	return nil
}

type GetAllProviderPactsReponse struct {
	Response *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XLinks   *Links    `protobuf:"bytes,2,opt,name=_links,json=Links" json:"_links,omitempty"`
}

func (m *GetAllProviderPactsReponse) Reset()                    { *m = GetAllProviderPactsReponse{} }
func (m *GetAllProviderPactsReponse) String() string            { return proto1.CompactTextString(m) }
func (*GetAllProviderPactsReponse) ProtoMessage()               {}
func (*GetAllProviderPactsReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAllProviderPactsReponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetAllProviderPactsReponse) GetXLinks() *Links {
	if m != nil {
		return m.XLinks
	}
	return nil
}

type GetProviderConsumerVersionPactRequest struct {
	ProviderId string `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	ConsumerId string `protobuf:"bytes,2,opt,name=consumerId" json:"consumerId,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *GetProviderConsumerVersionPactRequest) Reset()         { *m = GetProviderConsumerVersionPactRequest{} }
func (m *GetProviderConsumerVersionPactRequest) String() string { return proto1.CompactTextString(m) }
func (*GetProviderConsumerVersionPactRequest) ProtoMessage()    {}
func (*GetProviderConsumerVersionPactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12}
}

func (m *GetProviderConsumerVersionPactRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *GetProviderConsumerVersionPactRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *GetProviderConsumerVersionPactRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetProviderConsumerVersionPactResponse struct {
	Response *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Pact     []byte    `protobuf:"bytes,2,opt,name=pact,proto3" json:"pact,omitempty"`
}

func (m *GetProviderConsumerVersionPactResponse) Reset() {
	*m = GetProviderConsumerVersionPactResponse{}
}
func (m *GetProviderConsumerVersionPactResponse) String() string { return proto1.CompactTextString(m) }
func (*GetProviderConsumerVersionPactResponse) ProtoMessage()    {}
func (*GetProviderConsumerVersionPactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13}
}

func (m *GetProviderConsumerVersionPactResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetProviderConsumerVersionPactResponse) GetPact() []byte {
	if m != nil {
		return m.Pact
	}
	return nil
}

func init() {
	proto1.RegisterType((*Participant)(nil), "com.huawei.paas.cse.serviceregistry.api.Participant")
	proto1.RegisterType((*ParticipantVersion)(nil), "com.huawei.paas.cse.serviceregistry.api.ParticipantVersion")
	proto1.RegisterType((*Pact)(nil), "com.huawei.paas.cse.serviceregistry.api.Pact")
	proto1.RegisterType((*PactPublication)(nil), "com.huawei.paas.cse.serviceregistry.api.PactPublication")
	proto1.RegisterType((*Tag)(nil), "com.huawei.paas.cse.serviceregistry.api.Tag")
	proto1.RegisterType((*PublishPactRequest)(nil), "com.huawei.paas.cse.serviceregistry.api.PublishPactRequest")
	proto1.RegisterType((*Response)(nil), "com.huawei.paas.cse.serviceregistry.api.Response")
	proto1.RegisterType((*PublishPactResponse)(nil), "com.huawei.paas.cse.serviceregistry.api.PublishPactResponse")
	proto1.RegisterType((*GetAllProviderPactsRequest)(nil), "com.huawei.paas.cse.serviceregistry.api.GetAllProviderPactsRequest")
	proto1.RegisterType((*ConsumerInfo)(nil), "com.huawei.paas.cse.serviceregistry.api.ConsumerInfo")
	proto1.RegisterType((*Links)(nil), "com.huawei.paas.cse.serviceregistry.api.Links")
	proto1.RegisterType((*GetAllProviderPactsReponse)(nil), "com.huawei.paas.cse.serviceregistry.api.GetAllProviderPactsReponse")
	proto1.RegisterType((*GetProviderConsumerVersionPactRequest)(nil), "com.huawei.paas.cse.serviceregistry.api.GetProviderConsumerVersionPactRequest")
	proto1.RegisterType((*GetProviderConsumerVersionPactResponse)(nil), "com.huawei.paas.cse.serviceregistry.api.GetProviderConsumerVersionPactResponse")
	proto1.RegisterEnum("com.huawei.paas.cse.serviceregistry.api.Response_Code", Response_Code_name, Response_Code_value)
}

func init() { proto1.RegisterFile("pacts.proto", fileDescriptor0) }

/*
var fileDescriptor0 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xad, 0x63, 0xb7, 0xfd, 0x8d, 0x4b, 0x89, 0x96, 0x82, 0x2c, 0x0e, 0x28, 0x5a, 0xf1,
	0xa7, 0xe2, 0x60, 0x41, 0x80, 0x9e, 0xb8, 0x94, 0xa8, 0x54, 0xa1, 0x6d, 0x88, 0xdc, 0x16, 0x24,
	0x2e, 0x68, 0x63, 0x6f, 0x93, 0x15, 0x89, 0xd7, 0xec, 0x6e, 0x02, 0x1c, 0xe1, 0xca, 0x91, 0x2b,
	0x9f, 0x81, 0x03, 0x9f, 0x10, 0xed, 0xfa, 0x4f, 0x12, 0x12, 0x44, 0x40, 0x2a, 0xa7, 0xec, 0xcc,
	0xd8, 0x6f, 0xdf, 0x7b, 0x33, 0x13, 0x83, 0x9f, 0xd1, 0x58, 0xab, 0x30, 0x93, 0x42, 0x0b, 0x7c,
	0x27, 0x16, 0xa3, 0x70, 0x30, 0xa6, 0xef, 0x18, 0x0f, 0x33, 0x4a, 0x55, 0x18, 0x2b, 0x16, 0x2a,
	0x26, 0x27, 0x3c, 0x66, 0x92, 0xf5, 0xb9, 0xd2, 0xf2, 0x43, 0x48, 0x33, 0x4e, 0xce, 0xc0, 0xef,
	0x52, 0xa9, 0x79, 0xcc, 0x33, 0x9a, 0x6a, 0xbc, 0x05, 0x0e, 0x4f, 0x02, 0xd4, 0x40, 0x3b, 0x5e,
	0xe4, 0xf0, 0x04, 0x6f, 0x83, 0x47, 0xb3, 0xac, 0x9d, 0x04, 0x4e, 0x03, 0xed, 0xfc, 0x1f, 0xe5,
	0x01, 0x6e, 0x80, 0x5f, 0x60, 0x75, 0xe8, 0x88, 0x05, 0x35, 0x5b, 0x9b, 0x4d, 0x91, 0xf7, 0x80,
	0x67, 0x60, 0x5f, 0x30, 0xa9, 0xb8, 0x48, 0x17, 0xd0, 0xaf, 0xc1, 0x5a, 0x3a, 0x1e, 0xf5, 0x98,
	0x2c, 0xe0, 0x8b, 0x08, 0xdf, 0x84, 0x4b, 0xd9, 0xf4, 0xed, 0x76, 0x62, 0x6f, 0xf0, 0xa2, 0xf9,
	0xa4, 0xe1, 0x26, 0x64, 0xc2, 0x64, 0xe0, 0xda, 0x6a, 0x1e, 0x90, 0x6f, 0x08, 0xdc, 0x2e, 0x8d,
	0x17, 0xa5, 0x3c, 0x84, 0xab, 0xb1, 0x48, 0xd5, 0x78, 0xc4, 0x64, 0x77, 0x0e, 0xdc, 0xb1, 0x8f,
	0x2c, 0x2f, 0x9a, 0xb7, 0x32, 0x29, 0x26, 0x3c, 0xf9, 0xf9, 0xad, 0x9c, 0xd2, 0xf2, 0x22, 0xae,
	0x43, 0x4d, 0x0d, 0xa8, 0x25, 0xb6, 0x19, 0x99, 0x23, 0x0e, 0x60, 0x3d, 0x16, 0xa9, 0x66, 0xa9,
	0x0e, 0x3c, 0x9b, 0x2d, 0x43, 0xf2, 0x05, 0xc1, 0x65, 0x43, 0xb8, 0x3b, 0xee, 0x0d, 0x79, 0x4c,
	0xf5, 0x32, 0xa3, 0x9a, 0xb0, 0x9d, 0x2d, 0xd8, 0x59, 0x51, 0x5f, 0x5a, 0x33, 0xe6, 0x9a, 0x89,
	0xa8, 0xa8, 0x16, 0xd1, 0xa2, 0xb9, 0xee, 0x12, 0x73, 0xc9, 0x31, 0xd4, 0x4e, 0x69, 0x1f, 0x63,
	0x70, 0x53, 0xd3, 0x62, 0x64, 0xfb, 0x63, 0xcf, 0x7f, 0x43, 0x86, 0x7c, 0x42, 0x80, 0xad, 0x40,
	0x35, 0x30, 0x5a, 0x23, 0xf6, 0x76, 0xcc, 0x94, 0xc6, 0x37, 0x00, 0x4a, 0x03, 0xdb, 0x49, 0x71,
	0xc9, 0x4c, 0xc6, 0xd4, 0xcb, 0xb6, 0x54, 0x33, 0x38, 0x93, 0x31, 0xae, 0x4e, 0xf2, 0x3b, 0x8a,
	0x21, 0x2c, 0x43, 0x43, 0xdc, 0xe8, 0x2d, 0x5a, 0x60, 0xcf, 0xe4, 0x2b, 0x82, 0x8d, 0x88, 0xa9,
	0x4c, 0xa4, 0x8a, 0xe1, 0x67, 0xe0, 0xc6, 0x22, 0xc9, 0x95, 0x6d, 0x35, 0x77, 0xc3, 0x15, 0x17,
	0x26, 0x2c, 0x01, 0xc2, 0x96, 0x48, 0x58, 0x64, 0x31, 0x0c, 0x8d, 0x11, 0x53, 0x8a, 0xf6, 0x59,
	0xc1, 0xb1, 0x0c, 0xc9, 0x5d, 0x70, 0xcd, 0x73, 0xd8, 0x87, 0xf5, 0xb3, 0xce, 0x61, 0xe7, 0xf9,
	0xcb, 0x4e, 0xfd, 0x3f, 0x13, 0x9c, 0x9c, 0xb5, 0x5a, 0xfb, 0x27, 0x27, 0x75, 0x84, 0x37, 0xc0,
	0x7d, 0xba, 0xd7, 0x3e, 0xaa, 0x3b, 0x24, 0x81, 0x2b, 0x73, 0x16, 0x15, 0x44, 0x8f, 0x61, 0x43,
	0x16, 0x67, 0x4b, 0xd6, 0x6f, 0xde, 0xff, 0x63, 0xb2, 0x51, 0x05, 0x41, 0x1e, 0xc3, 0xf5, 0x03,
	0xa6, 0xf7, 0x86, 0xc3, 0x6e, 0x35, 0xb9, 0xb1, 0x56, 0x2b, 0x36, 0x84, 0xec, 0xc2, 0x66, 0xab,
	0xb4, 0x3f, 0x3d, 0x17, 0xc6, 0xe6, 0x81, 0x64, 0xe7, 0xe5, 0x7c, 0x98, 0x73, 0x35, 0x33, 0xce,
	0x74, 0x66, 0xc8, 0x29, 0x78, 0x47, 0x3c, 0x7d, 0xa3, 0xf0, 0x21, 0x78, 0xf6, 0x7f, 0x2a, 0x40,
	0x8d, 0xda, 0x8e, 0xdf, 0x7c, 0xb4, 0xb2, 0x94, 0xd9, 0x6b, 0xa3, 0x1c, 0x83, 0x7c, 0x47, 0xbf,
	0x10, 0x73, 0x11, 0xce, 0xe1, 0x7d, 0x58, 0x7b, 0x3d, 0x34, 0x22, 0xac, 0x32, 0xbf, 0x19, 0xae,
	0x0c, 0x66, 0xa5, 0x47, 0xb9, 0x03, 0xe4, 0x23, 0x82, 0x5b, 0x07, 0x4c, 0x97, 0x8c, 0x4b, 0x5d,
	0xc5, 0xae, 0xfc, 0x93, 0xed, 0x20, 0x9f, 0x11, 0xdc, 0xfe, 0x1d, 0x87, 0x0b, 0x19, 0xbf, 0x6a,
	0x2f, 0x9d, 0xe9, 0x5e, 0x3e, 0xb9, 0x07, 0xab, 0x7e, 0xae, 0x5e, 0x79, 0xf6, 0xf3, 0xd6, 0x5b,
	0xb3, 0x3f, 0x0f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x54, 0x1e, 0xa1, 0x98, 0xf4, 0x06, 0x00,
	0x00,
}
*/
