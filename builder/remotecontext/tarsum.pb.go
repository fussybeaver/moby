// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tarsum.proto

package remotecontext

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TarsumBackup struct {
	Hashes map[string]string `protobuf:"bytes,1,rep,name=Hashes,proto3" json:"Hashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TarsumBackup) Reset()      { *m = TarsumBackup{} }
func (*TarsumBackup) ProtoMessage() {}
func (*TarsumBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_375be7466b474c4d, []int{0}
}
func (m *TarsumBackup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TarsumBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TarsumBackup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TarsumBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarsumBackup.Merge(m, src)
}
func (m *TarsumBackup) XXX_Size() int {
	return m.Size()
}
func (m *TarsumBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_TarsumBackup.DiscardUnknown(m)
}

var xxx_messageInfo_TarsumBackup proto.InternalMessageInfo

func (m *TarsumBackup) GetHashes() map[string]string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func init() {
	proto.RegisterType((*TarsumBackup)(nil), "remotecontext.TarsumBackup")
	proto.RegisterMapType((map[string]string)(nil), "remotecontext.TarsumBackup.HashesEntry")
}

func init() { proto.RegisterFile("tarsum.proto", fileDescriptor_375be7466b474c4d) }

var fileDescriptor_375be7466b474c4d = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x49, 0x2c, 0x2a,
	0x2e, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x4a, 0xcd, 0xcd, 0x2f, 0x49,
	0x4d, 0xce, 0xcf, 0x2b, 0x49, 0xad, 0x28, 0x51, 0xea, 0x62, 0xe4, 0xe2, 0x09, 0x01, 0xcb, 0x3b,
	0x25, 0x26, 0x67, 0x97, 0x16, 0x08, 0xd9, 0x73, 0xb1, 0x79, 0x24, 0x16, 0x67, 0xa4, 0x16, 0x4b,
	0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xa9, 0xeb, 0xa1, 0x68, 0xd0, 0x43, 0x56, 0xac, 0x07, 0x51,
	0xe9, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0xd5, 0x26, 0x65, 0xc9, 0xc5, 0x8d, 0x24, 0x2c, 0x24,
	0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x89,
	0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5, 0x20, 0x1c, 0x2b, 0x26, 0x0b,
	0x46, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1,
	0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x8f, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xc6, 0xaf, 0xd0, 0xe8, 0x00, 0x00, 0x00,
}

func (this *TarsumBackup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TarsumBackup)
	if !ok {
		that2, ok := that.(TarsumBackup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Hashes) != len(that1.Hashes) {
		return false
	}
	for i := range this.Hashes {
		if this.Hashes[i] != that1.Hashes[i] {
			return false
		}
	}
	return true
}
func (this *TarsumBackup) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&remotecontext.TarsumBackup{")
	keysForHashes := make([]string, 0, len(this.Hashes))
	for k, _ := range this.Hashes {
		keysForHashes = append(keysForHashes, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForHashes)
	mapStringForHashes := "map[string]string{"
	for _, k := range keysForHashes {
		mapStringForHashes += fmt.Sprintf("%#v: %#v,", k, this.Hashes[k])
	}
	mapStringForHashes += "}"
	if this.Hashes != nil {
		s = append(s, "Hashes: "+mapStringForHashes+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTarsum(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TarsumBackup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TarsumBackup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TarsumBackup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for k := range m.Hashes {
			v := m.Hashes[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintTarsum(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTarsum(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTarsum(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTarsum(dAtA []byte, offset int, v uint64) int {
	offset -= sovTarsum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TarsumBackup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for k, v := range m.Hashes {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTarsum(uint64(len(k))) + 1 + len(v) + sovTarsum(uint64(len(v)))
			n += mapEntrySize + 1 + sovTarsum(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTarsum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTarsum(x uint64) (n int) {
	return sovTarsum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TarsumBackup) String() string {
	if this == nil {
		return "nil"
	}
	keysForHashes := make([]string, 0, len(this.Hashes))
	for k, _ := range this.Hashes {
		keysForHashes = append(keysForHashes, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForHashes)
	mapStringForHashes := "map[string]string{"
	for _, k := range keysForHashes {
		mapStringForHashes += fmt.Sprintf("%v: %v,", k, this.Hashes[k])
	}
	mapStringForHashes += "}"
	s := strings.Join([]string{`&TarsumBackup{`,
		`Hashes:` + mapStringForHashes + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTarsum(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TarsumBackup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTarsum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TarsumBackup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TarsumBackup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTarsum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTarsum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTarsum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hashes == nil {
				m.Hashes = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTarsum
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTarsum
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTarsum
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTarsum
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTarsum
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTarsum
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTarsum
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTarsum(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTarsum
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Hashes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTarsum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTarsum
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
func skipTarsum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTarsum
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
					return 0, ErrIntOverflowTarsum
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTarsum
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
			if length < 0 {
				return 0, ErrInvalidLengthTarsum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTarsum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTarsum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTarsum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTarsum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTarsum = fmt.Errorf("proto: unexpected end of group")
)
