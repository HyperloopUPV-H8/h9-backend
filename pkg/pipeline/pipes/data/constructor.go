package data

import (
	"encoding/binary"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

func NewPipe(codec *Codec) Pipe {
	return Pipe{
		codec: codec,
	}
}

func NewCodec(byteOrder binary.ByteOrder) Codec {
	return Codec{
		packetStructures: make(map[pipeline.PacketId]PacketStructure),
		enums:            make(map[ValueName]EnumDescriptor),
		byteOrder:        byteOrder,
	}
}

func (codec *Codec) AddStructure(id pipeline.PacketId, structure PacketStructure) {
	codec.packetStructures[id] = structure
}

func (codec *Codec) RemoveStructure(id pipeline.PacketId) {
	delete(codec.packetStructures, id)
}

func (codec *Codec) ClearStructures() {
	codec.packetStructures = make(map[pipeline.PacketId]PacketStructure)
}

func (codec *Codec) AddEnum(name ValueName, descriptor EnumDescriptor) {
	codec.enums[name] = descriptor
}

func (codec *Codec) RemoveEnum(name ValueName) {
	delete(codec.enums, name)
}

func (codec *Codec) ClearEnums() {
	codec.enums = make(map[ValueName]EnumDescriptor)
}
