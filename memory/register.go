package memory

import "errors"

// Register is an 8-bit block of memory
type Register struct {
	data uint8
}

// NewRegister creates a new Register and initializes its value
func NewRegister(data uint8) *Register {
	return &Register{
		data: data,
	}
}

// NewRegisterPair creates a new RegisterPair and initializes its values
// such that High is set to data1 and Low is set to data2
func NewRegisterPair(data1 uint8, data2 uint8) *RegisterPair {
	return &RegisterPair{
		High: *NewRegister(data1),
		Low:  *NewRegister(data2),
	}
}

// Write writes a single 8-bit byte to the register
func (r *Register) Write(buf []byte) (int, error) {
	if len(buf) != 1 {
		return 0, errors.New("buffer argument must be of type and []byte with length 1")
	}
	r.data = buf[0]
	return 1, nil
}

// Write8 writes a single 8-bit byte to the register
func (r *Register) Write8(buf uint8) (int, error) {
	r.data = buf
	return 1, nil
}

// Read reads the register into an 8-bit buffer
func (r *Register) Read(buf []byte) (int, error) {
	if len(buf) != 1 {
		return 0, errors.New("buffer argument must be of type and []byte with length 1")
	}
	buf[0] = r.data
	return 1, nil
}

// Read8 reads a single 8-bit byte into a uint8
func (r *Register) Read8(buf *uint8) (int, error) {
	*buf = r.data
	return 1, nil
}
