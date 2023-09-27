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

func (codec *Codec) Decode(id pipeline.PacketId, reader io.Reader) (Packet, int, error) {
	structure, ok := codec.packetStructures[id]
	if !ok {
		return Packet{id: id}, 0, ErrIdNotFound(id)
	}

	totalRead := 0
	packet := Packet{id: id, values: make(map[ValueName]Value, len(structure))}
	for _, descriptor := range structure {
		value, n, err := codec.decodeNext(descriptor, reader)
		totalRead += n
		packet.values[descriptor.Name] = value
		if err != nil {
			return packet, totalRead, err
		}
	}

	return packet, totalRead, nil
}

func (codec *Codec) Encode(packet Packet, writer io.Writer) (int, error) {
	id := packet.Id()

	structure, ok := codec.packetStructures[id]
	if !ok {
		return 0, ErrIdNotFound(id)
	}

	totalWrite := 0
	for _, descriptor := range structure {
		value, ok := packet.Value(descriptor.Name)
		if !ok {
			return totalWrite, ErrValueNotFound(descriptor.Name)
		}
		n, err := codec.encodeNext(descriptor, value, writer)
		totalWrite += n
		if err != nil {
			return totalWrite, err
		}
	}

	return totalWrite, nil
}
