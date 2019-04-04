// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/dbnode/generated/proto/namespace/schema.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package namespace

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SchemaOptions contains schema information for a namespace.
type SchemaOptions struct {
	// history contains a history of deployed schema definitions.
	History *SchemaHistory `protobuf:"bytes,1,opt,name=history" json:"history,omitempty"`
	// defaultMessageName identifies the proto message that contains the default schema for the namespace.
	DefaultMessageName string `protobuf:"bytes,2,opt,name=defaultMessageName,proto3" json:"defaultMessageName,omitempty"`
}

func (m *SchemaOptions) Reset()                    { *m = SchemaOptions{} }
func (m *SchemaOptions) String() string            { return proto.CompactTextString(m) }
func (*SchemaOptions) ProtoMessage()               {}
func (*SchemaOptions) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *SchemaOptions) GetHistory() *SchemaHistory {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *SchemaOptions) GetDefaultMessageName() string {
	if m != nil {
		return m.DefaultMessageName
	}
	return ""
}

// SchemaHistory is versioned FileDescriptorSet.
type SchemaHistory struct {
	// versions is a list of FileDescriptorSet sorted by version in ascending order.
	Versions []*FileDescriptorSet `protobuf:"bytes,1,rep,name=versions" json:"versions,omitempty"`
}

func (m *SchemaHistory) Reset()                    { *m = SchemaHistory{} }
func (m *SchemaHistory) String() string            { return proto.CompactTextString(m) }
func (*SchemaHistory) ProtoMessage()               {}
func (*SchemaHistory) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{1} }

func (m *SchemaHistory) GetVersions() []*FileDescriptorSet {
	if m != nil {
		return m.Versions
	}
	return nil
}

// FileDescriptorSet is a set of proto file descriptors.
type FileDescriptorSet struct {
	// id identifies a deployed version of FileDescriptorSet.
	DeployId string `protobuf:"bytes,1,opt,name=deployId,proto3" json:"deployId,omitempty"`
	// prevId identifies the previous deploy id of FileDescriptorSet.
	PrevId string `protobuf:"bytes,2,opt,name=prevId,proto3" json:"prevId,omitempty"`
	// descriptors is a list of proto file descriptors sorted by dependency in topological order.
	Descriptors [][]byte `protobuf:"bytes,3,rep,name=descriptors" json:"descriptors,omitempty"`
}

func (m *FileDescriptorSet) Reset()                    { *m = FileDescriptorSet{} }
func (m *FileDescriptorSet) String() string            { return proto.CompactTextString(m) }
func (*FileDescriptorSet) ProtoMessage()               {}
func (*FileDescriptorSet) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{2} }

func (m *FileDescriptorSet) GetDeployId() string {
	if m != nil {
		return m.DeployId
	}
	return ""
}

func (m *FileDescriptorSet) GetPrevId() string {
	if m != nil {
		return m.PrevId
	}
	return ""
}

func (m *FileDescriptorSet) GetDescriptors() [][]byte {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*SchemaOptions)(nil), "namespace.SchemaOptions")
	proto.RegisterType((*SchemaHistory)(nil), "namespace.SchemaHistory")
	proto.RegisterType((*FileDescriptorSet)(nil), "namespace.FileDescriptorSet")
}
func (m *SchemaOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.History != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.History.Size()))
		n1, err := m.History.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.DefaultMessageName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.DefaultMessageName)))
		i += copy(dAtA[i:], m.DefaultMessageName)
	}
	return i, nil
}

func (m *SchemaHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaHistory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, msg := range m.Versions {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSchema(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *FileDescriptorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileDescriptorSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeployId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.DeployId)))
		i += copy(dAtA[i:], m.DeployId)
	}
	if len(m.PrevId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.PrevId)))
		i += copy(dAtA[i:], m.PrevId)
	}
	if len(m.Descriptors) > 0 {
		for _, b := range m.Descriptors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSchema(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SchemaOptions) Size() (n int) {
	var l int
	_ = l
	if m.History != nil {
		l = m.History.Size()
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.DefaultMessageName)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *SchemaHistory) Size() (n int) {
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, e := range m.Versions {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func (m *FileDescriptorSet) Size() (n int) {
	var l int
	_ = l
	l = len(m.DeployId)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.PrevId)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, b := range m.Descriptors {
			l = len(b)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SchemaOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.History == nil {
				m.History = &SchemaHistory{}
			}
			if err := m.History.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMessageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultMessageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, &FileDescriptorSet{})
			if err := m.Versions[len(m.Versions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileDescriptorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileDescriptorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileDescriptorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeployId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, make([]byte, postIndex-iNdEx))
			copy(m.Descriptors[len(m.Descriptors)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchema
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSchema(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/dbnode/generated/proto/namespace/schema.proto", fileDescriptorSchema)
}

var fileDescriptorSchema = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xec, 0x30,
	0x14, 0x86, 0x6f, 0xee, 0xc0, 0x38, 0x4d, 0x15, 0x34, 0x0b, 0x29, 0x22, 0xa5, 0x74, 0xd5, 0x55,
	0x03, 0xed, 0xc6, 0xb5, 0x0c, 0x62, 0x17, 0x2a, 0x74, 0x9e, 0x20, 0x6d, 0x8e, 0x6d, 0xa0, 0x69,
	0x42, 0x92, 0x19, 0x98, 0xb7, 0xf0, 0xb1, 0x5c, 0xfa, 0x08, 0x52, 0x5f, 0x44, 0xec, 0x74, 0xca,
	0x80, 0x2e, 0xcf, 0xff, 0x7d, 0x9c, 0xf3, 0x27, 0x78, 0xdd, 0x08, 0xd7, 0x6e, 0xab, 0xb4, 0x56,
	0x92, 0xca, 0x9c, 0x57, 0x54, 0xe6, 0xd4, 0x9a, 0x9a, 0xf2, 0xaa, 0x57, 0x1c, 0x68, 0x03, 0x3d,
	0x18, 0xe6, 0x80, 0x53, 0x6d, 0x94, 0x53, 0xb4, 0x67, 0x12, 0xac, 0x66, 0x35, 0x50, 0x5b, 0xb7,
	0x20, 0x59, 0x3a, 0xc6, 0xc4, 0x9b, 0xf3, 0xd8, 0xe2, 0x8b, 0xcd, 0x88, 0x5e, 0xb4, 0x13, 0xaa,
	0xb7, 0x24, 0xc3, 0x67, 0xad, 0xb0, 0x4e, 0x99, 0x7d, 0x80, 0x22, 0x94, 0xf8, 0x59, 0x90, 0xce,
	0x76, 0x7a, 0x50, 0x1f, 0x0f, 0xbc, 0x3c, 0x8a, 0x24, 0xc5, 0x84, 0xc3, 0x2b, 0xdb, 0x76, 0xee,
	0x09, 0xac, 0x65, 0x0d, 0x3c, 0x33, 0x09, 0xc1, 0xff, 0x08, 0x25, 0x5e, 0xf9, 0x07, 0x89, 0x8b,
	0xe3, 0xd1, 0x69, 0x13, 0xb9, 0xc3, 0xab, 0x1d, 0x18, 0xfb, 0x53, 0x20, 0x40, 0xd1, 0x22, 0xf1,
	0xb3, 0xdb, 0x93, 0xab, 0x0f, 0xa2, 0x83, 0x35, 0xd8, 0xda, 0x08, 0xed, 0x94, 0xd9, 0x80, 0x2b,
	0x67, 0x3b, 0x16, 0xf8, 0xea, 0x17, 0x26, 0x37, 0x78, 0xc5, 0x41, 0x77, 0x6a, 0x5f, 0xf0, 0xf1,
	0x11, 0x5e, 0x39, 0xcf, 0xe4, 0x1a, 0x2f, 0xb5, 0x81, 0x5d, 0xc1, 0xa7, 0x7e, 0xd3, 0x44, 0x22,
	0xec, 0xf3, 0x79, 0x89, 0x0d, 0x16, 0xd1, 0x22, 0x39, 0x2f, 0x4f, 0xa3, 0xfb, 0xcb, 0xf7, 0x21,
	0x44, 0x1f, 0x43, 0x88, 0x3e, 0x87, 0x10, 0xbd, 0x7d, 0x85, 0xff, 0xaa, 0xe5, 0xf8, 0x9d, 0xf9,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xeb, 0x4b, 0x47, 0x96, 0x01, 0x00, 0x00,
}