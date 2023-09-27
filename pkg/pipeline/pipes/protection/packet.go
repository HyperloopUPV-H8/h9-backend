package protection

import (
	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

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
	// TODO: change to better type
	BoardId    uint16    `json:"boardId"`
	Timestamp  Timestamp `json:"timestamp"`
	Protection Details   `json:"protection"`
}

type Name string
type Type string

type Details struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Data Data `json:"data"`
}

func (packet Packet) Id() pipeline.PacketId { return packet.id }

type Data interface {
	Type() Type
}

type OutOfBounds struct {
	Value  float64    `json:"value"`
	Bounds [2]float64 `json:"bounds"`
}

func (data OutOfBounds) Type() Type {
	return "OUT_OF_BOUNDS"
}

type LowerBound struct {
	Value float64 `json:"value"`
	Bound float64 `json:"bound"`
}

func (data LowerBound) Type() Type {
	return "LOWER_BOUND"
}

type UpperBound struct {
	Value float64 `json:"value"`
	Bound float64 `json:"bound"`
}

func (data UpperBound) Type() Type {
	return "UPPER_BOUND"
}

type Equals struct {
	Value float64 `json:"value"`
}

func (data Equals) Type() Type {
	return "EQUALS"
}

type NotEquals struct {
	Value float64 `json:"value"`
	Want  float64 `json:"want"`
}

func (data NotEquals) Type() Type {
	return "NOT_EQUALS"
}

type TimeAccumulation struct {
	Value     float64 `json:"value"`
	Bound     float64 `json:"bound"`
	TimeLimit float64 `json:"timelimit"`
}

func (data TimeAccumulation) Type() Type {
	return "TIME_ACCUMULATION"
}

type ErrorHandler string

func (data ErrorHandler) Type() Type {
	return "ERROR_HANDLER"
}
