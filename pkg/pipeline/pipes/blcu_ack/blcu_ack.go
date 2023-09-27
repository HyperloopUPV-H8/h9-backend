package blcuAck

import (
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type Packet struct {
	id pipeline.PacketId
}

func (packet Packet) Id() pipeline.PacketId {
	return packet.id
}

type Pipe struct {
	output chan<- pipeline.Packet
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	pipe.output <- Packet{id}

	return 0, nil
}

func (pipe *Pipe) WritePacket(packet pipeline.Packet, writer io.Writer) (int, error) {
	return 0, ErrNoEncoding{}
}

type ErrNoEncoding struct{}

func (err ErrNoEncoding) Error() string { return "Pipe does not support encoding" }
