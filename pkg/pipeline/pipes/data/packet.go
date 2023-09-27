package data

import "github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"

type Value any
type ValueName string
type ValueType string

const (
	Uint8   ValueType = "uint8"
	Uint16  ValueType = "uint16"
	Uint32  ValueType = "uint32"
	Uint64  ValueType = "uint64"
	Int8    ValueType = "int8"
	Int16   ValueType = "int16"
	Int32   ValueType = "int32"
	Int64   ValueType = "int64"
	Float32 ValueType = "float32"
	Float64 ValueType = "float64"
	Bool    ValueType = "bool"
	Enum    ValueType = "enum"
)

type EnumVariant string
type EnumDescriptor []EnumVariant

type ValueDescriptor struct {
	Name ValueName
	Type ValueType
}

type PacketStructure []ValueDescriptor

type Packet struct {
	id     pipeline.PacketId
	values map[ValueName]Value
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }

func (packet Packet) Value(valueName ValueName) (value Value, ok bool) {
	value, ok = packet.values[valueName]
	return value, ok
}
