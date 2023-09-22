package data

import (
	"encoding/binary"
	"io"
)

func (pipe *Pipe) decodeNext(descriptor ValueDescriptor, reader io.Reader) (val Value, n int, err error) {
	switch descriptor.Type {
	case Uint8:
		val, n, err = decodeNextValue[uint8](reader, pipe.byteOrder)
		val = float64(val.(uint8))
	case Uint16:
		val, n, err = decodeNextValue[uint16](reader, pipe.byteOrder)
		val = float64(val.(uint16))
	case Uint32:
		val, n, err = decodeNextValue[uint32](reader, pipe.byteOrder)
		val = float64(val.(uint32))
	case Uint64:
		val, n, err = decodeNextValue[uint64](reader, pipe.byteOrder)
		val = float64(val.(uint64))
	case Int8:
		val, n, err = decodeNextValue[int8](reader, pipe.byteOrder)
		val = float64(val.(int8))
	case Int16:
		val, n, err = decodeNextValue[int16](reader, pipe.byteOrder)
		val = float64(val.(int16))
	case Int32:
		val, n, err = decodeNextValue[int32](reader, pipe.byteOrder)
		val = float64(val.(int32))
	case Int64:
		val, n, err = decodeNextValue[int64](reader, pipe.byteOrder)
		val = float64(val.(int64))
	case Float32:
		val, n, err = decodeNextValue[float32](reader, pipe.byteOrder)
		val = float64(val.(float32))
	case Float64:
		val, n, err = decodeNextValue[float64](reader, pipe.byteOrder)
	case Bool:
		val, n, err = decodeNextValue[bool](reader, pipe.byteOrder)
	case Enum:
		enumDescriptor, ok := pipe.enums[descriptor.Name]
		if !ok {
			err = ErrEnumNotFound(descriptor.Name)
		}
		var variant byte
		variant, n, err = decodeNextValue[byte](reader, pipe.byteOrder)
		val = enumDescriptor[variant]
	default:
		err = ErrInvalidType(descriptor.Type)
	}

	return val, n, err
}

func decodeNextValue[T any](reader io.Reader, byteOrder binary.ByteOrder) (T, int, error) {
	var value T
	err := binary.Read(reader, byteOrder, &value)
	totalRead := binary.Size(value)
	if err == io.EOF {
		totalRead = 0
	}
	return value, totalRead, err
}
