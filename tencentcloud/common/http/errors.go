package common

import "fmt"

type PacketTooLargeError struct {
	Size, Limit int64
}

func (e PacketTooLargeError) Error() string {
	return fmt.Sprintf("packet too large, limit: %d, size: %d", e.Limit, e.Size)
}
