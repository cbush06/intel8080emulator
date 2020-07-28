package alu

import "testing"

func TestConditionFlagsImpl_ClearFlags(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Zero:           true,
		Sign:           true,
		Parity:         true,
		Carry:          true,
		AuxillaryCarry: true,
	}

	cndFlags.ClearFlags()

	if cndFlags.Zero || cndFlags.Sign || cndFlags.Parity || cndFlags.Carry || cndFlags.AuxillaryCarry {
		t.Errorf("Expected all flags cleared but got %+v", cndFlags)
	}
}

func TestConditionFlagsImpl_CreateStatusWord(t *testing.T) {
	var (
		allFlagsSet uint8 = 0xD7 // 11010111
		noFlagsSet  uint8 = 0x02 // 00000010
	)

	cndFlags := &ConditionFlagsImpl{
		Carry:          true,
		Parity:         true,
		AuxillaryCarry: true,
		Zero:           true,
		Sign:           true,
	}

	if status := cndFlags.CreateStatusWord(); status != allFlagsSet {
		t.Errorf("Expected %bb but got %bb", allFlagsSet, status)
	}

	cndFlags.ClearFlags()

	if status := cndFlags.CreateStatusWord(); status != noFlagsSet {
		t.Errorf("Expected %bb but got %bb", noFlagsSet, status)
	}
}

func TestConditionFlagsImpl_ApplyStatusWord(t *testing.T) {
	var (
		allFlagsSet uint8 = 0xD7 // 11010111
		noFlagsSet  uint8 = 0x02 // 00000010
	)

	cndFlags := new(ConditionFlagsImpl)
	cndNoFlagsSet := &ConditionFlagsImpl{
		Zero:           false,
		Sign:           false,
		Parity:         false,
		Carry:          false,
		AuxillaryCarry: false,
	}
	cndAllFlagsSet := &ConditionFlagsImpl{
		Zero:           true,
		Sign:           true,
		Parity:         true,
		Carry:          true,
		AuxillaryCarry: true,
	}

	cndFlags.ApplyStatusWord(allFlagsSet)
	if *cndFlags != *cndAllFlagsSet {
		t.Errorf("Expected %+v but got %+v", cndAllFlagsSet, cndFlags)
	}

	cndFlags.ApplyStatusWord(noFlagsSet)
	if *cndFlags != *cndNoFlagsSet {
		t.Errorf("Expected %+v but got %+v", cndNoFlagsSet, cndFlags)
	}
}

func TestConditionFlagsImpl_IsZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}

	cndFlags.Zero = true
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_SetZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetZero()
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_UpdateZero(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	cndFlags.UpdateZero(0)
	if !cndFlags.IsZero() {
		t.Error("Expected true but got false")
	}

	cndFlags.UpdateZero(1)
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_ClearZero(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Zero: true,
	}

	cndFlags.ClearZero()
	if cndFlags.IsZero() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_IsSign(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}

	cndFlags.Sign = true
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_SetSign(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetSign()
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_UpdateSign(t *testing.T) {
	var (
		signedResult   uint8 = 0x80
		unsignedResult uint8 = 0x40
	)

	cndFlags := &ConditionFlagsImpl{
		Sign: false,
	}

	cndFlags.UpdateSign(signedResult)
	if !cndFlags.IsSign() {
		t.Error("Expected true but got false")
	}

	cndFlags.UpdateSign(unsignedResult)
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_ClearSign(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Sign: true,
	}

	cndFlags.ClearSign()
	if cndFlags.IsSign() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_IsParity(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.Parity = true
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_SetParity(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetParity()
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_UpdateParity(t *testing.T) {
	var (
		oddParity  uint8 = 0x57
		evenParity uint8 = 0x55
	)

	cndFlags := new(ConditionFlagsImpl)

	cndFlags.UpdateParity(oddParity)
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}

	cndFlags.UpdateParity(evenParity)
	if !cndFlags.IsParity() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_ClearParity(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Parity: true,
	}

	cndFlags.ClearParity()
	if cndFlags.IsParity() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_IsCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.Carry = true
	if !cndFlags.IsCarry() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_SetCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)
	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetCarry()
	if !cndFlags.IsCarry() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_UpdateCarry(t *testing.T) {
	var (
		addend      uint8 = 0x80
		addend2     uint8 = 0x08
		sum         uint8 = 0x00
		sum2        uint8 = 0x10
		minuend     uint8 = 0x80
		minuend2    uint8 = 0x81
		difference  uint8 = 0x40
		difference2 uint8 = 0x80
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("AddRegister w/ Carry", func(t *testing.T) {
		cndFlags.UpdateCarry(addend, sum) // 0x80 + 0x80 = 0x00 // overflow occurs as a carry out
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("AddRegister w/o Carry", func(t *testing.T) {
		cndFlags.UpdateCarry(addend2, sum2) // 0x08 + 0x08 = 0x10 // this qualifies as an axuillary carry, but not a carry
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})

	t.Run("Subtract w/ Borrow", func(t *testing.T) {
		cndFlags.UpdateCarry(minuend, difference) // 0x80 - 0x40 = 0x40 // borrow occurs from bit 7 to bit 6
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Subtract w/o Borrow", func(t *testing.T) {
		cndFlags.UpdateCarry(minuend2, difference2) // 0x81 - 0x01 = 0x80 // no borrow
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestConditionFlagsImpl_UpdateCarryDoublePrecision(t *testing.T) {
	var (
		addend      uint16 = 0x8000
		addend2     uint16 = 0x0800
		sum         uint16 = 0x0000
		sum2        uint16 = 0x1000
		minuend     uint16 = 0x8000
		minuend2    uint16 = 0x8100
		difference  uint16 = 0x4000
		difference2 uint16 = 0x8000
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("AddRegister with Carry", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(addend, sum) // 0x80 + 0x80 = 0x00 // overflow occurs as a carry out
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("AddRegister without Carry", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(addend2, sum2) // 0x08 + 0x08 = 0x10 // this qualifies as an auxillary carry, but not a carry
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})

	t.Run("Subtract with Borrow", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(minuend, difference) // 0x80 - 0x40 = 0x40 // borrow occurs from bit 7 to bit 6
		if !cndFlags.IsCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("Subtract without Borrow", func(t *testing.T) {
		cndFlags.UpdateCarryDoublePrecision(minuend2, difference2) // 0x81 - 0x01 = 0x80 // no borrow
		if cndFlags.IsCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestConditionFlagsImpl_ClearCarry(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		Carry: true,
	}
	cndFlags.ClearCarry()

	if cndFlags.IsCarry() {
		t.Error("Expected false but got true")
	}
}

func TestConditionFlagsImpl_IsAuxillaryCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)

	if cndFlags.IsAuxiliaryCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetAuxiliaryCarry()
	if !cndFlags.IsAuxiliaryCarry() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_UpdateAuxiliaryCarry(t *testing.T) {
	var (
		addend      uint8 = 0x08
		addend2     uint8 = 0x01
		sum         uint8 = 0x10
		sum2        uint8 = 0x09
	)

	cndFlags := new(ConditionFlagsImpl)

	t.Run("AddRegister with Auxiliary Carry", func(t *testing.T) {
		cndFlags.UpdateAuxiliaryCarry(addend, sum) // 0x08 + 0x08 = 0x10 // addition causes a carry from bit 3 into bit 4
		if !cndFlags.IsAuxiliaryCarry() {
			t.Error("Expected true but got false")
		}
	})

	t.Run("AddRegister without Auxiliary Carry", func(t *testing.T) {
		cndFlags.UpdateAuxiliaryCarry(addend2, sum2) // 0x08 + 0x01 = 0x09 // this does not cause an auxiliary carry
		if cndFlags.IsAuxiliaryCarry() {
			t.Error("Expected false but got true")
		}
	})
}

func TestConditionFlagsImpl_SetAuxiliaryCarry(t *testing.T) {
	cndFlags := new(ConditionFlagsImpl)

	if cndFlags.IsAuxiliaryCarry() {
		t.Error("Expected false but got true")
	}

	cndFlags.SetAuxiliaryCarry()
	if !cndFlags.IsAuxiliaryCarry() {
		t.Error("Expected true but got false")
	}
}

func TestConditionFlagsImpl_ClearAuxiliaryCarry(t *testing.T) {
	cndFlags := &ConditionFlagsImpl{
		AuxillaryCarry: true,
	}

	cndFlags.ClearAuxiliaryCarry()
	if cndFlags.IsAuxiliaryCarry() {
		t.Error("Expected false but go true")
	}
}
