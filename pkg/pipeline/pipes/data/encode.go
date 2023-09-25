package data

import (
	"encoding/binary"
	"io"
)

func (codec *Codec) encodeNext(descriptor ValueDescriptor, value Value, writer io.Writer) (int, error) {
	switch descriptor.Type {
	case Uint8:
		return encodeNextNumeric[uint8](value, writer, codec.byteOrder)
	case Uint16:
		return encodeNextNumeric[uint16](value, writer, codec.byteOrder)
	case Uint32:
		return encodeNextNumeric[uint32](value, writer, codec.byteOrder)
	case Uint64:
		return encodeNextNumeric[uint64](value, writer, codec.byteOrder)
	case Int8:
		return encodeNextNumeric[int8](value, writer, codec.byteOrder)
	case Int16:
		return encodeNextNumeric[int16](value, writer, codec.byteOrder)
	case Int32:
		return encodeNextNumeric[int32](value, writer, codec.byteOrder)
	case Int64:
		return encodeNextNumeric[int64](value, writer, codec.byteOrder)
	case Float32:
		return encodeNextNumeric[float32](value, writer, codec.byteOrder)
	case Float64:
		return encodeNextNumeric[float64](value, writer, codec.byteOrder)
	case Bool:
		return encodeNextBool(value, writer, codec.byteOrder)
	case Enum:
		enum, ok := codec.enums[descriptor.Name]
		if !ok {
			return 0, ErrEnumNotFound(descriptor.Name)
		}
		return encodeNextEnum(enum, value, writer, codec.byteOrder)
	default:
		return 0, ErrInvalidType(descriptor.Type)
	}
}

type numeric interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

func encodeNextNumeric[T numeric](value Value, writer io.Writer, byteOrder binary.ByteOrder) (int, error) {
	val, ok := value.(float64)
	if !ok {
		return 0, ErrValueNotNumeric{value}
	}
	num := T(val)
	err := binary.Write(writer, byteOrder, num)
	return binary.Size(num), err
}

func encodeNextBool(value Value, writer io.Writer, byteOrder binary.ByteOrder) (int, error) {
	boolean, ok := value.(bool)
	if !ok {
		return 0, ErrValueNotBool{value}
	}
	err := binary.Write(writer, byteOrder, boolean)
	return binary.Size(boolean), err
}

func encodeNextEnum(enum EnumDescriptor, value Value, writer io.Writer, byteOrder binary.ByteOrder) (int, error) {
	variant, ok := value.(EnumVariant)
	if !ok {
		return 0, ErrValueNotEnum{value}
	}

	var repr uint8
	found := false
	for i, enumVariant := range enum {
		if variant == enumVariant {
			repr = uint8(i)
			found = true
			break
		}
	}

	if !found {
		return 0, ErrVariantNotFound(variant)
	}

	err := binary.Write(writer, byteOrder, repr)
	return binary.Size(repr), err
}
