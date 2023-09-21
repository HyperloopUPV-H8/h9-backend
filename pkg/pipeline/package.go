package pipeline

// PacketId is the ID of any packet
type PacketId uint16

// Packet is an abstraction over any kind of packet.
// It is used as a common interface returned by all
// pipes. To actually use its contents, it must be casted
// to the actuall struct.
type Packet interface {
	Id() PacketId
}
