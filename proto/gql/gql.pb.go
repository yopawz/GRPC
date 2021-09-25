// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gql/gql.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetAllEmployeesWithGQLParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllEmployeesWithGQLParams) Reset()         { *m = GetAllEmployeesWithGQLParams{} }
func (m *GetAllEmployeesWithGQLParams) String() string { return proto.CompactTextString(m) }
func (*GetAllEmployeesWithGQLParams) ProtoMessage()    {}
func (*GetAllEmployeesWithGQLParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{0}
}

func (m *GetAllEmployeesWithGQLParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllEmployeesWithGQLParams.Unmarshal(m, b)
}
func (m *GetAllEmployeesWithGQLParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllEmployeesWithGQLParams.Marshal(b, m, deterministic)
}
func (m *GetAllEmployeesWithGQLParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllEmployeesWithGQLParams.Merge(m, src)
}
func (m *GetAllEmployeesWithGQLParams) XXX_Size() int {
	return xxx_messageInfo_GetAllEmployeesWithGQLParams.Size(m)
}
func (m *GetAllEmployeesWithGQLParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllEmployeesWithGQLParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllEmployeesWithGQLParams proto.InternalMessageInfo

type GetEmployeeWithIdWithGQLParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEmployeeWithIdWithGQLParams) Reset()         { *m = GetEmployeeWithIdWithGQLParams{} }
func (m *GetEmployeeWithIdWithGQLParams) String() string { return proto.CompactTextString(m) }
func (*GetEmployeeWithIdWithGQLParams) ProtoMessage()    {}
func (*GetEmployeeWithIdWithGQLParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{1}
}

func (m *GetEmployeeWithIdWithGQLParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEmployeeWithIdWithGQLParams.Unmarshal(m, b)
}
func (m *GetEmployeeWithIdWithGQLParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEmployeeWithIdWithGQLParams.Marshal(b, m, deterministic)
}
func (m *GetEmployeeWithIdWithGQLParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEmployeeWithIdWithGQLParams.Merge(m, src)
}
func (m *GetEmployeeWithIdWithGQLParams) XXX_Size() int {
	return xxx_messageInfo_GetEmployeeWithIdWithGQLParams.Size(m)
}
func (m *GetEmployeeWithIdWithGQLParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEmployeeWithIdWithGQLParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetEmployeeWithIdWithGQLParams proto.InternalMessageInfo

func (m *GetEmployeeWithIdWithGQLParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EmployeeGQL struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Salary               int32    `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeGQL) Reset()         { *m = EmployeeGQL{} }
func (m *EmployeeGQL) String() string { return proto.CompactTextString(m) }
func (*EmployeeGQL) ProtoMessage()    {}
func (*EmployeeGQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{2}
}

func (m *EmployeeGQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeGQL.Unmarshal(m, b)
}
func (m *EmployeeGQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeGQL.Marshal(b, m, deterministic)
}
func (m *EmployeeGQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeGQL.Merge(m, src)
}
func (m *EmployeeGQL) XXX_Size() int {
	return xxx_messageInfo_EmployeeGQL.Size(m)
}
func (m *EmployeeGQL) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeGQL.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeGQL proto.InternalMessageInfo

func (m *EmployeeGQL) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EmployeeGQL) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EmployeeGQL) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *EmployeeGQL) GetSalary() int32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type ResponseGQL struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseGQL) Reset()         { *m = ResponseGQL{} }
func (m *ResponseGQL) String() string { return proto.CompactTextString(m) }
func (*ResponseGQL) ProtoMessage()    {}
func (*ResponseGQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{3}
}

func (m *ResponseGQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGQL.Unmarshal(m, b)
}
func (m *ResponseGQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGQL.Marshal(b, m, deterministic)
}
func (m *ResponseGQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGQL.Merge(m, src)
}
func (m *ResponseGQL) XXX_Size() int {
	return xxx_messageInfo_ResponseGQL.Size(m)
}
func (m *ResponseGQL) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGQL.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGQL proto.InternalMessageInfo

func (m *ResponseGQL) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EmployeeListGQL struct {
	Employees            []*EmployeeGQL `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EmployeeListGQL) Reset()         { *m = EmployeeListGQL{} }
func (m *EmployeeListGQL) String() string { return proto.CompactTextString(m) }
func (*EmployeeListGQL) ProtoMessage()    {}
func (*EmployeeListGQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{4}
}

func (m *EmployeeListGQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeListGQL.Unmarshal(m, b)
}
func (m *EmployeeListGQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeListGQL.Marshal(b, m, deterministic)
}
func (m *EmployeeListGQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeListGQL.Merge(m, src)
}
func (m *EmployeeListGQL) XXX_Size() int {
	return xxx_messageInfo_EmployeeListGQL.Size(m)
}
func (m *EmployeeListGQL) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeListGQL.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeListGQL proto.InternalMessageInfo

func (m *EmployeeListGQL) GetEmployees() []*EmployeeGQL {
	if m != nil {
		return m.Employees
	}
	return nil
}

type InsertEmployeeWithGQLParams struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Salary               int32    `protobuf:"varint,3,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertEmployeeWithGQLParams) Reset()         { *m = InsertEmployeeWithGQLParams{} }
func (m *InsertEmployeeWithGQLParams) String() string { return proto.CompactTextString(m) }
func (*InsertEmployeeWithGQLParams) ProtoMessage()    {}
func (*InsertEmployeeWithGQLParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{5}
}

func (m *InsertEmployeeWithGQLParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertEmployeeWithGQLParams.Unmarshal(m, b)
}
func (m *InsertEmployeeWithGQLParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertEmployeeWithGQLParams.Marshal(b, m, deterministic)
}
func (m *InsertEmployeeWithGQLParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertEmployeeWithGQLParams.Merge(m, src)
}
func (m *InsertEmployeeWithGQLParams) XXX_Size() int {
	return xxx_messageInfo_InsertEmployeeWithGQLParams.Size(m)
}
func (m *InsertEmployeeWithGQLParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertEmployeeWithGQLParams.DiscardUnknown(m)
}

var xxx_messageInfo_InsertEmployeeWithGQLParams proto.InternalMessageInfo

func (m *InsertEmployeeWithGQLParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertEmployeeWithGQLParams) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *InsertEmployeeWithGQLParams) GetSalary() int32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type UpdateEmployeeWithGQLParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Salary               int32    `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEmployeeWithGQLParams) Reset()         { *m = UpdateEmployeeWithGQLParams{} }
func (m *UpdateEmployeeWithGQLParams) String() string { return proto.CompactTextString(m) }
func (*UpdateEmployeeWithGQLParams) ProtoMessage()    {}
func (*UpdateEmployeeWithGQLParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{6}
}

func (m *UpdateEmployeeWithGQLParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEmployeeWithGQLParams.Unmarshal(m, b)
}
func (m *UpdateEmployeeWithGQLParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEmployeeWithGQLParams.Marshal(b, m, deterministic)
}
func (m *UpdateEmployeeWithGQLParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEmployeeWithGQLParams.Merge(m, src)
}
func (m *UpdateEmployeeWithGQLParams) XXX_Size() int {
	return xxx_messageInfo_UpdateEmployeeWithGQLParams.Size(m)
}
func (m *UpdateEmployeeWithGQLParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEmployeeWithGQLParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEmployeeWithGQLParams proto.InternalMessageInfo

func (m *UpdateEmployeeWithGQLParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateEmployeeWithGQLParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateEmployeeWithGQLParams) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UpdateEmployeeWithGQLParams) GetSalary() int32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type DeleteEmployeeWithGQLParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEmployeeWithGQLParams) Reset()         { *m = DeleteEmployeeWithGQLParams{} }
func (m *DeleteEmployeeWithGQLParams) String() string { return proto.CompactTextString(m) }
func (*DeleteEmployeeWithGQLParams) ProtoMessage()    {}
func (*DeleteEmployeeWithGQLParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_996dd0c3ac707142, []int{7}
}

func (m *DeleteEmployeeWithGQLParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEmployeeWithGQLParams.Unmarshal(m, b)
}
func (m *DeleteEmployeeWithGQLParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEmployeeWithGQLParams.Marshal(b, m, deterministic)
}
func (m *DeleteEmployeeWithGQLParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEmployeeWithGQLParams.Merge(m, src)
}
func (m *DeleteEmployeeWithGQLParams) XXX_Size() int {
	return xxx_messageInfo_DeleteEmployeeWithGQLParams.Size(m)
}
func (m *DeleteEmployeeWithGQLParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEmployeeWithGQLParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEmployeeWithGQLParams proto.InternalMessageInfo

func (m *DeleteEmployeeWithGQLParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetAllEmployeesWithGQLParams)(nil), "proto.GetAllEmployeesWithGQLParams")
	proto.RegisterType((*GetEmployeeWithIdWithGQLParams)(nil), "proto.GetEmployeeWithIdWithGQLParams")
	proto.RegisterType((*EmployeeGQL)(nil), "proto.EmployeeGQL")
	proto.RegisterType((*ResponseGQL)(nil), "proto.ResponseGQL")
	proto.RegisterType((*EmployeeListGQL)(nil), "proto.EmployeeListGQL")
	proto.RegisterType((*InsertEmployeeWithGQLParams)(nil), "proto.InsertEmployeeWithGQLParams")
	proto.RegisterType((*UpdateEmployeeWithGQLParams)(nil), "proto.UpdateEmployeeWithGQLParams")
	proto.RegisterType((*DeleteEmployeeWithGQLParams)(nil), "proto.DeleteEmployeeWithGQLParams")
}

func init() { proto.RegisterFile("proto/gql/gql.proto", fileDescriptor_996dd0c3ac707142) }

var fileDescriptor_996dd0c3ac707142 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x49, 0xbb, 0xba, 0xb7, 0x52, 0xbb, 0xb7, 0xa4, 0x8d, 0xa9, 0x2e, 0x65, 0x44, 0xec,
	0xd3, 0x5a, 0x57, 0x41, 0xf1, 0xcd, 0x8d, 0x52, 0x16, 0x2a, 0xb8, 0x05, 0xf1, 0xc1, 0xa7, 0xd1,
	0x0c, 0x69, 0x60, 0xf2, 0x61, 0x26, 0x3e, 0xec, 0xa3, 0x3f, 0xcb, 0xfc, 0x3a, 0x99, 0x49, 0xa6,
	0x49, 0x4a, 0x0c, 0xec, 0x43, 0x69, 0x26, 0xf7, 0xe3, 0xdc, 0x73, 0xce, 0x9d, 0xc0, 0x34, 0xcd,
	0x92, 0x3c, 0x79, 0x19, 0xfc, 0xe2, 0xf2, 0x77, 0xa1, 0x4e, 0x38, 0x54, 0x7f, 0xee, 0xbc, 0x11,
	0xcb, 0x68, 0xba, 0xd7, 0x71, 0x72, 0x0e, 0x4f, 0x36, 0x2c, 0xff, 0xc0, 0xf9, 0xa7, 0x28, 0xe5,
	0xc9, 0x2d, 0x63, 0xe2, 0x5b, 0x98, 0xef, 0x37, 0x37, 0xdb, 0x2f, 0x34, 0xa3, 0x91, 0x20, 0x6f,
	0xe1, 0x7c, 0xc3, 0x72, 0x1d, 0x94, 0xb1, 0x6b, 0xbf, 0x95, 0x81, 0x36, 0x98, 0xa1, 0xef, 0x18,
	0x4b, 0x63, 0x35, 0xbc, 0x1a, 0x16, 0x9e, 0xf9, 0xc0, 0xd8, 0x99, 0xa1, 0x4f, 0xbe, 0xc3, 0x48,
	0x57, 0x6d, 0x6e, 0xb6, 0x38, 0xae, 0xb3, 0x64, 0x18, 0x11, 0x06, 0x31, 0x8d, 0x98, 0x63, 0x2e,
	0x8d, 0xd5, 0xe9, 0x4e, 0x3d, 0xe3, 0x04, 0x2c, 0x1a, 0x30, 0xc7, 0x52, 0x49, 0xf2, 0x11, 0x67,
	0x70, 0x22, 0x28, 0xa7, 0xd9, 0xad, 0x33, 0x50, 0x2f, 0xab, 0x13, 0x79, 0x01, 0xa3, 0x1d, 0x13,
	0x69, 0x12, 0x0b, 0xd5, 0xdc, 0x81, 0xfb, 0x11, 0x13, 0x42, 0x16, 0x1b, 0xaa, 0x9f, 0x3e, 0x12,
	0x0f, 0x1e, 0xe9, 0x29, 0xb6, 0xa1, 0xc8, 0x65, 0xf2, 0x1a, 0x4e, 0x99, 0xe6, 0xea, 0x18, 0x4b,
	0x6b, 0x35, 0xba, 0xc4, 0x52, 0x8c, 0x8b, 0xc6, 0xc0, 0xbb, 0x3a, 0x89, 0xa4, 0xb0, 0xb8, 0x8e,
	0x05, 0xcb, 0x5a, 0x32, 0xd4, 0x02, 0x3c, 0xae, 0xa8, 0x28, 0x68, 0x2d, 0x41, 0xc9, 0x68, 0x5e,
	0x32, 0x32, 0x9b, 0xe2, 0x28, 0x62, 0x4f, 0x0f, 0xc4, 0xac, 0x66, 0x4c, 0xf3, 0xfb, 0x63, 0xc0,
	0xe2, 0x6b, 0xea, 0xd3, 0x9c, 0x75, 0x43, 0x76, 0x6b, 0x8e, 0xf3, 0xa6, 0xa8, 0x57, 0x56, 0xe1,
	0xdd, 0xab, 0xe6, 0xb0, 0x1b, 0xca, 0x96, 0xef, 0xd5, 0x14, 0x8b, 0xb6, 0xbc, 0x65, 0x44, 0xcf,
	0xf0, 0x06, 0x16, 0x1f, 0x19, 0x67, 0x77, 0x1b, 0xe1, 0xf2, 0xef, 0x00, 0xec, 0xcf, 0x34, 0xa6,
	0xc1, 0x71, 0x19, 0xee, 0x61, 0xd6, 0xbd, 0x69, 0xf8, 0xac, 0x92, 0xbf, 0x6f, 0x11, 0xdd, 0xd9,
	0x91, 0x47, 0x95, 0x9d, 0xe4, 0xac, 0xf0, 0xc6, 0xf8, 0x30, 0xa8, 0x77, 0x54, 0xa0, 0x0f, 0x76,
	0xa7, 0x5f, 0x48, 0xaa, 0x1e, 0x3d, 0x6e, 0xba, 0x7a, 0x17, 0x1a, 0xfb, 0x45, 0xa6, 0x85, 0x37,
	0xc1, 0x71, 0xd8, 0xaa, 0x92, 0x28, 0x9d, 0x16, 0x1d, 0x50, 0x7a, 0x0c, 0xec, 0x41, 0xf9, 0xdd,
	0xaa, 0x92, 0x28, 0x9d, 0x2e, 0x1c, 0x50, 0x7a, 0x3c, 0xea, 0x41, 0xf1, 0x5b, 0x55, 0x18, 0x81,
	0xf3, 0xbf, 0x5b, 0x8e, 0xcf, 0x6b, 0x77, 0x7a, 0x3e, 0x03, 0x6e, 0xc7, 0x1d, 0x22, 0xb3, 0xc2,
	0x9b, 0xe2, 0x59, 0x70, 0x5c, 0xe8, 0xda, 0x85, 0x87, 0x30, 0xe6, 0xc9, 0x4f, 0xca, 0xf7, 0x89,
	0xc8, 0xdf, 0xbf, 0x5b, 0xaf, 0x5f, 0x4d, 0x8c, 0x1f, 0x27, 0xaa, 0xc3, 0xeb, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x9f, 0xe3, 0x63, 0xc9, 0x04, 0x00, 0x00,
}