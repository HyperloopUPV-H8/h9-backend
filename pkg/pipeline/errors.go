package pipeline

import "fmt"

type ErrIdNotFound PacketId

func (err ErrIdNotFound) Error() string {
	return fmt.Sprintf("Id %d not found", PacketId(err))
}

type ErrKindNotFound Kind

func (err ErrKindNotFound) Error() string {
	return fmt.Sprintf("Pipe for kind %s not found", Kind(err))
}
