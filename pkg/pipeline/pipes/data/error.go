package data

import (
	"fmt"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/pipeline"
)

type ErrIdNotFound pipeline.PacketId

func (err ErrIdNotFound) Error() string { return fmt.Sprintf("Id %d not found", err) }
