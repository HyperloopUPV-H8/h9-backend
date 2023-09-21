package pipeline

import (
	"encoding/binary"
	"io"
)

type Pipeline interface {
	ReadPacket(id PacketId, reader io.Reader) (int, error)
	WritePacket(packet Packet, writer io.Writer) (int, error)
	SetOutput(output chan<- Packet)
}

type Kind string

const (
	Data       Kind = "data"
	Info       Kind = "info"
	Protection Kind = "protection"
	StateOrder Kind = "state_order"
	StateSpace Kind = "state_space"
	BlcuAck    Kind = "blcu_ack"
	Order      Kind = "order"
)

type Mux struct {
	pipes     map[Kind]Pipeline
	idToKind  map[PacketId]Kind
	byteOrder binary.ByteOrder
}

func (mux *Mux) ReadFrom(reader io.Reader) (int64, error) {
	total_read := 0

	var err error = nil
	var n int
	for err == nil {
		n, err = mux.readNextPacket(reader)
		total_read += n
	}

	return int64(total_read), err
}

func (mux *Mux) readNextPacket(reader io.Reader) (int, error) {
	total_read := 0

	idBuf := make([]byte, 2)
	n, err := reader.Read(idBuf)
	total_read += n
	if err != nil {
		return total_read, err
	}

	nextId := PacketId(mux.byteOrder.Uint16(idBuf))
	kind := mux.idToKind[nextId]
	pipe := mux.pipes[kind]

	n, err = pipe.ReadPacket(nextId, reader)
	total_read += n

	return total_read, err
}
