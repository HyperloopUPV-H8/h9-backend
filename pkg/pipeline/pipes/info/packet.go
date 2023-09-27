package info

import "github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"

type Message string

type Timestamp struct {
	Counter uint16 `json:"counter"`
	Second  uint8  `json:"second"`
	Minute  uint8  `json:"minute"`
	Hour    uint8  `json:"hour"`
	Day     uint8  `json:"day"`
	Month   uint8  `json:"month"`
	Year    uint16 `json:"year"`
}

type Packet struct {
	id pipeline.PacketId
	// TODO: replace with better type
	BoardId   uint16    `json:"boardId"`
	Timestamp Timestamp `json:"timestamp"`
	Msg       Message   `json:"msg"`
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }
