package alu

const (
	parityMask      byte   = 0x01
	signMask        byte   = 0x80
	auxCarryMask1   uint8 = 0x08
	auxCarryMask2    uint8  = 0x10
	carryMask       uint8  = 0x80
	doubleCarryMask uint16 = 0x8000
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
	flags.ClearAuxiliaryCarry()
}

/*
 * 8-bit STATUS WORD FORMAT
 *
 *   7   6   5    4   3   2   1    0
 * +---+---+---+----+---+---+---+----+
 * | S | Z | 0 | AC | 0 | P | 1 | CY |
 * +---+---+---+----+---+---+---+----+
 *
 * See: page 4-13 of Intel 8080 User Manual
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
func (flags *ConditionFlagsImpl) UpdateZero(result uint8) bool {
	flags.Zero = result == 0
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
func (flags *ConditionFlagsImpl) UpdateSign(result uint8) bool {
	flags.Sign = result&signMask > 0
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
func (flags *ConditionFlagsImpl) UpdateParity(data uint8) bool {
	var bitCount uint8

	for i := 8; i > 0; i-- {
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

// UpdateCarry updates the Carry flag based on the change from the original
// value to the new value. It is set if a carry occurs out-of bit 7 (the highest bit)
// of the 8-bit value.
func (flags *ConditionFlagsImpl) UpdateCarry(original uint8, new uint8) bool {
	originalCarryBit := (original & carryMask) >> 7
	newCarryBit := (new & carryMask) >> 7
	flags.Carry = (originalCarryBit == 1) && (newCarryBit == 0)
	return flags.Carry
}

// UpdateCarryDoublePrecision updates the Carry flag based the change from the original
// value to the new value. It is set if a carry occurs out-of bit 15 (the highest bit)
// of the 16-bit value.
func (flags *ConditionFlagsImpl) UpdateCarryDoublePrecision(original uint16, new uint16) bool {
	originalCarryBit := (original & doubleCarryMask) >> 15
	newCarryBit := (new & doubleCarryMask) >> 15
	flags.Carry = (originalCarryBit == 1) && (newCarryBit == 0)
	return flags.Carry
}

// ClearCarry clears the Carry flag.
func (flags *ConditionFlagsImpl) ClearCarry() {
	flags.Carry = false
}

// IsAuxiliaryCarry returns the value of the Auxiliary Carry flag
func (flags *ConditionFlagsImpl) IsAuxiliaryCarry() bool {
	return flags.AuxillaryCarry
}

// UpdateAuxiliaryCarry updates the Auxiliary Carry flag based on the value of the provided
// result and returns the Auxiliary Carry flag. This flag is set when a carry occurs between
// bits 3 and 4 of the low nibble.
func (flags *ConditionFlagsImpl) UpdateAuxiliaryCarry(original uint8, new uint8) bool {
	origBits := uint8((0x18 & original) >> 3)
	newBits := uint8((0x18 & new) >> 3)
	flags.AuxillaryCarry = origBits == 1 && newBits == 2
	return flags.AuxillaryCarry
}

// SetAuxiliaryCarry sets the Auxiliary Carry flag.
func (flags *ConditionFlagsImpl) SetAuxiliaryCarry() {
	flags.AuxillaryCarry = true
}

// ClearAuxiliaryCarry clears the Auxiliary Carry flag.
func (flags *ConditionFlagsImpl) ClearAuxiliaryCarry() {
	flags.AuxillaryCarry = false
}
