package data

import (
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type Pipe struct {
	packetStructures map[pipeline.PacketId]PacketStructure
	output           chan<- pipeline.Packet
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	structure, ok := pipe.packetStructures[id]
	if !ok {
		return 0, ErrIdNotFound(id)
	}

	total_read := 0
	packet := Packet{id: id, values: make(map[ValueName]Value, len(structure))}
	for _, descriptor := range structure {
		value, n, err := decodeNext(descriptor, reader)
		total_read += n
		if err != nil {
			return total_read, err
		}
		packet.values[descriptor.Name] = value
	}

	return total_read, nil
}
