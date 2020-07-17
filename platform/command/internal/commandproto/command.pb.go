// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: command.proto

/*
	Package commandproto is a generated protocol buffer package.

	It is generated from these files:
		command.proto

	It has these top-level messages:
		Event
*/
package commandproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time         int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	DeviceUdid   string `protobuf:"bytes,4,opt,name=device_udid,json=deviceUdid,proto3" json:"device_udid,omitempty"`
	PayloadBytes []byte `protobuf:"bytes,5,opt,name=payload_bytes,json=payloadBytes,proto3" json:"payload_bytes,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorCommand, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Event) GetDeviceUdid() string {
	if m != nil {
		return m.DeviceUdid
	}
	return ""
}

func (m *Event) GetPayloadBytes() []byte {
	if m != nil {
		return m.PayloadBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "commandproto.Event")
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommand(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCommand(dAtA, i, uint64(m.Time))
	}
	if len(m.DeviceUdid) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCommand(dAtA, i, uint64(len(m.DeviceUdid)))
		i += copy(dAtA[i:], m.DeviceUdid)
	}
	if len(m.PayloadBytes) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCommand(dAtA, i, uint64(len(m.PayloadBytes)))
		i += copy(dAtA[i:], m.PayloadBytes)
	}
	return i, nil
}

func encodeVarintCommand(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCommand(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovCommand(uint64(m.Time))
	}
	l = len(m.DeviceUdid)
	if l > 0 {
		n += 1 + l + sovCommand(uint64(l))
	}
	l = len(m.PayloadBytes)
	if l > 0 {
		n += 1 + l + sovCommand(uint64(l))
	}
	return n
}

func sovCommand(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommand(x uint64) (n int) {
	return sovCommand(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommand
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommand
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
				return ErrInvalidLengthCommand
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommand
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceUdid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommand
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
				return ErrInvalidLengthCommand
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceUdid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommand
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
				return ErrInvalidLengthCommand
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayloadBytes = append(m.PayloadBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.PayloadBytes == nil {
				m.PayloadBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommand(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommand
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
func skipCommand(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommand
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
					return 0, ErrIntOverflowCommand
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
					return 0, ErrIntOverflowCommand
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
				return 0, ErrInvalidLengthCommand
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommand
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
				next, err := skipCommand(dAtA[start:])
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
	ErrInvalidLengthCommand = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommand   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("command.proto", fileDescriptorCommand) }

var fileDescriptorCommand = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x72, 0xc1, 0x3c, 0xa5,
	0x42, 0x2e, 0x56, 0xd7, 0xb2, 0xd4, 0xbc, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xcc, 0xdc, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x30, 0x5b, 0x48, 0x9e, 0x8b, 0x3b, 0x25, 0xb5, 0x2c,
	0x33, 0x39, 0x35, 0xbe, 0x34, 0x25, 0x33, 0x45, 0x82, 0x05, 0xac, 0x98, 0x0b, 0x22, 0x14, 0x9a,
	0x92, 0x99, 0x22, 0xa4, 0xcc, 0xc5, 0x5b, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0x12, 0x9f, 0x54,
	0x59, 0x92, 0x5a, 0x2c, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x03, 0x15, 0x74, 0x02, 0x89,
	0x39, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33,
	0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xdd, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xca,
	0xac, 0xde, 0xaa, 0x00, 0x00, 0x00,
}