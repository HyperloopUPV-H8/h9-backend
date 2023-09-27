package protection

import (
	"encoding/binary"
	"encoding/json"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline/pipes"
)

type Pipe struct {
	output    chan<- pipeline.Packet
	byteOrder binary.ByteOrder
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	msgBuf, totalRead, readErr := pipes.ReadLength(pipe.byteOrder, reader)
	if readErr != nil && readErr != io.EOF {
		return totalRead, readErr
	}

	var message Packet
	err := json.Unmarshal(msgBuf, &message)
	if err != nil {
		return totalRead, err
	}
	message.id = id

	pipe.output <- message

	return totalRead, readErr
}

func (pipe *Pipe) WritePacket(packet pipeline.Packet, writer io.Writer) (int, error) {
	return 0, pipes.ErrNoEncoding{}
}
