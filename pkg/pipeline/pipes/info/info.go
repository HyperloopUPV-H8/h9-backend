package info

import (
	"encoding/binary"
	"encoding/json"
	"io"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type Pipe struct {
	byteOrder binary.ByteOrder
	output    chan<- pipeline.Packet
}

func (pipe *Pipe) SetOutput(output chan<- pipeline.Packet) { pipe.output = output }

func (pipe *Pipe) ReadPacket(id pipeline.PacketId, reader io.Reader) (int, error) {
	var msgLength uint16
	err := binary.Read(reader, pipe.byteOrder, &msgLength)
	if err != nil {
		return 0, err
	}

	totalRead := binary.Size(msgLength)

	msgBuf := make([]byte, msgLength)
	m := 0
	for m < len(msgBuf) {
		n, err := reader.Read(msgBuf[m:msgLength])
		totalRead += n
		m += n
		if err != nil {
			return totalRead, err
		}
	}

	var message Packet
	err = json.Unmarshal(msgBuf, &message)
	if err != nil {
		return totalRead, err
	}
	message.id = id

	pipe.output <- message

	return totalRead, nil
}

func (pipe *Pipe) WritePacket(packet pipeline.Packet, writer io.Writer) (int, error) {
	return 0, ErrNoEncoding{}
}

type ErrNoEncoding struct{}

func (err ErrNoEncoding) Error() string {
	return "Pipeline does not suppoort encoding"
}
