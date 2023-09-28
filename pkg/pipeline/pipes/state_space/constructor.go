package state_space

import "encoding/binary"

func NewPipe(byteOrder binary.ByteOrder) Pipe {
	return Pipe{byteOrder: byteOrder}
}
