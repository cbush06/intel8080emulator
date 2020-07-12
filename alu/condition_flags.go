package alu

// ConditionFlags is the interface to the Intel 8080's ALU condition flags
type ConditionFlags interface {
	ClearFlags()
	CreateStatusWord() uint8
	ApplyStatusWord(statusWord uint8)
	IsZero() bool
	SetZero()
	UpdateZero(result uint16) bool
	ClearZero()
	IsSign() bool
	SetSign()
	UpdateSign(result uint16) bool
	ClearSign()
	IsParity() bool
	SetParity()
	UpdateParity(data uint16) bool
	ClearParity()
	IsCarry() bool
	SetCarry()
	UpdateBorrow(addend1 uint8, addend2 uint8) bool
	UpdateCarry(result uint16) bool
	UpdateCarryDoublePrecision(result uint32) bool
	ClearCarry()
	IsAuxillaryCarry() bool
	SetAuxillaryCarry()
	ClearAuxillaryCarry()
}
