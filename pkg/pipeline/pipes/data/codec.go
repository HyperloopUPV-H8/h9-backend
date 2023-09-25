package data

import (
	"encoding/binary"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type Codec struct {
	packetStructures map[pipeline.PacketId]PacketStructure
	enums            map[ValueName]EnumDescriptor
	byteOrder        binary.ByteOrder
}

func (codec Codec) Decode(id pipeline.PacketId, reader io.Reader) (Packet, int, error) {
	structure, ok := codec.packetStructures[id]
	if !ok {
		return Packet{id: id}, 0, ErrIdNotFound(id)
	}

	totalRead := 0
	packet := Packet{id: id, values: make(map[ValueName]Value, len(structure))}
	for _, descriptor := range structure {
		value, n, err := codec.decodeNext(descriptor, reader)
		totalRead += n
		if err != nil {
			return packet, totalRead, err
		}
		packet.values[descriptor.Name] = value
	}

	return packet, totalRead, nil
}
