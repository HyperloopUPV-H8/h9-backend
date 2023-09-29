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

func NewCodec(structures map[pipeline.PacketId]PacketStructure, enums map[ValueName]EnumDescriptor, byteOrder binary.ByteOrder) Codec {
	return Codec{
		packetStructures: structures,
		enums:            enums,
		byteOrder:        byteOrder,
	}
}
