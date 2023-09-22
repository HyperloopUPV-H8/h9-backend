package data

import (
	"fmt"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type ErrIdNotFound pipeline.PacketId

func (err ErrIdNotFound) Error() string {
	return fmt.Sprintf("Id %d not found", pipeline.PacketId(err))
}

type ErrInvalidType ValueType

func (err ErrInvalidType) Error() string { return fmt.Sprintf("Type %s is not valid", ValueType(err)) }

type ErrEnumNotFound ValueName

func (err ErrEnumNotFound) Error() string { return fmt.Sprintf("Enum %s not found", ValueName(err)) }
