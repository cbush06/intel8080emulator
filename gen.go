// +build ignore

package main

//go:generate mockgen -destination=alu/mocks/condition_flags_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ConditionFlags
//go:generate mockgen -destination=alu/mocks/alu_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ALU
