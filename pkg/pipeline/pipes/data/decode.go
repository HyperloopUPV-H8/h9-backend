package data

import "io"

func decodeNext(descriptor ValueDescriptor, reader io.Reader) (Value, int, error) {
	return Value(""), 0, nil
}
