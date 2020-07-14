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
	UpdateCarry(original uint8, new uint8) bool
	UpdateCarryDoublePrecision(original uint16, new uint16) bool
	ClearCarry()
	IsAuxillaryCarry() bool
	UpdateAuxillaryCarry(original uint8, new uint8) bool
	SetAuxillaryCarry()
	ClearAuxillaryCarry()
}
