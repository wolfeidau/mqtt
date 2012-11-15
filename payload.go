package mqtt

import (
	"io"
)

// Payload is the interface for Publish payloads. Typically the BytesPayload
// implementation will be sufficient for small payloads whose full contents
// will exist in memory.
type Payload interface {
	// Size returns the number of bytes that WritePayload will write.
	Size() int

	// WritePayload writes the payload data to w.
	WritePayload(w io.Writer) error

	// ReadPayload reads the payload data from r.
	ReadPayload(r io.Reader) error
}

type BytesPayload []byte

func (p BytesPayload) Size() int {
	return len(p)
}

func (p BytesPayload) WritePayload(w io.Writer) error {
	_, err := w.Write(p)
	return err
}

func (p BytesPayload) ReadPayload(r io.Reader) error {
	_, err := io.ReadFull(r, p)
	return err
}
