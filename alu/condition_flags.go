package alu

// ConditionFlags is the interface to the Intel 8080's ALU condition flags
type ConditionFlags interface {
	ClearFlags()
	CreateStatusWord() uint8
	ApplyStatusWord(statusWord uint8)
	IsZero() bool
	SetZero()
	UpdateZero(result uint8) bool
	ClearZero()
	IsSign() bool
	SetSign()
	UpdateSign(result uint8) bool
	ClearSign()
	IsParity() bool
	SetParity()
	UpdateParity(data uint8) bool
	ClearParity()
	IsCarry() bool
	SetCarry()
	UpdateCarry(result uint16) bool
	UpdateCarryDoublePrecision(result uint32) bool
	ClearCarry()
	IsAuxiliaryCarry() bool
	UpdateAuxiliaryCarry(original uint8, result uint8) bool
	SetAuxiliaryCarry()
	ClearAuxiliaryCarry()
}
