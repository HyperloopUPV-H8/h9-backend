package pipes

import (
	"encoding/binary"
	"io"
)

func ReadLength(byteOrder binary.ByteOrder, reader io.Reader) ([]byte, int, error) {
	var msgLength uint16
	err := binary.Read(reader, byteOrder, &msgLength)
	if err != nil {
		return []byte{}, 0, err
	}

	totalRead := binary.Size(msgLength)

	msgBuf := make([]byte, msgLength)
	m := 0
	for m < len(msgBuf) {
		n, err := reader.Read(msgBuf[m:msgLength])
		totalRead += n
		m += n
		if err != nil {
			return msgBuf, totalRead, err
		}
	}

	return msgBuf, totalRead, nil
}

type ErrNoEncoding struct{}

func (err ErrNoEncoding) Error() string {
	return "Pipeline does not suppoort encoding"
}
