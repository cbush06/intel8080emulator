package memory

import (
	"errors"
)

// RegisterPair presents an Intel 8080 register pair
type RegisterPair struct {
	Low  Register
	High Register
}

// NewRegisterPair creates a new RegisterPair and initializes its values
// such that High is set to data1 and Low is set to data2
func NewRegisterPair(data1 uint8, data2 uint8) *RegisterPair {
	return &RegisterPair{
		High: *NewRegister(data1),
		Low:  *NewRegister(data2),
	}
}

// Write writes 2 bytes of data to a register pair.
func (reg *RegisterPair) Write(buf []byte) (int, error) {
	if len(buf) != 2 {
		return 0, errors.New("buffer argument must be of type and []byte with length 2")
	}
	reg.Low.Write8(buf[0])
	reg.High.Write8(buf[1])
	return 2, nil
}

// WriteLow writes the given byte to the lower register of the RegisterPair.
func (reg *RegisterPair) WriteLow(buf byte) (int, error) {
	reg.Low.Write8(buf)
	return 1, nil
}

// WriteHigh writes the given byte to the higher register of the RegisterPair.
func (reg *RegisterPair) WriteHigh(buf byte) (int, error) {
	reg.High.Write8(buf)
	return 1, nil
}

// Write16 writes a uint16 to a register pair.
func (reg *RegisterPair) Write16(buf uint16) (int, error) {
	reg.Low.Write8(uint8(buf & 0x00ff))
	reg.High.Write8(uint8(buf >> 8))
	return 2, nil
}

// Read stores the low register's value in buf[0] and the high register's value in buf[1].
func (reg *RegisterPair) Read(buf []byte) (int, error) {
	if len(buf) != 2 {
		return 0, errors.New("buffer argument must be of type and []byte with length 2")
	}
	reg.Low.Read8(&buf[0])
	reg.High.Read8(&buf[1])
	return 2, nil
}

// ReadLow stores the low register's value in buf.
func (reg *RegisterPair) ReadLow(buf *uint8) (int, error) {
	reg.Low.Read8(buf)
	return 1, nil
}

// ReadHigh stores the high register's value in buf.
func (reg *RegisterPair) ReadHigh(buf *uint8) (int, error) {
	reg.High.Read8(buf)
	return 1, nil
}

// Read16 combines the low and high registers of the RegisterPair into a uint16 value and stores it in the specified buffer.
func (reg *RegisterPair) Read16(buf *uint16) (int, error) {
	var low uint8
	var high uint8

	reg.Low.Read8(&low)
	reg.High.Read8(&high)

	*buf = uint16(high)<<8 | uint16(low)
	return 2, nil
}
