package pebble

const (
	ErrKeyNotFound = Error("key not found")
)

// Error represents an error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}
