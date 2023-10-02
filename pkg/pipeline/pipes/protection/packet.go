package protection

import (
	"encoding/json"

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

func (packet Packet) Id() pipeline.PacketId { return packet.id }

type Name string
type Type string

type Details struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Data Data `json:"data"`
}

type detailsAdapter struct {
	Name Name            `json:"name"`
	Type Type            `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (details *Details) UnmarshalJSON(data []byte) error {
	var adapter detailsAdapter
	err := json.Unmarshal(data, &adapter)
	if err != nil {
		return err
	}

	var protection Data
	switch adapter.Type {
	case OutOfBounds{}.Type():
		var outOfBounds OutOfBounds
		err = json.Unmarshal(adapter.Data, &outOfBounds)
		protection = outOfBounds
	case LowerBound{}.Type():
		var lowerBound LowerBound
		err = json.Unmarshal(adapter.Data, &lowerBound)
		protection = lowerBound
	case UpperBound{}.Type():
		var upperBound UpperBound
		err = json.Unmarshal(adapter.Data, &upperBound)
		protection = upperBound
	case Equals{}.Type():
		var equals Equals
		err = json.Unmarshal(adapter.Data, &equals)
		protection = equals
	case NotEquals{}.Type():
		var notEquals NotEquals
		err = json.Unmarshal(adapter.Data, &notEquals)
		protection = notEquals
	case TimeAccumulation{}.Type():
		var timeAccumulation TimeAccumulation
		err = json.Unmarshal(adapter.Data, &timeAccumulation)
		protection = timeAccumulation
	case ErrorHandler("").Type():
		var errorHandler ErrorHandler
		err = json.Unmarshal(adapter.Data, &errorHandler)
		protection = errorHandler
	}

	details.Name = adapter.Name
	details.Type = adapter.Type
	details.Data = protection

	return err
}

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
