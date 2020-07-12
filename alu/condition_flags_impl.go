package alu

const (
	parityMask byte = 0x01
	signMask   byte = 0x80
)

// ConditionFlagsImpl is a struct representing the ALU condition flags register.
type ConditionFlagsImpl struct {
	Zero           bool
	Sign           bool
	Parity         bool
	Carry          bool
	AuxillaryCarry bool
}

// ClearFlags clears all flags.
func (flags *ConditionFlagsImpl) ClearFlags() {
	flags.ClearZero()
	flags.ClearSign()
	flags.ClearParity()
	flags.ClearCarry()
	flags.ClearAuxillaryCarry()
}

/*
 * 8-bit STATUS WORD FORMAT
 * [0] CARRY
 * [1] 1
 * [2] PARITY
 * [3] 0
 * [4] AUX CARRY
 * [5] 0
 * [6] ZERO
 * [7] SIGN
 */

// CreateStatusWord generates an 8-bit status word from the flags' values.
func (flags *ConditionFlagsImpl) CreateStatusWord() uint8 {
	var statusWord uint8 = 0x02
	if flags.Carry {
		statusWord |= 0x01
	}
	if flags.Parity {
		statusWord |= 0x04
	}
	if flags.AuxillaryCarry {
		statusWord |= 0x10
	}
	if flags.Zero {
		statusWord |= 0x40
	}
	if flags.Sign {
		statusWord |= 0x80
	}
	return statusWord
}

// ApplyStatusWord updates the flags' values based on statusWord.
func (flags *ConditionFlagsImpl) ApplyStatusWord(statusWord uint8) {
	flags.Carry = (statusWord & 0x01) > 0
	flags.Parity = (statusWord & 0x04) > 0
	flags.AuxillaryCarry = (statusWord & 0x10) > 0
	flags.Zero = (statusWord & 0x40) > 0
	flags.Sign = (statusWord & 0x80) > 0
}

// IsZero returns the value of the Zero flag
func (flags *ConditionFlagsImpl) IsZero() bool {
	return flags.Zero
}

// SetZero sets the Zero flag.
func (flags *ConditionFlagsImpl) SetZero() {
	flags.Zero = true
}

// UpdateZero sets the Zero flag if provided result is zero
// and returns the value of the Zero flag.
func (flags *ConditionFlagsImpl) UpdateZero(result uint16) bool {
	flags.Zero = result&0xFF == 0
	return flags.Zero
}

// ClearZero clears the Zero flag.
func (flags *ConditionFlagsImpl) ClearZero() {
	flags.Zero = false
}

// IsSign returns the value of the Sign flag
func (flags *ConditionFlagsImpl) IsSign() bool {
	return flags.Sign
}

// SetSign sets the Sign flag (indicating the last ALU operation resulted in a negative value).
func (flags *ConditionFlagsImpl) SetSign() {
	flags.Sign = true
}

// UpdateSign updates the Sign flag if the bit in position 7
// is set to 1 and returns the value of the Sign flag.
func (flags *ConditionFlagsImpl) UpdateSign(result uint16) bool {
	flags.Sign = result&0x80 > 0
	return flags.Sign
}

// ClearSign clears the Sign flag.
func (flags *ConditionFlagsImpl) ClearSign() {
	flags.Sign = false
}

// IsParity returns the value of the Parity flag
func (flags *ConditionFlagsImpl) IsParity() bool {
	return flags.Parity
}

// SetParity sets the Parity flag
func (flags *ConditionFlagsImpl) SetParity() {
	flags.Parity = true
}

// UpdateParity updates the Parity flag based on the bit parity of the provided data
// argument and returns the Parity flag value.
func (flags *ConditionFlagsImpl) UpdateParity(data uint16) bool {
	var bitCount uint8

	for i := 16; i > 0; i-- {
		if byte(data&0xFF)&parityMask > 0 {
			bitCount++
		}
		data = data >> 1
	}

	flags.Parity = bitCount%2 == 0
	return flags.Parity
}

// ClearParity clears the Parity flag.
func (flags *ConditionFlagsImpl) ClearParity() {
	flags.Parity = false
}

// IsCarry returns the value of the Carry flag
func (flags *ConditionFlagsImpl) IsCarry() bool {
	return flags.Carry
}

// SetCarry sets the Carry flag.
func (flags *ConditionFlagsImpl) SetCarry() {
	flags.Carry = true
}

// UpdateCarry updates the Carry flag based on the value of the provided result
// and returns the value of the Carry flag.
func (flags *ConditionFlagsImpl) UpdateCarry(result uint16) bool {
	flags.Carry = result > 0xff
	return flags.Carry
}

// UpdateBorrow updates the Carry flag based on the value of the provided addends
// assuming that addend2 is being subtracted from addend1.
func (flags *ConditionFlagsImpl) UpdateBorrow(addend1 uint8, addend2 uint8) bool {
	addend1Bit7 := (addend1 & 0x80) >> 6
	addend2Bit7 := (addend2 & 0x80) >> 6
	flags.Carry = addend1Bit7 < addend2Bit7
	return flags.Carry
}

// UpdateCarryDoublePrecision updates the Carry flag based on the value of the
// provided result and returns the value of the Carry flag.
func (flags *ConditionFlagsImpl) UpdateCarryDoublePrecision(result uint32) bool {
	flags.Carry = result > 0xffff
	return flags.Carry
}

// ClearCarry clears the Carry flag.
func (flags *ConditionFlagsImpl) ClearCarry() {
	flags.Carry = false
}

// IsAuxillaryCarry returns the value of the Auxillary Carry flag
func (flags *ConditionFlagsImpl) IsAuxillaryCarry() bool {
	return flags.AuxillaryCarry
}

// SetAuxillaryCarry sets the Auxillary Carry flag.
func (flags *ConditionFlagsImpl) SetAuxillaryCarry() {
	flags.AuxillaryCarry = true
}

// ClearAuxillaryCarry clears the Auxillary Carry flag.
func (flags *ConditionFlagsImpl) ClearAuxillaryCarry() {
	flags.AuxillaryCarry = false
}
