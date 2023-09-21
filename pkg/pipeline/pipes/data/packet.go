package data

import "github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"

type Value any
type ValueName string
type ValueType string

type ValueDescriptor struct {
	Name ValueName
	Type ValueType
}

type PacketStructure []ValueDescriptor

type Packet struct {
	id     pipeline.PacketId
	values map[ValueName]Value
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }
