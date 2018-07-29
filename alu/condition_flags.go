package alu

const (
	parityMask byte = 0x01
	signMask   byte = 0x80
)

// ConditionFlags is a struct representing the ALU condition flags register.
type ConditionFlags struct {
	Zero           bool
	Sign           bool
	Parity         bool
	Carry          bool
	AuxillaryCarry bool
}

// ClearFlags clears all flags.
func (flags *ConditionFlags) ClearFlags() {
	flags.ClearZero()
	flags.ClearSign()
	flags.ClearParity()
	flags.ClearCarry()
	flags.ClearAuxillaryCarry()
}

// SetZero sets the Zero flag.
func (flags *ConditionFlags) SetZero() {
	flags.Zero = true
}

// UpdateZero sets the Zero flag if provided result is zero
// and returns the value of the Zero flag.
func (flags *ConditionFlags) UpdateZero(result uint16) bool {
	flags.Zero = result&0xFF == 0
	return flags.Zero
}

// ClearZero clears the Zero flag.
func (flags *ConditionFlags) ClearZero() {
	flags.Zero = false
}

// SetSign sets the Sign flag (indicating the last ALU operation resulted in a negative value).
func (flags *ConditionFlags) SetSign() {
	flags.Sign = true
}

// UpdateSign updates the Sign flag if the bit in position 7
// is set to 1 and returns the value of the Sign flag.
func (flags *ConditionFlags) UpdateSign(result uint16) bool {
	flags.Sign = result&0x80 > 0
	return flags.Sign
}

// ClearSign clears the Sign flag.
func (flags *ConditionFlags) ClearSign() {
	flags.Sign = false
}

// UpdateParity updates the Parity flag based on the bit parity of the provided data
// argument and returns the Parity flag value.
func (flags *ConditionFlags) UpdateParity(data uint8) bool {
	var bitCount uint8

	for i := 8; i > 0; i-- {
		if byte(data)&parityMask > 0 {
			bitCount++
		}
		data = data >> 1
	}

	flags.Parity = bitCount%2 == 0
	return flags.Parity
}

// ClearParity clears the Parity flag.
func (flags *ConditionFlags) ClearParity() {
	flags.Parity = false
}

// SetCarry sets the Carry flag.
func (flags *ConditionFlags) SetCarry() {
	flags.Carry = true
}

// UpdateCarry updates the Carry flag based on the value of the provided result
// and returns the value of the Carry flag.
func (flags *ConditionFlags) UpdateCarry(result uint16) bool {
	flags.Carry = result > 0xff
	return flags.Carry
}

// ClearCarry clears the Carry flag.
func (flags *ConditionFlags) ClearCarry() {
	flags.Carry = false
}

// SetAuxillaryCarry sets the Auxillary Carry flag.
func (flags *ConditionFlags) SetAuxillaryCarry() {
	flags.AuxillaryCarry = true
}

// ClearAuxillaryCarry clears the Auxillary Carry flag.
func (flags *ConditionFlags) ClearAuxillaryCarry() {
	flags.AuxillaryCarry = false
}
