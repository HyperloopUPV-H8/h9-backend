package pipeline

import (
	"encoding/binary"
)

func NewMux(idToKind map[PacketId]Kind, kindToPipe map[Kind]Pipeline, byteOrder binary.ByteOrder) Mux {
	return Mux{
		pipes:     kindToPipe,
		idToKind:  idToKind,
		byteOrder: byteOrder,
	}
}
