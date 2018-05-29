package memory

// RegisterPair epresents an Intel 8080 register pair
type RegisterPair struct {
	low      *uint8
	high     *uint8
	register [2]uint8
}

// Init prepares the RegisterPair for use by assigning memory addresses to the low and high fields that correspond to register[0] and register[1].
func (reg *RegisterPair) Init() {
	reg.low = &reg.register[0]
	reg.high = &reg.register[1]
}

// Write writes 2 bytes of data to a register pair.
func (reg *RegisterPair) Write(buf []byte) (int, error) {
	reg.low = &reg.register[0]
	reg.high = &reg.register[1]
	return 2, nil
}

// WriteLow writes the given byte to the lower register of the RegisterPair.
func (reg *RegisterPair) WriteLow(buf byte) (int, error) {
	*reg.low = uint8(buf)
	return 1, nil
}

// WriteHigh writes the given byte to the higher register of the RegisterPair.
func (reg *RegisterPair) WriteHigh(buf byte) (int, error) {
	*reg.high = uint8(buf)
	return 1, nil
}

// Write16 writes a uint16 to a register pair.
func (reg *RegisterPair) Write16(buf uint16) (int, error) {
	*reg.low = uint8(buf & 0x00ff)
	*reg.high = uint8(buf >> 8)
	return 2, nil
}

// Read stores the low register's value in buf[0] and the high register's value in buf[1].
func (reg *RegisterPair) Read(buf []byte) (int, error) {
	buf[0] = *reg.low
	buf[1] = *reg.high
	return 2, nil
}

// Read16 combines the low and high registers of the RegisterPair into a uint16 value and stores it in the specified buffer.
func (reg *RegisterPair) Read16(buf *uint16) (int, error) {
	var low = uint16(*reg.low)
	var high = uint16(*reg.high)
	*buf = uint16((high&0x00ff)<<8) | uint16(low&0x00ff)
	return 2, nil
}
