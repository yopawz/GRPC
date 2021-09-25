// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc/service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InsertEmployeeParams struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Salary               int32    `protobuf:"varint,3,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertEmployeeParams) Reset()         { *m = InsertEmployeeParams{} }
func (m *InsertEmployeeParams) String() string { return proto.CompactTextString(m) }
func (*InsertEmployeeParams) ProtoMessage()    {}
func (*InsertEmployeeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{0}
}

func (m *InsertEmployeeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertEmployeeParams.Unmarshal(m, b)
}
func (m *InsertEmployeeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertEmployeeParams.Marshal(b, m, deterministic)
}
func (m *InsertEmployeeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertEmployeeParams.Merge(m, src)
}
func (m *InsertEmployeeParams) XXX_Size() int {
	return xxx_messageInfo_InsertEmployeeParams.Size(m)
}
func (m *InsertEmployeeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertEmployeeParams.DiscardUnknown(m)
}

var xxx_messageInfo_InsertEmployeeParams proto.InternalMessageInfo

func (m *InsertEmployeeParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertEmployeeParams) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *InsertEmployeeParams) GetSalary() int32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type DeleteEmployeeParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEmployeeParams) Reset()         { *m = DeleteEmployeeParams{} }
func (m *DeleteEmployeeParams) String() string { return proto.CompactTextString(m) }
func (*DeleteEmployeeParams) ProtoMessage()    {}
func (*DeleteEmployeeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{1}
}

func (m *DeleteEmployeeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEmployeeParams.Unmarshal(m, b)
}
func (m *DeleteEmployeeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEmployeeParams.Marshal(b, m, deterministic)
}
func (m *DeleteEmployeeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEmployeeParams.Merge(m, src)
}
func (m *DeleteEmployeeParams) XXX_Size() int {
	return xxx_messageInfo_DeleteEmployeeParams.Size(m)
}
func (m *DeleteEmployeeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEmployeeParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEmployeeParams proto.InternalMessageInfo

func (m *DeleteEmployeeParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAllEmployeesParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllEmployeesParams) Reset()         { *m = GetAllEmployeesParams{} }
func (m *GetAllEmployeesParams) String() string { return proto.CompactTextString(m) }
func (*GetAllEmployeesParams) ProtoMessage()    {}
func (*GetAllEmployeesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{2}
}

func (m *GetAllEmployeesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllEmployeesParams.Unmarshal(m, b)
}
func (m *GetAllEmployeesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllEmployeesParams.Marshal(b, m, deterministic)
}
func (m *GetAllEmployeesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllEmployeesParams.Merge(m, src)
}
func (m *GetAllEmployeesParams) XXX_Size() int {
	return xxx_messageInfo_GetAllEmployeesParams.Size(m)
}
func (m *GetAllEmployeesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllEmployeesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllEmployeesParams proto.InternalMessageInfo

type GetEmployeeWithIdParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEmployeeWithIdParams) Reset()         { *m = GetEmployeeWithIdParams{} }
func (m *GetEmployeeWithIdParams) String() string { return proto.CompactTextString(m) }
func (*GetEmployeeWithIdParams) ProtoMessage()    {}
func (*GetEmployeeWithIdParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{3}
}

func (m *GetEmployeeWithIdParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEmployeeWithIdParams.Unmarshal(m, b)
}
func (m *GetEmployeeWithIdParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEmployeeWithIdParams.Marshal(b, m, deterministic)
}
func (m *GetEmployeeWithIdParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEmployeeWithIdParams.Merge(m, src)
}
func (m *GetEmployeeWithIdParams) XXX_Size() int {
	return xxx_messageInfo_GetEmployeeWithIdParams.Size(m)
}
func (m *GetEmployeeWithIdParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEmployeeWithIdParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetEmployeeWithIdParams proto.InternalMessageInfo

func (m *GetEmployeeWithIdParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Employee struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Salary               int32    `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{4}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Employee) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Employee) GetSalary() int32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EmployeeList struct {
	Employees            []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EmployeeList) Reset()         { *m = EmployeeList{} }
func (m *EmployeeList) String() string { return proto.CompactTextString(m) }
func (*EmployeeList) ProtoMessage()    {}
func (*EmployeeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_983ccfe825f84659, []int{6}
}

func (m *EmployeeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeList.Unmarshal(m, b)
}
func (m *EmployeeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeList.Marshal(b, m, deterministic)
}
func (m *EmployeeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeList.Merge(m, src)
}
func (m *EmployeeList) XXX_Size() int {
	return xxx_messageInfo_EmployeeList.Size(m)
}
func (m *EmployeeList) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeList.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeList proto.InternalMessageInfo

func (m *EmployeeList) GetEmployees() []*Employee {
	if m != nil {
		return m.Employees
	}
	return nil
}

func init() {
	proto.RegisterType((*InsertEmployeeParams)(nil), "proto.InsertEmployeeParams")
	proto.RegisterType((*DeleteEmployeeParams)(nil), "proto.DeleteEmployeeParams")
	proto.RegisterType((*GetAllEmployeesParams)(nil), "proto.GetAllEmployeesParams")
	proto.RegisterType((*GetEmployeeWithIdParams)(nil), "proto.GetEmployeeWithIdParams")
	proto.RegisterType((*Employee)(nil), "proto.Employee")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*EmployeeList)(nil), "proto.EmployeeList")
}

func init() { proto.RegisterFile("proto/grpc/service.proto", fileDescriptor_983ccfe825f84659) }

var fileDescriptor_983ccfe825f84659 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0x5d, 0xdb, 0xfd, 0xbd, 0xbf, 0x1f, 0x9d, 0x5e, 0xa7, 0x2b, 0x53, 0x64, 0x04, 0x91, 0xf9,
	0xe0, 0x06, 0xd3, 0x57, 0x61, 0x8a, 0x6e, 0x0c, 0x14, 0xa4, 0x28, 0x8a, 0x6f, 0x71, 0xbd, 0xcc,
	0x42, 0xff, 0xd1, 0x14, 0x61, 0x5f, 0xd5, 0x4f, 0x23, 0xcb, 0x96, 0x8e, 0x75, 0xf1, 0x29, 0xc9,
	0xcd, 0xc9, 0xc9, 0x3d, 0xe7, 0x5c, 0x70, 0x92, 0x34, 0xce, 0xe2, 0xc1, 0x3c, 0x4d, 0x66, 0x03,
	0x41, 0xe9, 0xb7, 0x3f, 0xa3, 0xbe, 0x2c, 0x61, 0x45, 0x2e, 0xec, 0x05, 0x5a, 0xd3, 0x48, 0x50,
	0x9a, 0x3d, 0x84, 0x49, 0x10, 0x2f, 0x88, 0x9e, 0x79, 0xca, 0x43, 0x81, 0x08, 0xe5, 0x88, 0x87,
	0xe4, 0x18, 0x5d, 0xa3, 0xd7, 0x70, 0xe5, 0x1e, 0xf7, 0xc0, 0xe2, 0x73, 0x72, 0xcc, 0xae, 0xd1,
	0xab, 0xb8, 0xcb, 0x2d, 0x1e, 0x41, 0x55, 0xf0, 0x80, 0xa7, 0x0b, 0xc7, 0x92, 0xc5, 0xf5, 0x89,
	0x9d, 0x43, 0xeb, 0x9e, 0x02, 0xca, 0xa8, 0xc0, 0x6a, 0x83, 0xe9, 0x7b, 0x92, 0xb3, 0xe2, 0x9a,
	0xbe, 0xc7, 0xda, 0x70, 0x38, 0xa1, 0xec, 0x36, 0x08, 0x14, 0x4e, 0xac, 0x80, 0xec, 0x02, 0xda,
	0x13, 0xca, 0x7b, 0x7a, 0xf3, 0xb3, 0xaf, 0xa9, 0xf7, 0x07, 0xc7, 0x3b, 0xd4, 0x15, 0xae, 0x78,
	0x97, 0xab, 0x30, 0x77, 0x55, 0x58, 0x3a, 0x15, 0xe5, 0x2d, 0x15, 0x67, 0x50, 0x77, 0x49, 0x24,
	0x71, 0x24, 0x08, 0x1d, 0xa8, 0x85, 0x24, 0xc4, 0xf2, 0xe5, 0xca, 0x12, 0x75, 0x64, 0x37, 0xf0,
	0x5f, 0xfd, 0xff, 0xe8, 0x8b, 0x0c, 0x2f, 0xa1, 0x41, 0x4a, 0x8d, 0x63, 0x74, 0xad, 0xde, 0xbf,
	0x61, 0x73, 0xe5, 0x79, 0x5f, 0xe1, 0xdc, 0x0d, 0x62, 0xf8, 0x63, 0x82, 0xfd, 0xc4, 0x23, 0x3e,
	0xcf, 0xbd, 0xc2, 0x31, 0x34, 0x0b, 0xae, 0xe0, 0xc9, 0x9a, 0x41, 0xeb, 0x56, 0xe7, 0xa0, 0xc0,
	0xbf, 0xec, 0x83, 0x95, 0x70, 0x04, 0xf6, 0x76, 0xb6, 0x78, 0xbc, 0x06, 0xea, 0x22, 0xef, 0xa8,
	0x2e, 0x95, 0x66, 0x56, 0xc2, 0x6b, 0xb0, 0x5f, 0x13, 0x8f, 0x6f, 0x72, 0xc4, 0xa2, 0x14, 0xdd,
	0xab, 0x11, 0xd8, 0xdb, 0xe9, 0xe7, 0xff, 0xea, 0x86, 0x42, 0xc7, 0x30, 0x86, 0xfd, 0x9d, 0xf8,
	0xf1, 0x74, 0xe3, 0x81, 0x6e, 0x30, 0x3a, 0xc5, 0xd6, 0x58, 0xe9, 0xae, 0xf1, 0x51, 0xeb, 0x0f,
	0x64, 0xf5, 0xb3, 0x2a, 0x97, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x5d, 0xa4, 0x41,
	0x12, 0x03, 0x00, 0x00,
}