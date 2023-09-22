package data

import (
	"encoding/binary"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type Pipe struct {
	packetStructures map[pipeline.PacketId]PacketStructure
	enums            map[ValueName]EnumDescriptor
	output           chan<- pipeline.Packet
	byteOrder        binary.ByteOrder
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	structure, ok := pipe.packetStructures[id]
	if !ok {
		return 0, ErrIdNotFound(id)
	}

	totalRead := 0
	packet := Packet{id: id, values: make(map[ValueName]Value, len(structure))}
	for _, descriptor := range structure {
		value, n, err := pipe.decodeNext(descriptor, reader)
		totalRead += n
		if err != nil {
			return totalRead, err
		}
		packet.values[descriptor.Name] = value
	}

	pipe.output <- packet

	return totalRead, nil
}
