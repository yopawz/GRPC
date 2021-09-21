// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package go_proto_grpc

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
	return fileDescriptor_c33392ef2c1961ba, []int{0}
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
	return fileDescriptor_c33392ef2c1961ba, []int{1}
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
	return fileDescriptor_c33392ef2c1961ba, []int{2}
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
	return fileDescriptor_c33392ef2c1961ba, []int{3}
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
	return fileDescriptor_c33392ef2c1961ba, []int{4}
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
	return fileDescriptor_c33392ef2c1961ba, []int{5}
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
	proto.RegisterType((*Employee)(nil), "proto.Employee")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*EmployeeList)(nil), "proto.EmployeeList")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0x6d, 0x92, 0xb6, 0xb7, 0x9d, 0x7b, 0x49, 0x2f, 0xd3, 0xde, 0x6b, 0xa8, 0x3e, 0x94, 0xc5,
	0x8f, 0xbe, 0xa4, 0x85, 0xea, 0x9b, 0x08, 0x55, 0xfc, 0x40, 0x50, 0x90, 0xa0, 0x20, 0xbe, 0x94,
	0xb5, 0x1d, 0x43, 0x20, 0x1f, 0xcb, 0x6e, 0x10, 0xfa, 0x37, 0xfc, 0xc5, 0x92, 0x6d, 0xb7, 0xa5,
	0x31, 0x4f, 0x3b, 0x67, 0x39, 0x73, 0xe6, 0xcc, 0x1c, 0xe8, 0x0a, 0x99, 0xe5, 0xd9, 0x58, 0x91,
	0xfc, 0x8c, 0xe6, 0x34, 0xd2, 0x08, 0x1b, 0xfa, 0x61, 0xcf, 0xd0, 0xbb, 0x4f, 0x15, 0xc9, 0xfc,
	0x26, 0x11, 0x71, 0xb6, 0x24, 0x7a, 0xe2, 0x92, 0x27, 0x0a, 0x11, 0xea, 0x29, 0x4f, 0xc8, 0xb3,
	0x06, 0xd6, 0xb0, 0x1d, 0xe8, 0x1a, 0xff, 0x82, 0xc3, 0x43, 0xf2, 0xec, 0x81, 0x35, 0x6c, 0x04,
	0x45, 0x89, 0xff, 0xa1, 0xa9, 0x78, 0xcc, 0xe5, 0xd2, 0x73, 0xf4, 0xe7, 0x1a, 0xb1, 0x63, 0xe8,
	0x5d, 0x53, 0x4c, 0x39, 0x95, 0x54, 0x5d, 0xb0, 0xa3, 0x85, 0xd6, 0x6c, 0x04, 0x76, 0xb4, 0x60,
	0x7b, 0xf0, 0xef, 0x8e, 0xf2, 0xcb, 0x38, 0x36, 0x3c, 0xb5, 0x22, 0xb2, 0x57, 0x68, 0x99, 0xaf,
	0x72, 0xd3, 0xc6, 0x9a, 0xfd, 0xd3, 0x9a, 0x53, 0x65, 0xad, 0xbe, 0x63, 0xed, 0x10, 0x5a, 0x01,
	0x29, 0x91, 0xa5, 0x8a, 0xd0, 0x83, 0x5f, 0x09, 0x29, 0x55, 0x74, 0xae, 0xf6, 0x34, 0x90, 0x5d,
	0xc0, 0x1f, 0x33, 0xff, 0x21, 0x52, 0x39, 0xfa, 0xd0, 0x26, 0x63, 0xd1, 0xb3, 0x06, 0xce, 0xf0,
	0xf7, 0xa4, 0xb3, 0x3a, 0xe4, 0xc8, 0xf0, 0x82, 0x2d, 0x63, 0xf2, 0x65, 0x83, 0xfb, 0xc8, 0x53,
	0x1e, 0x6e, 0x0e, 0x80, 0xb7, 0xd0, 0x29, 0xad, 0x8a, 0x07, 0x6b, 0x85, 0xca, 0x13, 0xf4, 0xbb,
	0x25, 0xfd, 0xc2, 0x07, 0xab, 0xe1, 0x14, 0xdc, 0xdd, 0xc0, 0x70, 0x7f, 0x4d, 0xac, 0xca, 0xb1,
	0x6f, 0x5c, 0x9a, 0x9d, 0x59, 0x0d, 0xcf, 0xc0, 0x7d, 0x11, 0x0b, 0xbe, 0x0d, 0x07, 0xcb, 0xab,
	0x54, 0x75, 0x4d, 0xc1, 0xdd, 0x8d, 0x74, 0x33, 0xb7, 0x2a, 0xe9, 0x0a, 0x85, 0xab, 0x93, 0xb7,
	0xa3, 0x8f, 0x48, 0xaa, 0xdc, 0xe7, 0x42, 0x8c, 0xc3, 0xcc, 0xd7, 0x04, 0x3f, 0x94, 0x62, 0x7e,
	0x1e, 0x66, 0x33, 0x8d, 0x66, 0x05, 0x7a, 0x6f, 0xea, 0xfa, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0x90, 0x29, 0x41, 0xb8, 0x02, 0x00, 0x00,
}
