package data

import (
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

// Pipe if a pipe for both data and orders. It uses the structure specified in the
// ADE to decode and encode the packets. After this, it also converts the units of
// numeric values to those of the vehicle or the display./
type Pipe struct {
	codec  Codec
	output chan<- pipeline.Packet
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	totalRead := 0
	packet, n, err := pipe.codec.Decode(id, reader)
	totalRead += n
	if err != nil {
		return n, err
	}

	pipe.output <- packet

	return totalRead, nil
}
