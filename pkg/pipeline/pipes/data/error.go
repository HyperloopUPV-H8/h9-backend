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

func (err ErrInvalidType) Error() string {
	return fmt.Sprintf("Type %s is not valid", ValueType(err))
}

type ErrEnumNotFound ValueName

func (err ErrEnumNotFound) Error() string {
	return fmt.Sprintf("Enum %s not found", ValueName(err))
}

type ErrInvalidPacketType struct {
	packet pipeline.Packet
}

func (err ErrInvalidPacketType) Error() string {
	return fmt.Sprintf("Packet of type %T is invalid", err.packet)
}

type ErrValueNotFound ValueName

func (err ErrValueNotFound) Error() string {
	return fmt.Sprintf("Value %s not found", ValueName(err))
}

type ErrValueNotNumeric struct {
	value Value
}

func (err ErrValueNotNumeric) Error() string {
	return fmt.Sprintf("Type %T is not numeric", err.value)
}

type ErrValueNotBool struct {
	value Value
}

func (err ErrValueNotBool) Error() string {
	return fmt.Sprintf("Type %T is not boolean", err.value)
}

type ErrValueNotEnum struct {
	value Value
}

func (err ErrValueNotEnum) Error() string {
	return fmt.Sprintf("Type %T is not enum", err.value)
}

type ErrVariantNotFound EnumVariant

func (err ErrVariantNotFound) Error() string {
	return fmt.Sprintf("Variant %s not found", EnumVariant(err))
}
