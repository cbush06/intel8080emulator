package alu

// Add adds reg1 to reg2 and returns the unsigned result.
// The global AluFlags will be updated based on this
// operation's result.
func Add(reg1 uint8, reg2 uint8) uint8 {
	var result = uint16(reg1) + uint16(reg2)

	AluFlags.ClearFlags()
	AluFlags.UpdateZero(result)
	AluFlags.UpdateSign(result)
	AluFlags.UpdateParity(uint8(result))
	AluFlags.UpdateCarry(result)

	return uint8(result)
}

// Sub subtracts reg2 from reg1 and returns the unsigned result.
// The global AluFlags will be updated based on this
// operations result.
func Sub(reg1 uint8, reg2 uint8) uint8 {
	return Add(reg1, uint8(-1*int8(reg2)))
}
