package state_order

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline/pipes"
)

type Packet struct {
	id     pipeline.PacketId
	orders []pipeline.PacketId
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }

func (packet Packet) Orders() []pipeline.PacketId { return packet.orders }

type Pipe struct {
	output    chan<- pipeline.Packet
	byteOrder binary.ByteOrder
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	buf, totalRead, readErr := pipes.ReadLength(pipe.byteOrder, reader)
	if readErr != nil && readErr != io.EOF {
		return totalRead, readErr
	}

	orders := make([]pipeline.PacketId, binary.Size(buf)/binary.Size(pipeline.PacketId(0)))
	err := binary.Read(bytes.NewBuffer(buf), pipe.byteOrder, orders)
	if err != nil {
		return totalRead, err
	}

	pipe.output <- Packet{id, orders}

	return totalRead, nil
}

func (pipe *Pipe) WritePacket(packet pipeline.Packet, writer io.Writer) (int, error) {
	return 0, pipes.ErrNoEncoding{}
}
