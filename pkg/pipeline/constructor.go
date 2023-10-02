package pipeline

import (
	"encoding/binary"
)

func NewMux(byteOrder binary.ByteOrder) Mux {
	return Mux{
		pipes:     make(map[Kind]Pipeline),
		idToKind:  make(map[PacketId]Kind),
		byteOrder: byteOrder,
	}
}

func (mux *Mux) AddPacket(id PacketId, kind Kind) {
	mux.idToKind[id] = kind
}

func (mux *Mux) RemovePacket(id PacketId) {
	delete(mux.idToKind, id)
}

func (mux *Mux) ClearPackets() {
	mux.idToKind = make(map[PacketId]Kind)
}

func (mux *Mux) AddPipe(kind Kind, pipe Pipeline) {
	mux.pipes[kind] = pipe
}

// SetPipeOutput sets the pipe for the specified kind pipe.
// Returns true if the pipe exists and false if it doesn't.
func (mux *Mux) SetPipeOutput(kind Kind, output chan<- Packet) bool {
	pipe, ok := mux.pipes[kind]
	if !ok {
		return false
	}

	pipe.SetOutput(output)
	return true
}

func (mux *Mux) RemovePipe(kind Kind) {
	delete(mux.pipes, kind)
}
