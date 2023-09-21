package pipeline

type PacketId uint16

type Packet interface {
	Id() PacketId
}
