package state_space

import (
	"encoding/binary"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline/pipes"
)

type StateSpace [8][15]float32

type Packet struct {
	id    pipeline.PacketId
	state StateSpace
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }

func (packet Packet) State() StateSpace { return packet.state }

type Pipe struct {
	output    chan<- pipeline.Packet
	byteOrder binary.ByteOrder
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	var state StateSpace
	err := binary.Read(reader, pipe.byteOrder, state)
	totalRead := binary.Size(state)

	if err == io.EOF {
		return 0, err
	} else if err != nil {
		return totalRead, err
	}

	pipe.output <- Packet{id, state}

	return totalRead, nil
}

func (pipe *Pipe) WritePacket(packet pipeline.Packet, writer io.Writer) (int, error) {
	return 0, pipes.ErrNoEncoding{}
}
