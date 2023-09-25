package pipeline

import (
	"encoding/binary"
	"io"
)

// A Pipeline takes a binary input and turns it into usable data and
// vice versa.
type Pipeline interface {
	ReadPacket(id PacketId, reader io.Reader) (int, error)
	WritePacket(packet Packet, writer io.Writer) (int, error)
	SetOutput(output chan<- Packet)
}

// Kind of messages a pipeline produces. Used to identify each packet
// with a pipeline.
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

// Mux is a packet multiplexer for chosing the appropiate pipe
// both when encoding and decoding.
//
// When decoding, the mux decodes the packet ID and uses it to
// determine the pipe. On the other hand, when encoding, the mux
// uses the already provided packet id.
type Mux struct {
	pipes     map[Kind]Pipeline
	idToKind  map[PacketId]Kind
	byteOrder binary.ByteOrder
}

// Read packets from the reader until the source is drained.
// The source must provide valid packets, each starting with
// the ID as an uint16. After reading the ID, the pipe is responsible
// of only taking as much as it needs to read it. Failing to read
// exactly the packet data will cause further reads to return garbage
func (mux *Mux) ReadFrom(reader io.Reader) (int64, error) {
	totalRead := 0

	var err error = nil
	var n int
	for err == nil {
		n, err = mux.readNextPacket(reader)
		totalRead += n
	}

	return int64(totalRead), err
}

func (mux *Mux) readNextPacket(reader io.Reader) (int, error) {
	totalRead := 0

	idBuf := make([]byte, 2)
	n, err := reader.Read(idBuf)
	totalRead += n
	if err != nil {
		return totalRead, err
	}

	nextId := PacketId(mux.byteOrder.Uint16(idBuf))
	// FIXME: Check if id or kind exists in maps
	kind := mux.idToKind[nextId]
	pipe := mux.pipes[kind]

	n, err = pipe.ReadPacket(nextId, reader)
	totalRead += n

	return totalRead, err
}

func (mux *Mux) WritePacket(packet Packet, writer io.Writer) (int, error) {
	// FIXME: Check if id or kind exists in maps
	kind := mux.idToKind[packet.Id()]
	pipe := mux.pipes[kind]
	return pipe.WritePacket(packet, writer)
}
